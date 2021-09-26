/**
    package: sco_tracers
    filename: model
    author: diogo@gmail.com
    time: 2021/9/23 15:22
**/
package model


// 设备信息
type DeviceMode struct {
	Device   string `bson:"device" json:"device"`
	Usr      string `bson:"sr" json:"sr"`                // 屏幕分辨率
	Up       string `bson:"vp" json:"vp"`                // view 分辨率
	Ppr      int    `bson:"dpr" json:"dpr"`              // dpr
	ClientIP string `bson:"client_ip"  json:"client_ip"` // 请求的IP地址
}

/**
设备类型及数据
collection: c_model
*/
type CoModel struct {
	DeviceName string `bson:"device_name" json:"device_name"` // 设备名称
	DeviceCode string `bson:"device_code" json:"device_code"` // 设备型号
	DeviceOs   string `bson:"device_os" json:"device_os"`     // 系统型号
	DeviceSr   string `bson:"device_sr" json:"device_sr"`     // 设备屏幕 分辩率
	DeviceVp   string `bson:"device_vp" json:"device_vp"`     // 设备可视分辩率
	DeviceDpr  int    `bson:"device_dpr" json:"device_dpr"`   // Dpr
	//Counts     int    `bson:"counts" json:"counts"`           // counts 次数
}
/**
	IP 地址数据
	collection: c_ip_logs
*/
type IpModel struct {
	PublicFields `bson:",inline"`
	//Ip string `bson:"ip" json:"ip"`
	AppKey          string `bson:"AppKey" json:"AppKey"`                   // 项目 appkey
	Uid             string `bson:"uid" json:"uid"`                         // user id  sdk 生成的
	ClientIp        string `bson:"clientIp" json:"clientIp"`               // ip
	ClientProvinces string `bson:"clientProvinces" json:"clientProvinces"` // 省份
	ClientCity      string `bson:"clientCity" json:"clientCity"`           // 城市
	DeviceCode      string `bson:"device_code" json:"device_code"`         // 设备型号
	Device          string `bson:"-" json:"device"`
	//Counts          int    `bson:"counts" json:"counts"` // counts 次数
}