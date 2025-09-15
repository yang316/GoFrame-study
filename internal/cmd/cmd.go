package cmd

import (
	"context"
	"gf_study/internal/controller/v1/goods"
	"gf_study/internal/controller/v1/user"

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
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 绑定控制器到路由组
				group.Bind(
					// 用户控制器
					user.NewUser(),
					goods.NewGoods(),
					
				)
			})
			s.AddStaticPath("/uploads", "./resource/uploads")
			s.Run()
			return nil
		},
	}
)
