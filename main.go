package main

import (
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

	serverAgent := g.Cfg().GetString("system.apiServiceName")
	g.Server().SetServerAgent(serverAgent)
	g.Server().Run()
}
