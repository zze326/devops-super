package ci_pipeline

import (
	"context"
	"devops-super/api/ci_pipeline/v1"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	var (
		ePipeline *entity.CiPipeline
		config    mid.CiPipelineConfig
		entityMap map[int]*entity.CiEnv
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

	entityMap, err = service.CiEnv().GetEntityMap(ctx, config.GetEnvIds())
	for id, e := range entityMap {
		res.EnvMap.Set(id, g.Map{"name": e.Name, "isKaniko": e.IsKaniko})
	}

	return
}
