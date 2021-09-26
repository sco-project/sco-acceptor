/**
    package: sco_tracers
    filename: gpool
    author: diogo@gmail.com
    time: 2021/9/23 14:24
**/
package gpool

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var (
	// gpool 是否开启 批量插入, 默认是true 开启
	gpoolStatus = g.Cfg().GetBool("system.gpoolStatus")
	// gpool 消费者 10s 间隔
	gpoolTimes = g.Cfg().GetDuration("system.gpoolTimes")
	// gpool 消费 100 个
	itemNum = g.Cfg().GetInt("system.itemNum")
)

// 初始化
func init() {

	glog.Println("Start gpool")

	// 动态判断
	if gpoolStatus {
		// ip
		go ipConsumer()
		//go ipProducer()

		// model
		go modelConsumer()
		//go modelProducer()

		// report
		go reportConsumer()
		//go reportProducer()
	}

}
