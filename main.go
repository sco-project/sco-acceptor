package main

import (
	_ "sco-acceptor/boot"
	_ "sco-acceptor/router"
	// 开启 pool
	_ "sco-acceptor/app/service/gpool"

	"github.com/gogf/gf/frame/g"
)

func main() {

	serverAgent := g.Cfg().GetString("system.apiServiceName")
	g.Server().SetServerAgent(serverAgent)
	g.Server().Run()
}
