package context

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/model/entity/comb"
	"devops-super/internal/service"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

func (*sContext) initCasbin(ctx context.Context) (*casbin.Enforcer, error) {
	g.Log().Debug(ctx, "init casbin enforcer start.")
	enforcer, err := casbin.NewEnforcer(newCasbinModel())
	if err != nil {
		return nil, err
	}
	roleList, err := service.Role().GetCombList(ctx)
	if err != nil {
		return nil, err
	}

	type recursiveHandlePermissionsFunc func(roleCode string, permissions []*comb.Permission) error
	var recursiveHandlePermissions recursiveHandlePermissionsFunc

	recursiveHandlePermissions = func(roleCode string, permissions []*comb.Permission) error {
		for _, permission := range permissions {
			if permission.Type == consts.PERMISSION_TYPE_ABLE {
				for _, backendRouteInfo := range permission.BRoutes.Array() {
					routeInfoArr := strings.SplitN(backendRouteInfo.(string), ":", 2)
					method := strings.ToUpper(routeInfoArr[0])
					routePath := routeInfoArr[1]
					if enforcer.HasPolicy(roleCode, routePath, method) {
						continue
					}
					if _, err := enforcer.AddPolicy(roleCode, routePath, method); err != nil {
						return err
					}
					g.Log().Debugf(ctx, "casbin policy added: %s, %s, %s", roleCode, routePath, method)
				}
			} else {
				if len(permission.Children) > 0 {
					if err = recursiveHandlePermissions(roleCode, permission.Children); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}

	for _, role := range roleList {
		if err := recursiveHandlePermissions(role.Code, role.Permissions); err != nil {
			return nil, err
		}
	}

	userList, err := service.User().GetCombLst(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range userList {
		var roleCodes []string
		for _, role := range user.Roles {
			roleCodes = append(roleCodes, role.Code)
		}
		if _, err := enforcer.AddRolesForUser(user.Username, roleCodes); err != nil {
			return nil, err
		}
		g.Log().Debugf(ctx, "casbin roles added: %s, %v", user.Username, roleCodes)
	}
	g.Log().Debug(ctx, "init casbin enforcer successful.")
	return enforcer, nil
}

func newCasbinModel() model.Model {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", `g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || g(r.sub, "admin")`)
	return m
}
