/**
    package: sco-acceptor
    filename: model
    author: diogo@gmail.com
    time: 2022/3/2 16:07
**/
package model

import "github.com/gogf/gf/frame/g"

// Wxmodel v3.0 微信小程序 定制的一些参数
type Wxmodel struct {
	PublicFields `bson:",inline"`
	// 上报解析的，可能存在为空的情况
	//Debug int `json:"__debug"`
	AppKey	    string  `json:"app_key" bson:"AppKey"`
	AppType      string `bson:"app_type" json:"app_type"`
	Tv           string `bson:"tv" json:"tv"`
	Model        string `bson:"model" json:"model"`
	Osn          string `bson:"osn" json:"osn"`
	Osv          string `bson:"osv" json:"osv"`
	Nt           string `bson:"nt" json:"nt"`
	Wv           string `bson:"wv" json:"wv"`
	Dsw          int    `bson:"dsw" json:"dsw"`
	Dsh          int    `bson:"dsh" json:"dsh"`
	Wsw          int    `bson:"wsw" json:"wsw"`
	Wsh          int    `bson:"wsh" json:"wsh"`
	Source       int    `bson:"source" json:"source"`
	SourcePath   string `bson:"source_path" json:"source_path"`
	SourceAppID  string `bson:"source_app_id,omitempty" json:"source_app_id,omitempty"`
	SourceParams string `bson:"source_params,omitempty" json:"source_params,omitempty"`
	SourceSrcKey string `bson:"source_src_key,omitempty" json:"source_src_key,omitempty"`
	TrackID      string `bson:"track_id" json:"track_id"`
	Tracktype    string `bson:"tracktype" json:"tracktype"`
	Action       string `bson:"action" json:"action"`

	ProductID   string `bson:"product_id" json:"product_id"`
	ProductName string `bson:"product_name" json:"product_name"`
	ProductNum  string `bson:"product_num" json:"product_num"`
	Price       string `bson:"price" json:"price"`
	PageID      string `bson:"page_id" json:"page_id"`
	PageTitle   string `bson:"page_title,omitempty" json:"page_title"`
	PageURL     string `bson:"page_url,omitempty" json:"page_url"`
	PrevPageURL string `bson:"prev_page_url,omitempty" json:"prev_page_url"`
	PrevPageID  string `bson:"prev_page_id,omitempty" json:"prev_page_id"`
	LastPageID  string `bson:"last_page_id,omitempty" json:"last_page_id"`
	LastPageURL string `bson:"last_page_url,omitempty" json:"last_page_url"`

	Ip    	string `bson:"ip" json:"ip"`
	Provice string `bson:"provice,omitempty" json:"provice"`
	City    string `bson:"city,omitempty" json:"city"`
	Area    string `bson:"area,omitempty" json:"area"`

	OpenID    string `bson:"open_id,omitempty" json:"open_id"`
	UnionID   string `bson:"union_id,omitempty" json:"union_id"`
	Phone     string `json:"phone,omitempty" bson:"phone"`
	UserName  string `bson:"user_name,omitempty" json:"user_name"`
	Timestamp int64  `bson:"timestamp" json:"timestamp"`
}

// ToJsonString 转换成
func (w *Wxmodel) ToJsonString()  g.Map{
	//delete(w,"id")
	return g.Map{}
}