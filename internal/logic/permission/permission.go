package permission

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sPermission struct{}

var (
	cols = dao.Permission.Columns()
)

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Add(ctx context.Context, in *entity.Permission) (err error) {
	r, err := s.Get(ctx, &do.Permission{Name: in.Name})
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.Newf("已存在名称为 %s 权限", in.Name)
	}
	if (in.Type == consts.PERMISSION_TYPE_DIR || in.Type == consts.PERMISSION_TYPE_MENU) && !gutil.IsEmpty(in.FRoute) {
		r, err = s.Get(ctx, &do.Permission{FRoute: in.FRoute})
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.Newf("已存在前端路由为 %s 权限", in.FRoute)
		}
	}
	_, err = dao.Permission.Ctx(ctx).Insert(in)
	return
}

func (*sPermission) Get(ctx context.Context, in *do.Permission) (out *entity.Permission, err error) {
	err = dao.Permission.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sPermission) GetLst(ctx context.Context, search string) (out []*entity.Permission, err error) {
	m := dao.Permission.Ctx(ctx).Order(dao.Permission.Columns().Rank).Safe(true)
	if !gutil.IsEmpty(search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, util.SqlLikeStr(search)).WhereOrLike(cols.Title, util.SqlLikeStr(search)))
	}
	err = m.Scan(&out)
	return
}

func (*sPermission) Upt(ctx context.Context, in *do.Permission) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.Permission.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}

		if r != nil {
			return gerror.Newf("已存在名称为 %s 权限", in.Name)
		}
	}
	if (in.Type == consts.PERMISSION_TYPE_DIR || in.Type == consts.PERMISSION_TYPE_MENU) && !gutil.IsEmpty(in.FRoute) {
		r, err := dao.Permission.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.FRoute, in.FRoute).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.Newf("已存在前端路由为 %s 权限", in.FRoute)
		}
	}

	_, err = dao.Permission.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sPermission) Del(ctx context.Context, in *do.Permission) (err error) {
	_, err = dao.Permission.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}

func (s *sPermission) SystemRequired(ctx context.Context) (ePermission *entity.Permission, err error) {
	requiredPermission, err := s.Get(ctx, &do.Permission{Name: consts.PERMISSION_SYSTEM_REQUIRED_NAME})
	if err != nil {
		return nil, err
	}
	if requiredPermission != nil && requiredPermission.Id > 0 {
		return requiredPermission, nil
	}
	return nil, gerror.New("系统必需权限缺失")
}
