package middleware

import (
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (*sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	var (
		requestPath                   = r.Request.URL.Path
		method                        = r.Request.Method
		refreshPermissionPathPrefixes = []string{"/permission", "/role", "/user"}
	)

	pass, err := service.Context().Ctx().CasbinEnforcer.Enforce(service.CurrentUser(r.GetCtx()).Username, requestPath, method)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusForbidden,
			"message": err,
		})
		r.ExitAll()
	}

	if !pass {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusForbidden,
			"message": "没有接口权限",
		})
		r.ExitAll()
	}

	r.Middleware.Next()

	if method != "GET" {
		for _, prefix := range refreshPermissionPathPrefixes {
			if gstr.HasPrefix(requestPath, prefix) {
				if err := service.Context().RefreshCasbin(r.GetCtx()); err != nil {
					r.Response.WriteJson(g.Map{
						"code":    http.StatusInternalServerError,
						"message": gerror.Newf("刷新权限失败: %v", err),
					})
					r.ExitAll()
				}
				break
			}
		}
	}

}
