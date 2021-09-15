/**
    package: sco_tracers
    filename: boot
    author: diogo@gmail.com
    time: 2021/9/14 11:14
**/
package boot

import (
	"strings"

	"github.com/asim/mq/go/client"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

type PubOption struct {
	host string
}

var Mqconn client.Client

// 初始化链接
func NewMqCliet(opt PubOption) error {

	Mqconn = client.New(
		client.WithServers(strings.Split(opt.host, ",")...),
	)
	var testStr = "test"
	err := Mqconn.Publish(testStr, []byte(`diogo`))

	if err != nil {
		glog.Println("---mq 链接失败 -----")
		glog.Warning("---请检查mq , 如没用到也是可以的---")
		return err
	}
	glog.Println("---mq 链接成功 -----")
	// 开启订阅
	//go InitSub(Mqconn)
	InitSub(Mqconn)
	return nil
}

// 初始化 mq 订阅
func InitSub(cli client.Client) {

	/**
	当收到信息推送的时候.及时 进行数据处理及存储
	*/
	topic := g.Cfg().GetString("mqueue.subscribe")
	ch, err := cli.Subscribe(topic)

	if err != nil {
		glog.Println(err.Error())
		return
	}

	for {
		select {
		case e := <-ch:
			glog.Println(string(e))
			//saveSub(e)
		}
	}

}

/**
解析并保存
*/
//func saveSub(doc []byte) error {
//	//reportS := service.NewReportService()
//
//	if j, err := gjson.DecodeToJson(doc); err != nil {
//		//panic(err)
//		glog.Println(err.Error())
//	} else {
//		var ctx context.Context
//		var proinfo interface{}
//		etype := j.GetString("t")
//		appkey := j.GetString("token")
//
//		proCli := DBm.Database.Collection("c_project")
//
//		filter := bson.M{
//			"AppKey": appkey,
//		}
//
//		err := proCli.Find(ctx, filter).One(&proinfo)
//
//		glog.Println(etype)
//		if err != nil {
//			glog.Println(err)
//			return err
//		}
//
//		reportCli := DBm.Database.Collection("c_app_" + appkey)
//
//		switch etype {
//		case "error":
//			reqblob := new(model.ErrorMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Println(err)
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "pv":
//			reqblob := new(model.PvMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "api":
//			reqblob := new(model.ApiMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "res":
//			reqblob := new(model.ResourceMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "health":
//			reqblob := new(model.HealthMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "perf":
//			reqblob := new(model.PerfMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "sum":
//			reqblob := new(model.SumMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "avg":
//			reqblob := new(model.AvgMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		case "msg":
//			reqblob := new(model.MsgMsg)
//			err := j.ToStruct(reqblob)
//			if err != nil {
//				glog.Warning(err.Error())
//			}
//			reqblob.SetFieldsValue()
//			_, err = reportCli.InsertOne(ctx, &reqblob)
//		}
//
//		//return nil
//	}
//
//	return nil
//}
