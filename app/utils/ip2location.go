/**
    package: sco_tracers
    filename: utils
    author: diogo@gmail.com
    time: 2021/9/14 11:37
**/
package utils

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"os"


	"github.com/ip2location/ip2location-go"
)

var (
	c = g.Cfg().SetFileName("storage.config.toml")
	//tokenDB      = c.GetString("setting.localCacheDB")
	iplocationDB = c.GetString("system.iplocationDB")
	// 缓存一下IP
	Ipcache = gcache.New()
)

type IpRecord struct {
	Country_short string
	Country_long  string
	Region        string
	City          string
	Latitude      float32
	Longitude     float32
}

// Location return lat and long
func Location(file, ip string) (*IpRecord, error) {

	cacheRes := new(IpRecord)
	// 先取缓存
	res, _ := Ipcache.Get(ip)
	if res != nil {
		if err := gconv.Struct(res, &cacheRes); err != nil {
			glog.Warning(err)
			return nil, err
		}
		return cacheRes, nil
		//return res, err
	}

	// 当为空的时候
	if file == "" {
		file = iplocationDB
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return cacheRes, err
	}

	ipdb, _ := ip2location.OpenDB(file)
	record, _ := ipdb.Get_all(ip)
	//cacheRes.Country_long,_= ipdb.Get_country_long(ip)
	cacheRes.Country_long = record.Country_long
	cacheRes.Country_short = record.Country_short
	cacheRes.Region = record.Region
	cacheRes.City = record.City
	cacheRes.Latitude = record.Latitude
	cacheRes.Longitude = record.Longitude

	Ipcache.SetIfNotExist(ip, cacheRes, g.Cfg().GetDuration("system.gcacheTimes")*gtime.M)

	//ress, _ :=Ipcache.Get(ip)
	//glog.Println(ress)
	//ip2location.Close()
	ipdb.Close()

	return cacheRes, nil
}
