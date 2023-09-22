package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/entity/comb"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"github.com/gogf/gf/v2/util/gutil"
)

type sUser struct{}

var cols = dao.User.Columns()

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (*sUser) Add(ctx context.Context, in *entity.User) (err error) {
	in.Password, err = util.EncryptPassword(in.Password)
	if err != nil {
		return
	}
	_, err = dao.User.Ctx(ctx).Insert(in)
	return
}

func (*sUser) Upt(ctx context.Context, in *do.User) (err error) {
	if !gutil.IsEmpty(in.Password) {
		in.Password, err = util.EncryptPassword(in.Password.(string))
		if err != nil {
			return
		}
	}
	_, err = dao.User.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sUser) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.User], err error) {
	out = &api.PageLstRes[*entity.User]{}
	m := dao.User.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Username, in.SearchStr()).WhereOrLike(cols.RealName, in.SearchStr()))
	}

	if enabled := in.Wheres.Get("enabled"); !enabled.IsNil() {
		m = m.Where(cols.Enabled, enabled.Bool())
	}

	if deptId := in.Wheres.Get("deptId"); !deptId.IsNil() {
		m = m.Where(cols.DeptId, deptId.Int())
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).FieldsEx(cols.Password).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sUser) Get(ctx context.Context, userDo *do.User) (out *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(userDo).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sUser) GetComb(ctx context.Context, userDo *do.User) (out *comb.User, err error) {
	if err = dao.User.Ctx(ctx).Where(userDo).OmitNilWhere().Limit(1).Scan(&out); err != nil {
		return
	}
	if err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, out.RoleIds.Array()).Scan(&out.Roles); err != nil {
		return
	}
	return
}

func (*sUser) GetCombLst(ctx context.Context) (out []*comb.User, err error) {
	if err = dao.User.Ctx(ctx).Where(cols.Enabled, true).Scan(&out); err != nil {
		return
	}

	for _, user := range out {
		if err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, user.RoleIds.Array()).Scan(&user.Roles); err != nil {
			return
		}
	}
	return
}

func (*sUser) Del(ctx context.Context, in *do.User) (err error) {
	_, err = dao.User.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
