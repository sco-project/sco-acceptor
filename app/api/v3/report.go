/**
    package: sco_tracers
    filename: v3
    author: diogo@gmail.com
    time: 2022/3/1 16:49
**/
package v3

import (
	"github.com/gogf/gf/net/ghttp"

	"sco-acceptor/app/service"
	"sco-acceptor/library/mq/kafka"
)

// ReportCtl 上报的
type ReportCtl struct{}

// Report 3.0 上报的接口 主要是用于 监听,wx
func (c *ReportCtl) Report(r *ghttp.Request) {

	// 默认 URL上还是带有 token
	// 判断是否为,已申请的项目
	appkey := r.GetString("token")
	found := service.FindProject(appkey)
	if found != nil {
		// 客户端请求中的方法被禁止
		r.Response.WriteStatusExit(502)
	}
	// BUG 当数据正常的话直接返回.及可,不用等后续的操作
	r.Response.WriteStatus(200)
	sendAsyncMessage("name", "diogoxiang")
}

// KafkaSyncEr kafka syncer

func sendAsyncMessage(key, value string) error {
	kafka.Producer.SendMsg(key, value)
	return nil
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
