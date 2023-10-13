package user

import (
	"context"
	"devops-super/api"
	"devops-super/internal/consts"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sCiPipeline struct{}

var cols = dao.CiPipeline.Columns()

func init() {
	service.RegisterCiPipeline(New())
}

func New() *sCiPipeline {
	return &sCiPipeline{}
}

func (*sCiPipeline) Add(ctx context.Context, in *entity.CiPipeline) (err error) {
	r, err := dao.CiPipeline.Ctx(ctx).Where(cols.Name, in.Name).One()
	if err != nil {
		return err
	}
	if r != nil {
		return gerror.New("已存在该名称的流水线")
	}
	_, err = dao.CiPipeline.Ctx(ctx).Insert(in)
	return
}

func (*sCiPipeline) Upt(ctx context.Context, in *do.CiPipeline) (err error) {
	if !gutil.IsEmpty(in.Name) {
		r, err := dao.CiPipeline.Ctx(ctx).WhereNot(cols.Id, in.Id).Where(cols.Name, in.Name).One()
		if err != nil {
			return err
		}
		if r != nil {
			return gerror.New("已存在该名称的流水线")
		}
	}
	_, err = dao.CiPipeline.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sCiPipeline) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiPipeline], err error) {
	out = &api.PageLstRes[*entity.CiPipeline]{}
	m := dao.CiPipeline.Ctx(ctx).Safe(true)
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Name, in.SearchStr()))
	}

	err = m.Offset(in.Offset()).Limit(in.Limit()).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sCiPipeline) GetLst(ctx context.Context) (out []*entity.CiPipeline, err error) {
	err = dao.CiPipeline.Ctx(ctx).OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sCiPipeline) Get(ctx context.Context, in *do.CiPipeline) (out *entity.CiPipeline, err error) {
	err = dao.CiPipeline.Ctx(ctx).Where(in).OmitNilWhere().Limit(1).Scan(&out)
	return
}

func (*sCiPipeline) Del(ctx context.Context, in *do.CiPipeline) (err error) {
	_, err = dao.CiPipeline.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}

func (s *sCiPipeline) Run(ctx context.Context, id int) (err error) {
	var (
		ePipeline *entity.CiPipeline
		config    mid.CiPipelineConfig
		envMap    map[int]*entity.CiEnv
	)
	ePipeline, err = s.Get(ctx, &do.CiPipeline{Id: id})
	if err != nil {
		return
	}

	if err = ePipeline.Config.Scan(&config); err != nil {
		return
	}

	if envMap, err = service.CiEnv().GetEntityMap(ctx, config.GetEnvIds()); err != nil {
		return
	}

	// 组装环境关联的镜像和秘钥名称
	for _, envItem := range config {
		envItem.Image = envMap[envItem.Id].Image
		envItem.SecretName = envMap[envItem.Id].SecretName

		for _, stageItem := range envItem.Stages {
			for _, taskItem := range stageItem.Tasks {
				// 如果设置了 git 认证秘钥，组装一下
				if secretId := taskItem.GitPullData.SecretId; secretId > 0 {
					var (
						eSecret      *entity.Secret
						gitBasicAuth *mid.UsernamePasswordContent
					)
					eSecret, err = service.Secret().Get(ctx, &do.Secret{Id: secretId})
					if err != nil {
						return err
					}

					if err = eSecret.Content.Scan(&gitBasicAuth); err != nil {
						return
					}
					if eSecret.Type == consts.SECRET_TYPE_GIT_BASIC_AUTH {
						taskItem.GitPullData.GitBasicAuth = gitBasicAuth
					}
				}
			}
		}

		jsonStr, err := gjson.EncodeString(envItem.Stages)
		if err != nil {
			return err
		}

		fmt.Println(jsonStr)
	}
	return
}
