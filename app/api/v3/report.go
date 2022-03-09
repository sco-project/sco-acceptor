/**
    package: sco_tracers
    filename: v3
    author: diogo@gmail.com
    time: 2022/3/1 16:49
**/
package v3

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"sco-acceptor/app/global"
	"sco-acceptor/app/model"
	"sco-acceptor/app/service"
	"sco-acceptor/app/utils"
	"sco-acceptor/library/mq/kafka"
)

// ReportCtl 上报的
type ReportCtl struct{}

// Report 3.0 上报的接口 主要是用于 监听,wx
func (c *ReportCtl) Report(r *ghttp.Request) {
	// 默认 URL上还是带有 token
	// 判断是否为,已申请的项目
	appkey := r.GetString("a")
	found := service.FindProject(appkey)
	if found != nil {
		// 客户端请求中的方法被禁止
		r.Response.WriteStatusExit(502)
	}

	// BUG 当数据正常的话直接返回.及可,不用等后续的操作
	r.Response.WriteStatus(200)
	//fmt.Printf(`%+v`, resblob)
	saveReport(appkey,r)
	//sendAsyncMessage("hk",string(rJSON))

}

// KafkaSyncEr kafka syncer
func sendAsyncMessage(key, value string) error {
	kafka.Producer.SendMsg(key, value)
	return nil
}

// saveReport 处理上报的信息 report
func saveReport(appkey string, r *ghttp.Request) {
	cip := r.GetClientIp()
	if cip == "" {
		cip = "0.0.0.0"
	}
	requestJSON := r.GetBodyString()
	rJSON,_ := utils.Decode(requestJSON)
	j, err := gjson.DecodeToJson(rJSON);
	//glog.Println(rJSON)
	resblob := new(model.Wxmodel)
	resblob.SetFieldsValue()
	resblob.Ip = cip
	resblob.AppKey = appkey

	if err = j.Struct(resblob); err != nil {
		//panic(err)
		glog.Warning(err)
		return
	}
	//g.Dump(gconv.String(resblob))

	//g.Dump(gconv.MapStrStr(resblob))
	if global.AppReportList[appkey] != nil {
		ptrList := global.AppReportList[appkey]
		//fmt.Println(item)
		ptrList.Append(resblob)
	} else {
		arr := garray.NewArray(true)
		arr.Append(resblob)
		global.AppReportList[appkey] = arr
	}

	sendAsyncMessage("hk", gconv.String(resblob))
}

//type KafkaSyncEr struct {
//	Producer *kafka.KafkaProducer
//}
//
//func (kfk *KafkaSyncEr) Write(p []byte) (n int, err error) {
//	logID := bson.NewObjectId().Hex()
//	kfk.Producer.SendMsg(logID, string(p))
//	return 0, err
//}
