/**
    package: sco_tracers
    filename: api
    author: diogo@gmail.com
    time: 2021/9/14 15:54
**/
package api

import "github.com/gogf/gf/net/ghttp"

// 上报的
type ReportCtl struct{}

// 2.0 上报的接口
func (c *ReportCtl) Report(r *ghttp.Request) {
	r.Response.Writeln("ReportCtl v2")
}