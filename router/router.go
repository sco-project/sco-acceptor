package router

import "github.com/gogf/gf/frame/g"

func init() {

	s := g.Server()
	c := g.Config()

	// 配置对象及视图对象配置
	//_ = c.AddPath("config")

	// 公用 及错误码 初始化
	StatusCode(s)

	// 初始化 biz 控制器
	initBiz(s, c)
}
