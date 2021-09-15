/**
    package: sco_tracers
    filename: middleware
    author: diogo@gmail.com
    time: 2021/9/14 11:32
**/
package middleware


import "github.com/gogf/gf/net/ghttp"

// 允许接口跨域请求
func CORS(r *ghttp.Request) {
	//r.Response.CORSDefault()
	corsOptions := r.Response.DefaultCORSOptions()
	// 限制域名
	//corsOptions.AllowDomain = []string{"apiend.com", "localhost"}
	// 限制方法
	corsOptions.AllowMethods = "GET,POST,OPTIONS"
	r.Response.CORS(corsOptions)

	r.Middleware.Next()

}
