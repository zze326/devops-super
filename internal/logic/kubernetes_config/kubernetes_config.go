package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sKubernetesConfig struct{}

var cols = dao.KubernetesConfig.Columns()

func init() {
	service.RegisterKubernetesConfig(New())
}

func New() *sKubernetesConfig {
	return &sKubernetesConfig{}
}

func (*sKubernetesConfig) Add(ctx context.Context, in *entity.KubernetesConfig) (err error) {
	r, err := dao.KubernetesConfig.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的配置")
	}
	_, err = dao.KubernetesConfig.Ctx(ctx).Insert(in)
	return
}

func (*sKubernetesConfig) Upt(ctx context.Context, in *do.KubernetesConfig) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.KubernetesConfig.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的配置")
		}
	}
	_, err = dao.KubernetesConfig.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sKubernetesConfig) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.KubernetesConfig], err error) {
	out = &api.PageLstRes[*entity.KubernetesConfig]{}
	m := dao.KubernetesConfig.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()))
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sKubernetesConfig) GetLst(ctx context.Context) (out []*entity.KubernetesConfig, err error) {
	err = dao.KubernetesConfig.Ctx(ctx).OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sKubernetesConfig) GetPartialLst(ctx context.Context) (out []*mid.KubernetesConfigPartial, err error) {
	err = dao.KubernetesConfig.Ctx(ctx).FieldsEx(cols.Config, cols.UpdatedAt).OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sKubernetesConfig) Get(ctx context.Context, in *do.KubernetesConfig) (out *entity.KubernetesConfig, err error) {
	err = dao.KubernetesConfig.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sKubernetesConfig) Del(ctx context.Context, in *do.KubernetesConfig) (err error) {
	_, err = dao.KubernetesConfig.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
