package cmd

import (
	"context"
	"devops-super/internal/controller/dept"
	"devops-super/internal/controller/permission"
	"devops-super/internal/controller/public"
	"devops-super/internal/controller/role"
	"devops-super/internal/controller/user"
	"devops-super/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					public.NewV1(),
				)

				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						user.NewV1(),
						permission.NewV1(),
						role.NewV1(),
						dept.NewV1(),
					)
				})
			})

			// 初始化
			if err = service.Context().Init(ctx); err != nil {
				g.Log().Fatal(ctx, err)
			}
			s.Run()
			return nil
		},
	}
)
