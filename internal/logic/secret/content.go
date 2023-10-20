package user

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/model/do"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sSecret) GetKubernetesConfig(ctx context.Context, in *do.Secret) (*mid.TextContent, error) {
	e, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if e.Type != consts.SECRET_TYPE_KUBERNETES_CONFIG {
		return nil, gerror.New("秘钥类型错误")
	}

	result := new(mid.TextContent)
	if err = e.Content.Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}
