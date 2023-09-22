package permission

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/entity/comb"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
)

func (*sPermission) GetRouteLst(ctx context.Context) (out []*mid.Route, err error) {
	cUser, err := service.User().GetComb(ctx, &do.User{Id: service.CurrentUser(ctx).UserId})
	if err != nil {
		return nil, err
	}

	var permissionIdSet = gset.New()
	for _, role := range cUser.Roles {
		if role.Permission.IsNil() {
			continue
		}
		for _, pId := range role.Permission.Array() {
			permissionIdSet.Add(pId)
		}
	}

	var cPermissions []*comb.Permission
	if err = dao.Permission.Ctx(ctx).Order(cols.Rank).WithAll().WhereIn(cols.Id, permissionIdSet.Slice()).Scan(&cPermissions); err != nil {
		return
	}

	var ePermissionMap = make(map[int]*entity.Permission, len(cPermissions))
	for _, permission := range cPermissions {
		ePermissionMap[permission.Id] = permission.Permission
	}

	type (
		handleManyFunc   = func(children []*comb.Permission) error
		handleOneFunc    = func(cPermission *comb.Permission) error
		handleParentFunc = func(parentId int) (*mid.Route, error)
	)

	var (
		finalRoutes mid.FrontendRouteList
		routeMap    = make(map[int]*mid.Route)
		getOne      = func(id int) *entity.Permission {
			return ePermissionMap[id]
		}
		permissionToRoute = func(ePermission *entity.Permission) (route *mid.Route) {
			_ = gconv.Struct(ePermission, &route)
			return
		}
		handleMany   handleManyFunc
		handleOne    handleOneFunc
		handleParent handleParentFunc
	)

	handleParent = func(parentId int) (*mid.Route, error) {
		var (
			parent *mid.Route
			ok     bool
		)
		parent, ok = routeMap[parentId]
		if !ok {
			ePermission := new(entity.Permission)
			if permissionIdSet.Contains(parentId) {
				ePermission = getOne(parentId)
			} else {
				if err := dao.Permission.Ctx(ctx).WherePri(parentId).Scan(ePermission); err != nil {
					return nil, err
				}
				parent = permissionToRoute(ePermission)
				routeMap[parentId] = parent
				finalRoutes = append(finalRoutes, parent)
			}
			if parent.ParentId > 0 {
				if _, err := handleParent(parent.ParentId); err != nil {
					return nil, err
				}
			}
		}

		return parent, nil
	}

	handleMany = func(children []*comb.Permission) error {
		for _, child := range children {
			if err := handleOne(child); err != nil {
				return err
			}
		}
		return nil
	}

	handleOne = func(cPermission *comb.Permission) error {
		switch cPermission.Type {
		case consts.PERMISSION_TYPE_DIR, consts.PERMISSION_TYPE_MENU:
			route := permissionToRoute(cPermission.Permission)
			if cPermission.Type == consts.PERMISSION_TYPE_DIR {
				if err := handleMany(cPermission.Children); err != nil {
					return err
				}
			} else {
				route.Auths = cPermission.AuthCodes()
			}

			if _, ok := routeMap[route.Id]; !ok {
				routeMap[route.Id] = route
				finalRoutes = append(finalRoutes, route)
			}
		}

		if parentId := cPermission.ParentId; parentId > 0 {
			parent, err := handleParent(parentId)
			if err != nil {
				return err
			}

			if cPermission.Type == consts.PERMISSION_TYPE_ABLE {
				parent.Auths = append(parent.Auths, cPermission.Name)
			}
		}
		return nil
	}

	if err = handleMany(cPermissions); err != nil {
		return
	}
	sort.Sort(finalRoutes)
	return finalRoutes, nil
}
