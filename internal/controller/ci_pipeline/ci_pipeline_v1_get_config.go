package ci_pipeline

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline/v1"
)

func (c *ControllerV1) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	var (
		ePipeline *entity.CiPipeline
		config    mid.CiPipelineConfig
	)
	res = new(v1.GetConfigRes)
	ePipeline, err = service.CiPipeline().Get(ctx, &do.CiPipeline{Id: req.Id})
	if err != nil {
		return
	}
	res.Config = ePipeline.Config

	if err = ePipeline.Config.Scan(&config); err != nil {
		return
	}

	res.EnvMap, err = service.CiEnv().GetIdNameMap(ctx, config.GetEnvIds())
	return
}
