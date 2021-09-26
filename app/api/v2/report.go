/**
    package: sco_tracers
    filename: v2
    author: diogo@gmail.com
    time: 2021/9/18 15:29
**/
package v2

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"sco-acceptor/app/global"
	"sco-acceptor/app/model"
	"sco-acceptor/app/service"
	"sco-acceptor/app/utils"
	"sco-acceptor/library/chinaMap"
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
	mtype := r.GetString("t")
	// HACK: 在Nginx 反代的过程中会有可能出错.所以需要用
	//cip := r.GetClientIp()
	//if cip == "" {
	//	cip = "0.0.0.0"
	//}
	// 判断是否为,已申请的项目
	appkey := r.GetString("token")

	found := service.FindProject(appkey)
	if found != nil {
		// 客户端请求中的方法被禁止
		r.Response.WriteStatusExit(502)
	}
	// BUG 当数据正常的话直接返回.及可,不用等后续的操作
	r.Response.WriteStatus(200)

	// 判断是否有
	if !global.AppNameList.Contains(appkey) {
		//fmt.Println(appkey)
		global.AppNameList.Append(appkey)

	}

	//glog.Println(global.AppNameList)
	//g.Dump(global.AppNameList)

	// 记录一次 用户退出则记录一次
	if mtype == "health" {
		// 并发 写入设备信息数据
		go saveUidAndModel(appkey, r)
		// 并发 写入 ip请求数据
		go saveUidAndIps(appkey, r)
	}

	// 处理
	saveReport(mtype, appkey, r)

	//// ---ip
	//ipModel := new(model.IpModel)
	//
	//err := r.Parse(&ipModel)
	//if err != nil {
	//	glog.Warning(err.Error())
	//}
	////appkey := r.GetString("token")
	//deviceInfo := strings.Split(ipModel.Device, "|")
	//ipModel.DeviceCode = deviceInfo[0]
	//ipModel.AppKey = appkey
	////ipModel.ClientProvinces = results.Region
	////ipModel.ClientCity = results.City
	//ipModel.ClientIp = cip
	//ipModel.SetFieldsValue()
	//
	//if global.AppIpList[appkey] != nil {
	//	ptrList := global.AppIpList[appkey]
	//	//fmt.Println(item)
	//	ptrList.Append(ipModel)
	//} else {
	//	arr := garray.NewArray(true)
	//	arr.Append(ipModel)
	//	global.AppIpList[appkey] = arr
	//}
	//// --- ip global ---

}

// 不同类型分开处理 , 并行

// saveReport 处理上报的信息 report
func saveReport(mtype, appkey string, r *ghttp.Request) {
	var rawblob interface{}
	//mtype := r.Get("t")
	cip := r.GetClientIp()
	if cip == "" {
		cip = "0.0.0.0"
	}
	switch mtype {
	case "error":
		resblob := new(model.ErrorMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()
		rawblob = resblob
	case "pv":
		resblob := new(model.PvMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip

		resblob.SetFieldsValue()

		rawblob = resblob
	case "res":
		resblob := new(model.ResourceMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		// TODO: 还有没body 中的数据没有获取
		resblob.SetFieldsValue()

		rawblob = resblob
	case "api":
		resblob := new(model.ApiMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob
	case "health":
		resblob := new(model.HealthMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob
	case "perf":
		resblob := new(model.PerfMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob

	case "sum":
		resblob := new(model.SumMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob
	case "avg":
		resblob := new(model.AvgMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob
	case "msg":
		resblob := new(model.MsgMsg)
		err := r.Parse(&resblob)
		if err != nil {
			glog.Warning(err.Error())
		}
		resblob.ClientIP = cip
		resblob.SetFieldsValue()

		rawblob = resblob

	}

	if global.AppReportList[appkey] != nil {
		ptrList := global.AppReportList[appkey]
		//fmt.Println(item)
		ptrList.Append(rawblob)
	} else {
		arr := garray.NewArray(true)
		arr.Append(rawblob)
		global.AppReportList[appkey] = arr
	}
	//g.Dump(rawblob)
	//glog.Println(rawblob)

}

/**
saveUidAndModel 处理 手机型号 及 当前型号的手机 访问次数
*/
func saveUidAndModel(appkey string, r *ghttp.Request) {

	deviceMode := new(model.DeviceMode)
	coModel := new(model.CoModel)
	err := r.Parse(&deviceMode)
	if err != nil {
		glog.Warning(err.Error())
	}

	deviceInfo := strings.Split(deviceMode.Device, "|")
	coModel.DeviceCode = deviceInfo[0]
	coModel.DeviceOs = deviceInfo[1]
	coModel.DeviceSr = deviceMode.Usr
	coModel.DeviceVp = deviceMode.Up
	coModel.DeviceDpr = deviceMode.Ppr

	filter := bson.M{
		"device_code": deviceInfo[0],
	}

	update := bson.M{
		operator.Set: &coModel,
		operator.Inc: bson.M{
			"counts": 1,
		},
	}

	err = service.FindOneAndUpdateModel(filter, update)
	if err != nil {
		if gstr.Contains(err.Error(), "duplicate") {
			//response.Refail(r, 40001, "")
			glog.Warning("已经记录的code型号")
		} else {
			//response.Refail(r, 40030, fmt.Sprint(err))
			glog.Warning(err)
		}
	}

}

/**
saveUidAndIps 存储 用户的ip 访问信息 及 解析ip , 并存储用户的一些
*/
func saveUidAndIps(appkey string, r *ghttp.Request) {
	cip := r.GetClientIp()
	if cip == "" {
		cip = "0.0.0.0"
	}

	results, err := utils.Location("", cip)
	if err != nil {
		glog.Println(err)
	}
	ipModel := new(model.IpModel)

	err = r.Parse(&ipModel)
	if err != nil {
		glog.Warning(err.Error())
	}
	//appkey := r.GetString("token")
	deviceInfo := strings.Split(ipModel.Device, "|")
	ipModel.DeviceCode = deviceInfo[0]
	ipModel.AppKey = appkey
	ipModel.ClientProvinces = gconv.String(chinaMap.PinyinMap[results.Region])
	ipModel.ClientCity, _ = service.FindCityName(results.City)
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

}
