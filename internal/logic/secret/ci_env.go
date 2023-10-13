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

type sSecret struct{}

var cols = dao.Secret.Columns()

func init() {
	service.RegisterSecret(New())
}

func New() *sSecret {
	return &sSecret{}
}

func (*sSecret) Add(ctx context.Context, in *entity.Secret) (err error) {
	r, err := dao.Secret.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的秘钥")
	}
	_, err = dao.Secret.Ctx(ctx).Insert(in)
	return
}

func (*sSecret) Upt(ctx context.Context, in *do.Secret) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.Secret.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的秘钥")
		}
	}
	_, err = dao.Secret.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sSecret) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Secret], err error) {
	out = &api.PageLstRes[*entity.Secret]{}
	m := dao.Secret.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()))
	}

	if typeV := in.Wheres.Get("type"); !typeV.IsNil() {
		m = m.Where(cols.Type, typeV.Int())
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sSecret) GetLst(ctx context.Context, in *do.Secret) (out []*entity.Secret, err error) {
	err = dao.Secret.Ctx(ctx).Fields(cols.Id, cols.Name).Where(in).OmitNilWhere().OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sSecret) Get(ctx context.Context, in *do.Secret) (out *entity.Secret, err error) {
	err = dao.Secret.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sSecret) Del(ctx context.Context, in *do.Secret) (err error) {
	_, err = dao.Secret.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
