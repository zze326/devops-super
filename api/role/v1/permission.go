package v1

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type UptPermissionReq struct {
	g.Meta        `method:"patch" path:"/role/{id}/permission" summary:"更新角色关联的权限" tags:"角色"`
	Id            int         `v:"min:1#id必须" path:"id"`
	PermissionIds *gjson.Json `json:"permissionIds"`
}

type UptPermissionRes struct{}
