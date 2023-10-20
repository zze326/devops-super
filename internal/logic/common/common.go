package common

import (
	"context"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}

func (*sCommon) TestConnectKubernetes(ctx context.Context, config string) (err error) {
	timeoutCtx, _ := context.WithTimeout(ctx, time.Second*2)
	client, err := kubernetes.NewClient(timeoutCtx, config)
	if err != nil {
		return gerror.Wrap(err, "配置格式错误")
	}
	if _, err = client.GetNamespaces(); err != nil {
		return gerror.Wrap(err, "连接失败")
	}
	return nil
}
