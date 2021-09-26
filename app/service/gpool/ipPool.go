/**
    package: sco_tracers
    filename: gpool
    author: diogo@gmail.com
    time: 2021/9/17 16:05
**/
package gpool

import (
	"fmt"
	"sco-acceptor/app/service"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"

	"sco-acceptor/app/global"
)

// ip collection
var ipColl = "c_ip_logs"

// IpConsumer IP 消费者
func ipConsumer() {
	glog.Print("ipConsumer")

	uticker := time.NewTicker(gpoolTimes * time.Second)
	for {
		<-uticker.C
		g.Dump("=== IP Consumer ===")
		// 随机从 AppIpList  取 itemNum 个
		for key, array := range global.AppIpList {

			clist := array.PopLefts(itemNum)
			g.Dump("=== IP ===")
			fmt.Println(key)
			fmt.Println(len(clist))
			//appBag := "c_app_" + key
			// 不为 clist
			if len(clist) != 0 {
				res,err :=service.InsertArray(ipColl, clist)

				glog.Print(res)
				glog.Print(err)
			}
			//g.Dump(clist)
		}


	}

}

// IpProducer IP 生产者
func ipProducer() {

}
