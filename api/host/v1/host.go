package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `method:"get" path:"/host/{id}/one" summary:"获取主机" tags:"主机"`
	Id     int `v:"required" path:"id"`
}

type GetRes struct {
	*entity.Host
}

type GetAuthorizedLstReq struct {
	g.Meta `method:"get" path:"/host/authorized-list" summary:"获取当前用户拥有权限的主机列表" tags:"主机"`
}

type GetAuthorizedLstRes struct {
	List []*entity.Host `json:"list"`
}

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/host/page-list" summary:"分页获取主机列表" tags:"主机"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.Host]
}

type AddReq struct {
	g.Meta `method:"post" path:"/host" summary:"添加主机" tags:"主机"`
	*mid.Host
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/host/{id}" summary:"更新主机" tags:"主机"`
	Id     int `v:"min:1#id必须" path:"id"`
	*mid.Host
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/host/{id}" summary:"删除主机" tags:"主机"`
	Id     int `v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type TestSshReq struct {
	g.Meta `method:"get" path:"/host/{id}/ssh-ok" summary:"测试主机是否可以成功建立 ssh 连接" tags:"主机"`
	Id     int `v:"min:1#id必须" path:"id"`
}

type TestSshRes struct{}

type DownloadFileReq struct {
	g.Meta `method:"get" path:"/host/{id}/download-file" summary:"下载文件" tags:"主机"`
	Id     int    `v:"required" path:"id"`
	Path   string `v:"required" p:"path"`
}

type DownloadFileRes struct{}
