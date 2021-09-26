/**
    package: sco_tracers
    filename: model
    author: diogo@gmail.com
    time: 2021/9/16 11:11
**/
package model


/**
	collection = "c_project"
*/

type ProjectModel struct {
	// 公共字段，id和时间
	PublicFields `bson:",inline"`
	ProjectName  string   `bson:"ProjectName" json:"ProjectName"`           // 项目名称
	CurUrl       string   `bson:"CurUrl,omitempty" json:"CurUrl,omitempty"` // 网站url
	AppKey       string   `bson:"AppKey" json:"AppKey"`                     // 秘钥，唯一
	AppType      string   `bson:"AppType" json:"AppType"`                   // 浏览器：web  微信小程序 ：wx
	AdminUid     int      `bson:"AdminUid" json:"AdminUid"`                 // 应用创建者的 UID
	UserIds      []string `bson:"UserIds"  json:"UserIds"`                  // 应用所属用户 UID
	Status       int      `bson:"Status" json:"Status"`                     //  项目状态 1 为正常  -1,0 为禁止
	/**
	  slow_page_time: { type: Number, default: 5 }, // 页面加载页面阀值  单位：s
	  slow_js_time: { type: Number, default: 2 }, // js慢资源阀值 单位：s
	  slow_css_time: { type: Number, default: 2 }, // 慢加载css资源阀值  单位：S
	  slow_img_time: { type: Number, default: 2 }, // 慢图片加载资源阀值  单位:S
	  slow_ajax_time: { type: Number, default: 2 }, // AJAX加载阀值
	  is_statisi_pages: { type: Number, default: 1 }, // 是否统计页面性能信息  1：是  0：否
	  is_statisi_ajax: { type: Number, default: 1 }, // 是否统计页面Ajax性能资源 1：是  0：否
	  is_statisi_resource: { type: Number, default: 1 }, // 是否统计页面加载资源性能信息 1：是  0：否
	  is_statisi_system: { type: Number, default: 1 }, // 是否存储用户系统信息资源信息 1：是  0：否
	  is_statisi_error: { type: Number, default: 1 }, // 是否上报页面错误信息  1：是  0：否
	*/
	// 是否发送日报  1：是  0：否
	DailyUse int `bson:"DailyUse,omitempty" json:"DailyUse,omitempty"`
	// 日报邮箱列表
	DailyLists []string `bson:"DailyLists,omitempty" json:"DailyLists,omitempty"`
}
