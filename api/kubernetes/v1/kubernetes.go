package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestConnectReq struct {
	g.Meta `method:"post" path:"/kubernetes/test-connect" summary:"测试配置是否可以正常连接到 Kubernetes" tags:"Kubernetes"`
	Config string `v:"required" json:"config"`
}

type TestConnectRes struct{}

type GetNamespaceLstReq struct {
	g.Meta   `method:"get" path:"/kubernetes/namespace/list" summary:"获取集群命名空间" tags:"Kubernetes"`
	SecretId int `v:"required" p:"secretId"`
}

type GetNamespaceLstRes struct {
	Namespaces []string `json:"namespaces"`
}

type GetPersistentVolumeClaimLstReq struct {
	g.Meta    `method:"get" path:"/kubernetes/persistent-volume-claim/list" summary:"获取集群指定命名空间中的 PVC 列表" tags:"Kubernetes"`
	SecretId  int    `v:"required" p:"secretId"`
	Namespace string `v:"required" p:"namespace"`
}

type GetPersistentVolumeClaimLstRes struct {
	Pvcs []string `json:"pvcs"`
}
