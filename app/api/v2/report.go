/**
    package: sco_tracers
    filename: v2
    author: diogo@gmail.com
    time: 2021/9/18 15:29
**/
package v2

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"sco-acceptor/app/global"
	"sco-acceptor/app/model"
	"sco-acceptor/app/service"
	"strings"
)

// ReportCtl 上报的
type ReportCtl struct{}

// Report 2.0 上报的接口
func (c *ReportCtl) Report(r *ghttp.Request) {
	//cache := cache_service.New()
	//uVal := cache.Get("cacheKEY")
	// cache.Set("cacheKEY","cacheValue",0,"test")
	//formCfg := g.Cfg().GetString("runtimes.testDome")

	//r.Response.Writeln("ReportCtl v2 ",uVal, formCfg)
	//var rawblob interface{}
	//var ctx context.Context

	// 上报的类型
	//mtype := r.Get("t")
	// HACK: 在Nginx 反代的过程中会有可能出错.所以需要用
	cip := r.GetClientIp()
	if cip == "" {
		cip = "0.0.0.0"
	}
	// 判断是否为,已申请的项目
	appkey := r.GetString("token")

	found := service.FindProject(appkey)

	if found != nil {
		// 客户端请求中的方法被禁止
		r.Response.WriteStatusExit(502)
	}
	// 当数据正常的话直接返回.及可,不用等后续的操作
	r.Response.WriteStatus(200)

	// 判断是否有
	if !global.AppNameList.Contains(appkey) {
		//fmt.Println(appkey)
		global.AppNameList.Append(appkey)

	}

	g.Dump(global.AppNameList)

	// ---ip
	ipModel := new(model.IpModel)

	err := r.Parse(&ipModel)
	if err != nil {
		glog.Warning(err.Error())
	}
	//appkey := r.GetString("token")
	deviceInfo := strings.Split(ipModel.Device, "|")
	ipModel.DeviceCode = deviceInfo[0]
	ipModel.AppKey = appkey
	//ipModel.ClientProvinces = results.Region
	//ipModel.ClientCity = results.City
	ipModel.ClientIp = cip
	ipModel.SetFieldsValue()

	if global.AppIpList[appkey] != nil {
		ptrList := global.AppIpList[appkey]
		//fmt.Println(item)
		ptrList.Append(ipModel)
	} else {
		arr := garray.NewArray(true)
		arr.Append(ipModel)
		global.AppIpList[appkey] = arr
	}
	// --- ip global ---

}
