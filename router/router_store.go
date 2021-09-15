/**
    package: sco_tracers
    filename: router
    author: diogo@gmail.com
    time: 2021/9/14 11:31
**/
package router

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcfg"
	"sco-acceptor/app/api"
	"sco-acceptor/library/middleware"
)

func initBiz(s *ghttp.Server, c *gcfg.Config) {

	maxLimiter := c.GetInt("system.maxLimiter")
	// ip 限流
	cfg := middleware.LimiterConfig{
		// seconds
		Timeout: 60,
		Max:     maxLimiter,
	}

	backGlobal := c.GetString("system.backGlobal")

	// 其他类型
	s.Group(backGlobal+"/v1", func(group *ghttp.RouterGroup) {
		w := new(api.ReportCtl)
		group.Middleware(middleware.CORS)
		group.Middleware(middleware.NewLimiter(cfg))
		group.ALL("/{.method}", w)
	})
}

