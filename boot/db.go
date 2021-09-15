/**
    package: sco_tracers
    filename: boot
    author: diogo@gmail.com
    time: 2021/9/14 11:12
**/
package boot

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/qiniu/qmgo"
)

var (
	//DBclient    *qmgo.Client
	DBm *qmgo.QmgoClient
	err error
)

// mongo config
//type qmgoconfig struct {
//	Uri              string
//	Database         string
//	Coll             string
//	ConnectTimeoutMS int64
//	SocketTimeoutMS  int64
//	MaxPoolSize      int64
//	MinPoolSize      int64
//}

/**
初始化链接 mongo
*/
func initClient(cfg qmgo.Config) error {
	ctx := context.Background()
	//DBclient, err = qmgo.NewClient(ctx, &cfg)

	DBm, err = qmgo.Open(ctx, &cfg)

	if err != nil {
		glog.Println(err.Error())
		glog.Warning("----mongo 链接失败----检测 mongodb -")
		g.Log().Printf("----请检测链接是否正常 %s -", cfg.Uri)
		return err
	}
	glog.Println("----mongo 链接成功--")
	return nil
}
