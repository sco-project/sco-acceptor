/**
    package: sco_tracers
    filename: api
    author: diogo@gmail.com
    time: 2021/9/14 15:54
**/
package v1

import (
	"github.com/gogf/gf/net/ghttp"
	"sco-acceptor/app/service/cache_service"
)

// ReportCtl 上报的
type ReportCtl struct{}

// Report 1.0 上报的接口
func (c *ReportCtl) Report(r *ghttp.Request) {
	r.Response.Writeln("ReportCtl v1")
	cache := cache_service.New()

	cache.Set("cacheKEY","cacheValue",0,"test")
}