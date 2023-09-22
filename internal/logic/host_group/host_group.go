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

type sHostGroup struct{}

var (
	cols = dao.HostGroup.Columns()
)

func init() {
	service.RegisterHostGroup(New())
}

func New() *sHostGroup {
	return &sHostGroup{}
}

func (s *sHostGroup) Add(ctx context.Context, in *entity.HostGroup) (err error) {
	r, err := s.Get(ctx, &do.HostGroup{Name: in.Name})
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.Newf("已存在名称为 %s 主机组", in.Name)
	}
	_, err = dao.HostGroup.Ctx(ctx).Insert(in)
	return
}

func (*sHostGroup) Get(ctx context.Context, in *do.HostGroup) (out *entity.HostGroup, err error) {
	err = dao.HostGroup.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sHostGroup) GetLst(ctx context.Context, search string) (out []*entity.HostGroup, err error) {
	m := dao.HostGroup.Ctx(ctx).Order(dao.HostGroup.Columns().Rank).Safe(true)
	if !gutil.IsEmpty(search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, util.SqlLikeStr(search)))
	}
	err = m.Scan(&out)
	return
}

func (*sHostGroup) Upt(ctx context.Context, in *do.HostGroup) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.HostGroup.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}

		if r != nil {
			return gerror.Newf("已存在名称为 %s 主机组", in.Name)
		}
	}
	_, err = dao.HostGroup.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sHostGroup) Del(ctx context.Context, in *do.HostGroup) (err error) {
	_, err = dao.HostGroup.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
