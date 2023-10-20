package kubernetes

import (
	"context"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

type sKubernetes struct{}

func init() {
	service.RegisterKubernetes(New())
}

func New() *sKubernetes {
	return &sKubernetes{}
}

func (*sKubernetes) TestConnect(ctx context.Context, config string) (err error) {
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

func (*sKubernetes) GetNamespaces(ctx context.Context, config string) ([]string, error) {
	client, err := kubernetes.NewClient(ctx, config)
	if err != nil {
		return nil, gerror.Wrap(err, "配置格式错误")
	}
	return client.GetNamespaces()
}

func (*sKubernetes) GetPersistentVolumeClaims(ctx context.Context, config, namespace string) ([]string, error) {
	client, err := kubernetes.NewClient(ctx, config)
	if err != nil {
		return nil, gerror.Wrap(err, "配置格式错误")
	}
	return client.GetPersistentVolumeClaims(namespace)
}
