package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/entity/comb"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sRole struct{}

var cols = dao.Role.Columns()

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (*sRole) Add(ctx context.Context, in *entity.Role) (err error) {
	r, err := dao.Role.Ctx(ctx).Where(cols.Code, in.Code).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该代码的权限")
	}
	_, err = dao.Role.Ctx(ctx).Insert(in)
	return
}

func (*sRole) Upt(ctx context.Context, in *do.Role) (err error) {
	if !gutil.IsEmpty(in.Code) {
		r, err := dao.Role.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Code, in.Code).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该代码的权限")
		}
	}
	_, err = dao.Role.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sRole) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Role], err error) {
	out = &api.PageLstRes[*entity.Role]{}
	m := dao.Role.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()).WhereOrLike(cols.Code, in.SearchStr()))
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)

	permission, err := service.Permission().SystemRequired(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range out.List {
		u.Permission = gjson.New(append(u.Permission.Array(), permission.Id))
	}
	return
}

func (*sRole) GetLst(ctx context.Context) (out []*entity.Role, err error) {
	err = dao.Role.Ctx(ctx).OrderDesc(cols.Id).FieldsEx(cols.Permission).Scan(&out)
	return
}

func (*sRole) GetCombList(ctx context.Context) (out []*comb.Role, err error) {
	if err = dao.Role.Ctx(ctx).Scan(&out); err != nil {
		return
	}

	permission, err := service.Permission().SystemRequired(ctx)
	if err != nil {
		return nil, err
	}

	for _, role := range out {
		role.Permission = gjson.New(append(role.Permission.Array(), permission.Id))
		if err = dao.Permission.Ctx(ctx).WithAll().WhereIn(dao.Permission.Columns().Id, role.Permission.Array()).Scan(&role.Permissions); err != nil {
			return
		}
	}
	return
}

func (*sRole) Get(ctx context.Context, in *do.Role) (out *entity.Role, err error) {
	err = dao.Role.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	permission, err := service.Permission().SystemRequired(ctx)
	if err != nil {
		return nil, err
	}
	out.Permission = gjson.New(append(out.Permission.Array(), permission.Id))
	return
}

func (*sRole) Del(ctx context.Context, in *do.Role) (err error) {
	_, err = dao.Role.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
