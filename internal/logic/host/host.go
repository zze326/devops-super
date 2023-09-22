package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sHost struct{}

var cols = dao.Host.Columns()

func init() {
	service.RegisterHost(New())
}

func New() *sHost {
	return &sHost{}
}

func (*sHost) Add(ctx context.Context, in *entity.Host) (err error) {
	r, err := dao.Host.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的主机")
	}
	_, err = dao.Host.Ctx(ctx).Insert(in)
	return
}

func (*sHost) Upt(ctx context.Context, in *do.Host) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.Host.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的主机")
		}
	}
	_, err = dao.Host.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sHost) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Host], err error) {
	out = &api.PageLstRes[*entity.Host]{}
	m := dao.Host.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()).WhereOrLike(cols.Host, in.SearchStr()))
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sHost) Get(ctx context.Context, in *do.Host) (out *entity.Host, err error) {
	err = dao.Host.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sHost) Del(ctx context.Context, in *do.Host) (err error) {
	_, err = dao.Host.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
