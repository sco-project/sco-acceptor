package main

import (
	"flag"
	"fmt"

	_ "sco-acceptor/boot"
	_ "sco-acceptor/router"
	// 开启 pool
	_ "sco-acceptor/app/service/gpool"

	"github.com/gogf/gf/frame/g"
)

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

// go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"
func main() {
	g.Log().Println("---当前的版本号---")
	fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
	//go ghttp.StartPProfServer(8199)

	// 这里加入启动的参数, 方便多应用启动
	var port int
	flag.IntVar(&port, "port", 9003, "端口号，默认为9003")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	g.Server().SetPort(port)
	//g.Cfg().SetFileName("storage.config.toml")

	serverAgent := g.Cfg().GetString("system.apiServiceName")
	g.Server().SetServerAgent(serverAgent)
	g.Server().Run()
}
