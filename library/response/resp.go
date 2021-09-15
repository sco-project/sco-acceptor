/**
    package: sco_tracers
    filename: response
    author: diogo@gmail.com
    time: 2021/9/14 11:34
**/
package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
标准返回结果数据结构封装。
返回固定数据结构的JSON:
	code:  错误码(200:成功, 201:失败, >200:错误码);
	msg:  请求结果信息;
	data: 请求结果,根据不同接口返回结果的数据结构不同;
*/
func Json(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	redata := &Resp{
		Code: code,
		Msg:  msg,
		Data: responseData,
	}
	//_ = r.Response.WriteJson()

	r.Response.WriteJsonExit(redata)

	//r.Exit()
}

// 返回错误 信息
func Refail(r *ghttp.Request, Code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	// 统一 返回码
	errCode := gconv.String(Code)
	msgInfo := gconv.String(ReturnCode[errCode])
	_ = r.Response.WriteJson(g.Map{
		"code": Code,
		"msg":  msgInfo + msg,
		"data": responseData,
	})
	r.Exit()

}