package context

import (
	"context"
	"devops-super/internal/model"
	"devops-super/internal/service"
)

type sContext struct {
	svcCtx *model.ServiceContext
	ctx    context.Context
}

func init() {
	service.RegisterContext(New())
}

func New() *sContext {
	return &sContext{}
}

func (s *sContext) Init(ctx context.Context) error {
	s.ctx = ctx
	return s.initServiceContext()
}

func (s *sContext) initServiceContext() (err error) {
	s.svcCtx = new(model.ServiceContext)
	// 初始化 casbin
	err = s.RefreshCasbin(s.ctx)
	if err != nil {
		return
	}
	return
}

func (s *sContext) RefreshCasbin(ctx context.Context) error {
	enforcer, err := s.initCasbin(ctx)
	if err != nil {
		return err
	}
	s.svcCtx.CasbinEnforcer = enforcer
	return nil
}

func (s *sContext) Ctx() *model.ServiceContext {
	return s.svcCtx
}
