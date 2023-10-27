package cmd

import (
	"context"
	"devops-super/internal/controller/ci_env"
	"devops-super/internal/controller/ci_pipeline"
	"devops-super/internal/controller/ci_pipeline_run"
	"devops-super/internal/controller/common"
	"devops-super/internal/controller/dept"
	"devops-super/internal/controller/host"
	"devops-super/internal/controller/host_group"
	"devops-super/internal/controller/host_terminal_session"
	"devops-super/internal/controller/kubernetes"
	"devops-super/internal/controller/permission"
	"devops-super/internal/controller/public"
	"devops-super/internal/controller/role"
	"devops-super/internal/controller/secret"
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
						// 系统管理
						user.NewV1(),       // 用户
						permission.NewV1(), // 权限
						role.NewV1(),       // 角色
						dept.NewV1(),       // 部门
						// 资源管理
						host_group.NewV1(),            // 主机组
						host.NewV1(),                  // 主机
						host_terminal_session.NewV1(), // 主机会话
						secret.NewV1(),                // 秘钥管理
						// 持续集成
						ci_env.NewV1(),          // ci 环境
						ci_pipeline.NewV1(),     // ci 流水线
						ci_pipeline_run.NewV1(), // ci 流水线运行记录
						// 其它
						kubernetes.NewV1(), // Kubernetes 相关
						common.NewV1(),     // 公共、通用接口
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
