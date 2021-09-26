/**
    package: sco_tracers
    filename: test
    author: diogo@gmail.com
    time: 2021/10/11 9:59
**/
package test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"sco-acceptor/app/service"
	"sco-acceptor/library/response"
)

// TestCtl 测试一些参数及加解密的功能
type TestCtl struct{}

func (T *TestCtl) Test1(r *ghttp.Request) {

	pinyin := r.GetString("pinyin")

	cityName, err := service.FindCityName(pinyin)

	g.Log().Println(err)

	response.Json(r, 200, "done", cityName)

}
