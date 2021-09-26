/**
    package: sco_tracers
    filename: global
    author: diogo@gmail.com
    time: 2021/9/23 14:03
**/
package global

import "github.com/gogf/gf/container/garray"

var (
	// AppNameList   name
	AppNameList = garray.NewArray(true)
	// AppReportList   report list
	AppReportList = map[string]*garray.Array{}
	// AppIpList ip 记录
	AppIpList =  map[string]*garray.Array{}
	// AppModelList  model list
	AppModelList =map[string]*garray.Array{}
)
