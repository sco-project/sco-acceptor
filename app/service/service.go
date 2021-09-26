/**
    package: sco_tracers
    filename: service
    author: diogo@gmail.com
    time: 2021/9/16 16:09
**/
package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"sco-acceptor/app/dao"
	"sco-acceptor/app/model"
	"sco-acceptor/app/service/cache_service"
	"time"
)

// FindProject 查询project 项目
func FindProject(appkey string) error{
	//var project model.ProjectModel
	project := &model.ProjectModel{}
	cache := cache_service.New()

	// 取缓存中的数据
	fond := cache.Get(appkey)
	if fond == 1 {
		return  nil
	}

	// selector
	selector := bson.M{
		"AppKey": appkey,
		//"AdminUid": req.AdminUid,
	}
	// fields
	fields := bson.M{}

	err := dao.FindOne(ProjectColl,selector,fields,project)

	if err != nil{
		return  err
	}

	// 当项目未禁用的时候
	if project.Status  == 1  {
		// 存
		cache.SetIfNotExist(appkey,1,cacheTime*time.Minute,"project")
	}else{
		// 如果是-1,0了,则
		cache.SetIfNotExist(appkey,0,cacheTime*time.Minute,"project")
		return  ErrNotStatus
	}


	return  nil

}


// InsertOne 
func InsertOne(collection string, doc interface{})  error {

	return nil
}

// InsertMany 批量插入
func InsertArray(collection string, doc []interface{}) (result interface{}, err error){

	result ,err  = dao.InsertArray(collection,doc)

	return
}