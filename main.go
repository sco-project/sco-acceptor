package main

import (
	_ "sco-acceptor/boot"
	_ "sco-acceptor/router"

	"github.com/gogf/gf/frame/g"
)

func main() {

	g.Server().Run()
}
