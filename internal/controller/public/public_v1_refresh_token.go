package public

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"time"

	"devops-super/api/public/v1"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	token, expire := service.Auth().RefreshHandler(ctx)
	refreshAfter := time.Now().Add(expire.Sub(time.Now()) / 2).UnixMilli()
	eUser, err := service.User().GetComb(ctx, &do.User{Id: service.Auth().GetIdentity(ctx)})
	if err != nil {
		return nil, err
	}
	res = &v1.RefreshTokenRes{
		Username:     eUser.Username,
		RealName:     eUser.RealName,
		Token:        token,
		Expire:       expire.UnixMilli(),
		RefreshAfter: refreshAfter,
		Roles:        eUser.RoleCodes(),
	}
	return
}
