package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestConnectKubernetesReq struct {
	g.Meta `method:"post" path:"/common/kubernetes/test-connect" summary:"测试配置是否可以正常连接到 Kubernetes" tags:"公共"`
	Config string `v:"required" json:"config"`
}

type TestConnectKubernetesRes struct{}
