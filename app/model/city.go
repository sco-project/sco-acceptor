/**
    package sco_tracers
    filename model
    author diogo@gmail.com
    time 2021/10/8 17:41
**/
package model

// CityModel
type CityModel struct {
	CityCode   string  `bson:"CityCode" json:"CityCode"`
	ID         int     `bson:"ID" json:"ID"`
	Lat        float64 `bson:"Lat" json:"Lat"`
	LevelType  int     `bson:"LevelType" json:"LevelType"`
	Lng        float64 `bson:"Lng" json:"Lng"`
	MergerName string  `bson:"MergerName" json:"MergerName"`
	Name       string  `bson:"Name" json:"Name"`
	ParentId   int     `bson:"ParentId" json:"ParentId"`
	Pinyin     string  `bson:"Pinyin" json:"Pinyin"`
	Scode      string  `bson:"Scode" json:"Scode"`
	ShortName  string  `bson:"ShortName" json:"ShortName"`
	Sname      string  `bson:"Sname" json:"Sname"`
	ZipCode    string  `bson:"ZipCode" json:"ZipCode"`
}
