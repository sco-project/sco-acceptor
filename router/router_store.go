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
	"sco-acceptor/app/api/v1"
	v2 "sco-acceptor/app/api/v2"

	"sco-acceptor/library/middleware"
)

func initBiz(s *ghttp.Server, c *gcfg.Config) {

	Environment := c.GetString("system.Environment")
	maxLimiter := c.GetInt("system.maxLimiter")
	backGlobal := c.GetString("system.backGlobal")

	// ip 限流
	cfg := middleware.LimiterConfig{
		Environment: Environment,
		// seconds
		Timeout: 60,
		Max:     maxLimiter,
	}

	// v1 版本
	s.Group(backGlobal+"/v1", func(group *ghttp.RouterGroup) {
		w := new(v1.ReportCtl)
		group.Middleware(middleware.CORS)
		group.Middleware(middleware.NewLimiter(cfg))
		group.ALL("/{.method}", w)
	})

	// v2 版本 2021-9-18 15:30:41
	s.Group(backGlobal+"/v2", func(group *ghttp.RouterGroup) {
		w := new(v2.ReportCtl)
		group.Middleware(middleware.CORS)
		group.Middleware(middleware.NewLimiter(cfg))
		group.ALL("/{.method}", w)
	})
}
