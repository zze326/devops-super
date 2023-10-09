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

type sCiEnv struct{}

var cols = dao.CiEnv.Columns()

func init() {
	service.RegisterCiEnv(New())
}

func New() *sCiEnv {
	return &sCiEnv{}
}

func (*sCiEnv) Add(ctx context.Context, in *entity.CiEnv) (err error) {
	r, err := dao.CiEnv.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的环境")
	}
	_, err = dao.CiEnv.Ctx(ctx).Insert(in)
	return
}

func (*sCiEnv) Upt(ctx context.Context, in *do.CiEnv) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.CiEnv.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的环境")
		}
	}
	_, err = dao.CiEnv.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sCiEnv) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiEnv], err error) {
	out = &api.PageLstRes[*entity.CiEnv]{}
	m := dao.CiEnv.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()))
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sCiEnv) GetLst(ctx context.Context) (out []*entity.CiEnv, err error) {
	err = dao.CiEnv.Ctx(ctx).OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sCiEnv) Get(ctx context.Context, in *do.CiEnv) (out *entity.CiEnv, err error) {
	err = dao.CiEnv.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sCiEnv) Del(ctx context.Context, in *do.CiEnv) (err error) {
	_, err = dao.CiEnv.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
