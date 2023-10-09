package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/kubernetes-config/page-list" summary:"分页获取 Kubernetes 配置列表" tags:"Kubernetes 配置"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.KubernetesConfig]
}

type AddReq struct {
	g.Meta `method:"post" path:"/kubernetes-config" summary:"添加 Kubernetes 配置" tags:"Kubernetes 配置"`
	*mid.KubernetesConfig
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/kubernetes-config/{id}" summary:"更新 Kubernetes 配置" tags:"Kubernetes 配置"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.KubernetesConfig
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/kubernetes-config/{id}" summary:"删除 Kubernetes 配置" tags:"Kubernetes 配置"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/kubernetes-config/list" summary:"获取所有 Kubernetes 配置列表" tags:"Kubernetes 配置"`
	*api.PageLstReq
}

type GetLstRes struct {
	List []*entity.KubernetesConfig `json:"list"`
}
