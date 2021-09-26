/**
    package: sco_tracers
    filename: gpool
    author: diogo@gmail.com
    time: 2021/9/17 16:05
**/
package gpool

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"sco-acceptor/app/global"
	"sco-acceptor/app/service"
	"time"
)

//  reportConsumer 消费者
func reportConsumer() {
	glog.Print("reportConsumer")
	uticker := time.NewTicker(gpoolTimes * time.Second)

	for {
		<-uticker.C
		g.Dump("=== report Consumer ===")
		// 随机从 AppIpList  取 itemNum 个
		for key, array := range global.AppReportList {

			clist := array.PopLefts(itemNum)
			g.Dump("=== report ===")
			fmt.Println(key)
			fmt.Println(len(clist))
			appBag := "c_app_" + key
			// 不为 clist
			if len(clist) != 0 {
				res, err := service.InsertArray(appBag, clist)

				glog.Print(res)
				glog.Print(err)
			}
			//g.Dump(clist)
		}

	}
}

//  reportProducer 生产者
func reportProducer() {

}
