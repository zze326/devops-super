package ci_pipeline_run

import (
	"bufio"
	"context"
	"devops-super/utility/thirdclients/kubernetes"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/logrusorgru/aurora"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
)

type wsContext struct {
	request      *ghttp.Request
	ws           *ghttp.WebSocket
	ctx          context.Context
	cancelFunc   context.CancelFunc
	kubeClient   *kubernetes.Client
	watcher      watch.Interface
	namespace    string
	podName      string
	lastPingTime *gtime.Time
}

func (wsCtx *wsContext) writeErr(err error) {
	wsCtx.ws.WriteMessage(ghttp.WsMsgText, []byte(aurora.BgRed(err.Error()).String()))
}

// 获取 Pod 日志
func (wsCtx *wsContext) tailLog(status corev1.ContainerStatus) error {
	var (
		containerName = status.Name
		isKaniko      = strings.Contains(containerName, "kaniko")
		msg           = "Kaniko 构建上传镜像"
	)
	if isKaniko {
		if strings.Contains(containerName, "warmer") {
			msg = "更新基础镜像缓存"
		}
		if err := wsCtx.ws.WriteMessage(ghttp.WsMsgText, []byte(aurora.Green(fmt.Sprintf("=====【%s】开始=====", msg)).String())); err != nil {
			return err
		}
	}
	line := int64(100000)
	req := wsCtx.kubeClient.CoreV1().Pods(wsCtx.namespace).GetLogs(wsCtx.podName, &corev1.PodLogOptions{
		Container: containerName,
		Follow:    true,
		TailLines: &line,
	})
	stream, err := req.Stream(wsCtx.ctx)
	if err != nil {
		return err
	}
	defer stream.Close()
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		if err := wsCtx.ws.WriteMessage(ghttp.WsMsgText, scanner.Bytes()); err != nil {
			return err
		}
	}
	if isKaniko && ((status.State.Terminated != nil && status.State.Terminated.ExitCode == 0) || status.State.Running != nil) {
		if err := wsCtx.ws.WriteMessage(ghttp.WsMsgText, []byte(aurora.Green(fmt.Sprintf("=====【%s】结束=====", msg)).String())); err != nil {
			return err
		}
	}
	return nil
}

func (wsCtx *wsContext) checkClientClose() {
	var handleClose = func() {
		wsCtx.cancelFunc()
		if wsCtx.watcher != nil {
			wsCtx.watcher.Stop()
		}
		wsCtx.ws.Close()
	}

	for {
		_, _, err := wsCtx.ws.ReadMessage() // 仅用来监听客户端连接关闭
		if err != nil {
			handleClose()
			break
		}
	}
}
