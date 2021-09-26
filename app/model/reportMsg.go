/**
    package: sco_tracers
    filename: model
    author: diogo@gmail.com
    time: 2021/9/16 11:12
**/
package model

/**
collection = "c_app_" + $appkey
每个collection 都立

*/

type MsgType string

// 类型的定义
//const (
//	Tnull     MsgType = ""
//	Terror    MsgType = "error"
//	Tres      MsgType = "res"
//	Tapi      MsgType = "api"
//	Tpv       MsgType = "pv"  pv
//	Thealth   MsgType = "health" 健康度
//	Tperf     MsgType = "perf" 页面性能
//	Tbehavior MsgType = "behavior" 行为
//	Tsum      MsgType = "sum"
//	Tavg      MsgType = "avg"
//	Tpercent  MsgType = "percent"
//	Tmsg      MsgType = "msg" 手动上会以
//	Tduration MsgType = "duration" 用户在线时长统计
//)

// CommonMsg 基础信息 支持 长参数 与简写参数(两参数)
type CommonMsg struct {
	PublicFields `bson:",inline"`
	Utype        MsgType `bson:"type" json:"t"`                             //类型
	Utimes       int     `bson:"times,omitempty" json:"ts,times,omitempty"` // 次数
	Upage        string  `bson:"page" json:"pa,page"`                       // 页面路径
	Uhash        string  `bson:"hash" json:"ha,hash"`                       // 页面hash
	Version      string  `bson:"version" json:"v"`                          // 版本
	Uenv         string  `bson:"env" json:"e"`                              // 开发生产环境
	AppKey       string  `bson:"appkey" json:"to,token"`                    // 项目id
	Udevice      string  `bson:"device" json:"dv,device"`                   // 运行环境平台的信息
	Begin        int     `bson:"begin" json:"be,begin"`                     // 开始时间戳
	Usr          string  `bson:"sr" json:"sr"`                              // 屏幕分辨率
	Up           string  `bson:"vp" json:"vp"`                              // view 分辨率
	Uid          string  `bson:"uid" json:"uid"`                            // user id  sdk 生成的
	Usid         string  `bson:"sid" json:"sid"`                            // session id
	Uct          string  `bson:"ct" json:"ct"`                              // 网络
	Uul          string  `bson:"ul" json:"ul"`                              // 语言
	Uorign       string  `bson:"origin" json:"o"`                           // 原始 url
	Uuser        string  `bson:"userinfo" json:"us,user"`                   // 自定义 的user info
	// 内部信息转换用的
	TraceId  string `bson:"traceId" json:"traceId"`      // trace id
	ClientIP string `bson:"client_ip"  json:"client_ip"` // 请求的IP地址
}

// DurationMsg 用户在线时长统计
type DurationMsg struct {
	CommonMsg  `bson:",inline"`
	DurationMs int `bson:"duration_ms" json:"dms,duration_ms"` // 时间
}

// PvMsg 上报
type PvMsg struct {
	CommonMsg `bson:",inline"`
	Ptitle    string `bson:"pageTitle" json:"dt"` // document title
	Plocation string `bson:"location" json:"dl"`  //document location
	Pregion   string `bson:"region" json:"dr"`    // 来源
	Ppr       int    `bson:"dpr" json:"dpr"`      // dpr
	Pde       string `bson:"de" json:"de"`        // document 编码
}

// ErrorMsg 错误上报
type ErrorMsg struct {
	CommonMsg `bson:",inline"`
	Etype     string `bson:"subtype" json:"st"`                                       // 子类别
	Emsg      string `bson:"msg" json:"msg"`                                          // 信息
	Ecate     string `bson:"cate,omitempty" json:"co,cate,omitempty" p:"cate"`        // 类别
	Edetail   string `bson:"detail,omitempty" json:"det,detail,omitempty" p:"detail"` // 错误栈 或 出错标签
	Efile     string `bson:"file,omitempty" json:"fe,file,omitempty" p:"file"`        // 出错文件
	Eline     int    `bson:"line,omitempty"  json:"l,line,omitempty" p:"line"`        // 行
	Ecol      int    `bson:"col,omitempty" json:"c,col,omitempty" p:"col"`            // 列
}

// ResourceMsg 资源msg
type ResourceMsg struct {
	CommonMsg `bson:",inline"`
	Rdom      int         `bson:"dom" json:"dom" p:"dom"` //   // 所有解析时间 domInteractive - responseEnd
	Rload     int         `bson:"load" json:"load" p:"lo"` // 所有资源加载完时间 loadEventStart- fetchStart
	Rres      interface{} `bson:"res" json:"res" p:"res"`   // 资源信息
}

// ApiMsg Api上报
type ApiMsg struct {
	CommonMsg `bson:",inline"`
	Aurl      string `bson:"ajaxurl" json:"url" p:"au"`     // 接口
	Asuccess  bool   `bson:"success" json:"success" p:"su"` // 成功
	Atime     int    `bson:"time" json:"time" p:"tms"`       // 耗时
	Acode     string `bson:"code" json:"code" p:"ce"`       // http 返回的 FAILED
	Amsg      string `bson:"msg" json:"msg" p:"msg"`         // 信息 与ErrorMsg 不会同时存在
	// FIXME 后续看情况 是否增加下面一些字段
	//method string  // 方法
	//params  string // 参数
	//query string  // query
	//body string   // body
	//response string  // 回参
}

// HealthMsg 健康检查上报
type HealthMsg struct {
	CommonMsg `bson:",inline"`
	Hhealthy  int `bson:"healthy" json:"healthy" p:"hl"`   // 健康？ 0/1
	Hstay     int `bson:"stay" json:"stay" p:"sy"`         // 停留时间
	Herrcount int `bson:"errcount" json:"errcount" p:"er"` // error次数
	Hapisucc  int `bson:"apisucc" json:"apisucc" p:"suc"`   // api成功次数
	Hapifail  int `bson:"apifail" json:"apifail" p:"fai"`   // api错误次数
}

// PerfMsg 页面性能上报
type PerfMsg struct {
	CommonMsg  `bson:",inline"`
	Pdns       int    `bson:"dns" json:"dns" `             // dns时间
	Ptcp       int    `bson:"tcp" json:"tcp"`             // tcp时间
	Pssl       int    `bson:"ssl" json:"ssl"`             // ssl时间
	Pttfb      int    `bson:"ttfb" json:"ttfb" `           // ResponseStart - RequestStart (首包时间，关注网络链路耗时)
	Ptrans     int    `bson:"trans" json:"trans" p:"trs"`         // 重镜像 停留时间
	Pdom       int    `bson:"dom" json:"dom"`             // dom解析时间
	Pres       int    `bson:"res" json:"res"`             // 资源加载 停留时间
	Pfirstbyte int    `bson:"firstbyte" json:"firstbyte" p:"fby"` // 首字节时间
	Pfpt       int    `bson:"fpt" json:"fpt"`             // ResponseEnd - FetchStart （首次渲染时间 / 白屏时间）
	Ptti       int    `bson:"tti" json:"tti"`             // DomInteractive - FetchStart （首次可交付时间）
	Pready     int    `bson:"ready" json:"ready"  p:"rdy"`         // DomContentLoadEventEnd - FetchStart （加载完成时间）
	Pload      int    `bson:"load" json:"load"`           // LoadEventStart - FetchStart （页面完全加载时间）
	Pbandwidth int    `bson:"bandwidth" json:"bandwidth"  p:"bdh"` // 估计的带宽 单位M/s
	Pnavtype   string `bson:"navtype" json:"navtype" p:"nty"`     // nav方式 如reload
}


// SumMsg 统计总量,每次都要传数值过来 可分组 通过 "group::key"
type SumMsg struct {
	CommonMsg `bson:",inline"`
	Group     string `bson:"group" json:"group" p:"gr"` // 类tag 与key
	Ukey      string `bson:"key" json:"key"` // 取的是 ::key
	Uval      int    `bson:"val" json:"val"` // val
}

// AvgMsg  统计平均值  暂没有用到
type AvgMsg struct {
	CommonMsg `bson:",inline"`
	Group     string `bson:"group" json:"group" p:"gr"`
	Ukey      string `bson:"key" json:"key"`
	Uval      int    `bson:"val" json:"val"`
}

// MsgMsg 手动上报的一些信息   "group::msg"
type MsgMsg struct {
	CommonMsg `bson:",inline"`
	Group     string `bson:"group" json:"group" p:"gr"` // 分组key
	Msg       string `bson:"msg" json:"msg"`
}


//属性标签
//我们可以通过 c/gconv/json
//标签来自定义转换后的map键名，当多个标签存在时，
//按照 gconv/c/json 的标签顺序进行优先级识别。

//type User struct {
//	Uid      int
//	Name     string `c:"-"`
//	NickName string `c:"nickname, omitempty"`
//	Pass1    string `c:"password1"`
//	Pass2    string `c:"password2"`
//}

//自定义标签
//此外，我们也可以给struct的属性自定义自己的标签名称，
//并在map转换时通过第二个参数指定标签优先级。

//type User struct {
//	Id   int    `c:"uid"`
//	Name string `my-tag:"nick-name" c:"name"`
//}
//user := &User{
//	Id:   1,
//	Name: "john",
//}
//g.Dump(gconv.Map(user, "my-tag"))
//
//执行后，输出结果为：
//
//{
//"nick-name": "john",
//"uid": 1
//}
