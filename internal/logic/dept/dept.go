package permission

import (
	"context"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sDept struct{}

var (
	cols = dao.Dept.Columns()
)

func init() {
	service.RegisterDept(New())
}

func New() *sDept {
	return &sDept{}
}

func (s *sDept) Add(ctx context.Context, in *entity.Dept) (err error) {
	r, err := s.Get(ctx, &do.Dept{Name: in.Name})
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.Newf("已存在名称为 %s 部门", in.Name)
	}
	_, err = dao.Dept.Ctx(ctx).Insert(in)
	return
}

func (*sDept) Get(ctx context.Context, in *do.Dept) (out *entity.Dept, err error) {
	err = dao.Dept.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sDept) GetLst(ctx context.Context, search string) (out []*entity.Dept, err error) {
	m := dao.Dept.Ctx(ctx).Order(dao.Dept.Columns().Rank).Safe(true)
	if !gutil.IsEmpty(search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, util.SqlLikeStr(search)))
	}
	err = m.Scan(&out)
	return
}

func (*sDept) Upt(ctx context.Context, in *do.Dept) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.Dept.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}

		if r != nil {
			return gerror.Newf("已存在名称为 %s 部门", in.Name)
		}
	}
	_, err = dao.Dept.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sDept) Del(ctx context.Context, in *do.Dept) (err error) {
	_, err = dao.Dept.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
