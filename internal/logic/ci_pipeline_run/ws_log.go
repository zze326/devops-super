package ci_pipeline_run

import (
	"context"
	"devops-super/internal/model/entity"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
)

func (s *sCiPipelineRun) WsLog(ctx context.Context, id int) (err error) {
	var (
		eCiPipelineRun = new(entity.CiPipelineRun) // 运行记录
		podInfo        *corev1.Pod                 // ci pod 信息
		podFinished    bool                        // pod 是否已执行完毕
		wsCtx          = &wsContext{
			request: g.RequestFromCtx(ctx),
		}
	)
	wsCtx.ctx, wsCtx.cancelFunc = context.WithCancel(ctx)
	ws, err := g.RequestFromCtx(ctx).WebSocket()
	if err != nil {
		return err
	}
	defer ws.Close()
	wsCtx.ws = ws
	go wsCtx.checkClientClose() // 监听客户端连接关闭

	eCiPipelineRun, wsCtx.kubeClient, err = s.GetWithKubernetesClient(ctx, id)
	if err != nil {
		wsCtx.writeErr(err)
		return err
	}

	wsCtx.namespace = eCiPipelineRun.Namespace
	wsCtx.podName = eCiPipelineRun.PodName

	_, err = wsCtx.kubeClient.GetPod(wsCtx.ctx, eCiPipelineRun.Namespace, eCiPipelineRun.PodName)
	if err != nil {
		wsCtx.writeErr(err)
		return err
	}

	logIndex := 0 // 从第一个容器开始查看日志
	// 创建 pod 监听
	watcher, err := wsCtx.kubeClient.CoreV1().Pods(eCiPipelineRun.Namespace).Watch(wsCtx.ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", eCiPipelineRun.PodName),
	})
	if err != nil {
		err = gerror.Wrap(err, "Failed to create watcher")
		wsCtx.writeErr(err)
		return err
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
				for _, status := range allContainerStatus {
					if status.Ready || status.State.Terminated != nil {
						if err = wsCtx.tailLog(status); err != nil {
							return
						}
					}
				}
				break WATCH
			} else {
				// 把已经运行完毕和正在运行的容器的日志先输出
				for _, containerStatus := range allContainerStatus {
					if containerStatus.Ready {
						if err = wsCtx.tailLog(containerStatus); err != nil {
							return err
						}
						logIndex++
					} else {
						if containerStatus.State.Running != nil {
							logIndex++
							if err = wsCtx.tailLog(containerStatus); err != nil {
								return err
							}
							break
						}
					}
				}
				// 如果发现所有容器的日志已经输出完，则中断输出
				if logIndex > len(allContainerStatus)-1 {
					break WATCH
				}
			}
		case watch.Modified: // 监听 pod 变化
			podInfo = event.Object.(*corev1.Pod)
			// 如果 Pod 运行失败了，直接输出当前容器的日志然后中断
			if podInfo.Status.Phase == corev1.PodFailed {
				//if err := wsCtx.tailLog(logIndex); err != nil {
				//	return err
				//}
				break WATCH
			}
			for _, status := range append(podInfo.Status.InitContainerStatuses, podInfo.Status.ContainerStatuses...) {
				maxIndex := len(podInfo.Status.InitContainerStatuses) + len(podInfo.Status.ContainerStatuses) - 1
				if logIndex > maxIndex {
					break WATCH
				}
				if containerName := fmt.Sprintf("env-%d", logIndex); status.Name == containerName || strings.HasPrefix(status.Name, fmt.Sprintf("%s-", containerName)) {
					containerName = status.Name
					canLog := status.Ready || status.State.Running != nil || status.State.Terminated != nil
					if !canLog {
						continue
					}
					if err := wsCtx.tailLog(status); err != nil {
						return err
					}
					// 最后一个容器日志获取完毕才终止监听
					if logIndex == maxIndex && podInfo.Status.Phase != corev1.PodRunning {
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
