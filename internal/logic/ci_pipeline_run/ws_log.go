package ci_pipeline_run

import (
	"context"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (s *sCiPipelineRun) WsLog(ctx context.Context, id int) (err error) {
	var (
		eCiPipelineRun = new(entity.CiPipelineRun) // 运行记录
		eCiPipeline    *entity.CiPipeline          // 源 pipeline
		eSecret        *entity.Secret              // 秘钥
		kubeConfig     = new(mid.TextContent)      // k8s 配置内容
		podInfo        *corev1.Pod                 // ci pod 信息
		podFinished    bool                        // pod 是否已执行完毕
		wsCtx          = &wsContext{
			request: g.RequestFromCtx(ctx),
		}
	)
	wsCtx.ctx, wsCtx.cancelFunc = context.WithCancel(ctx)
	if err = dao.CiPipelineRun.Ctx(ctx).WherePri(id).Scan(eCiPipelineRun); err != nil {
		return
	}
	wsCtx.namespace = eCiPipelineRun.Namespace
	wsCtx.podName = eCiPipelineRun.PodName

	if eCiPipeline, err = service.CiPipeline().Get(ctx, &do.CiPipeline{Id: eCiPipelineRun.PipelineId}); err != nil {
		return
	}

	if eCiPipeline == nil {
		return gerror.New("找不到源流水线")
	}

	if eSecret, err = service.Secret().Get(ctx, &do.Secret{Id: eCiPipeline.KubernetesConfigId}); err != nil {
		return
	}
	if eSecret == nil {
		return gerror.New("找不到 Kubernetes 配置")
	}

	if err = eSecret.Content.Scan(kubeConfig); err != nil {
		return err
	}

	if wsCtx.kubeClient, err = kubernetes.NewClient(ctx, kubeConfig.Text); err != nil {
		return gerror.Wrap(err, "创建 kubernetes client 失败")
	}

	ws, err := g.RequestFromCtx(ctx).WebSocket()
	if err != nil {
		return err
	}
	defer ws.Close()
	wsCtx.ws = ws

	go wsCtx.checkClientClose() // 监听客户端连接关闭
	logIndex := 0               // 从第一个容器开始查看日志
	// 创建 pod 监听
	watcher, err := wsCtx.kubeClient.CoreV1().Pods(eCiPipelineRun.Namespace).Watch(wsCtx.ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", eCiPipelineRun.PodName),
	})
	if err != nil {
		return gerror.Wrap(err, "Failed to create watcher")
	}
	defer watcher.Stop()
	wsCtx.watcher = watcher
	// 监听 Pod 运行状态
WATCH:
	for event := range watcher.ResultChan() {
		switch event.Type {
		case watch.Added:
			podInfo = event.Object.(*corev1.Pod)
			podFinished = podInfo.Status.Phase == corev1.PodFailed || podInfo.Status.Phase == corev1.PodSucceeded // pod 是否已经运行结束
			allContainerStatus := append(podInfo.Status.InitContainerStatuses, podInfo.Status.ContainerStatuses...)
			if podFinished { // 如果容器已经运行结束，直接输出所有日志
				for idx, status := range allContainerStatus {
					if status.Ready || status.State.Terminated != nil {
						if err = wsCtx.tailLog(idx); err != nil {
							return
						}
					}
				}
				break WATCH
			} else {
				// 把已经运行完毕和正在运行的容器的日志先输出
				for idx, containerStatus := range allContainerStatus {
					if containerStatus.Ready {
						if err = wsCtx.tailLog(idx); err != nil {
							return err
						}
						logIndex++
					} else {
						if containerStatus.State.Running != nil {
							logIndex++
							if err = wsCtx.tailLog(idx); err != nil {
								return err
							}
							break
						}
					}
				}
				// 如果发现所有容器的日志已经输出完，则终端输出
				if logIndex > len(allContainerStatus)-1 {
					break WATCH
				}
			}
		case watch.Modified: // 监听 pod 变化
			podInfo = event.Object.(*corev1.Pod)
			// 如果 Pod 运行失败了，直接输出当前容器的日志然后中断
			if podInfo.Status.Phase == corev1.PodFailed {
				if err := wsCtx.tailLog(logIndex); err != nil {
					return err
				}
				break WATCH
			}
			for _, status := range append(podInfo.Status.InitContainerStatuses, podInfo.Status.ContainerStatuses...) {
				if containerName := fmt.Sprintf("env-%d", logIndex); status.Name == containerName {
					canLog := status.Ready || status.State.Running != nil || status.State.Terminated != nil
					if !canLog {
						continue
					}
					if err := wsCtx.tailLog(logIndex); err != nil {
						return err
					}

					// 最后一个容器日志获取完毕才终止监听
					if logIndex == len(podInfo.Status.InitContainerStatuses)+len(podInfo.Status.ContainerStatuses)-1 {
						break WATCH
					}
					logIndex++
				}
			}
		case watch.Error:
			glog.Errorf(wsCtx.ctx, "Received watch error: %v", event.Object)
			break WATCH
		}
	}
	return nil
}
