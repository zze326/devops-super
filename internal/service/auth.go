package service

import (
	"context"
	v1 "devops-super/api/public/v1"
	"devops-super/internal/model"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/utility/util"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func CurrentUser(ctx context.Context) (u *model.RequestUser) {
	_ = gjson.New(authService.GetPayload(ctx)).Scan(&u)
	return
}

func init() {
	ctx := context.Background()
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "devops-super",
		Key:             g.Cfg().MustGet(context.Background(), "jwt.secret").Bytes(),
		Timeout:         g.Cfg().MustGet(ctx, "jwt.expire").Duration(),
		MaxRefresh:      g.Cfg().MustGet(ctx, "jwt.expire").Duration(),
		IdentityKey:     "userId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data any) jwt.MapClaims {
	userInfo := data.(*entity.User)
	claims := make(jwt.MapClaims)
	claims["userId"] = userInfo.Id
	claims["username"] = userInfo.Username
	claims["realName"] = userInfo.RealName
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	g.Log().Debug(ctx, message)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r   = g.RequestFromCtx(ctx)
		req *v1.LoginReq
	)
	if err := r.Parse(&req); err != nil {
		return nil, err
	}

	eUser, err := User().Get(ctx, &do.User{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if eUser == nil {
		return nil, gerror.New("用户不存在")
	}

	if !util.ComparePassword(eUser.Password, req.Password) {
		return nil, gerror.New("密码错误")
	}

	if !eUser.Enabled {
		return nil, gerror.New("该用户处于禁用状态")
	}

	return eUser, nil
}
