/**
    package: sco_tracers
    filename: service
    author: diogo@gmail.com
    time: 2021/9/16 16:10
**/
package service

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


var  (
	ErrNotFound = errors.New("not found")
	ErrNotStatus = errors.New("project stop")
	ProjectColl = "c_project"
	cacheTime = g.Cfg().GetDuration("system.gcacheTimes")
)
// ----------------------------------- 公共函数 ----------------------------------------

// CheckUpdateContent 执行更新操作之前先判断有没有$操作符
func CheckUpdateContent(update bson.M) error {
	for k := range update {
		if k[0] != '$' {
			return errors.New("update content must start with '$'")
		}
	}
	return nil
}

// ExcludeDeleted 不包含已删除的
func ExcludeDeleted(selector bson.M) bson.M {
	selector["deletedAt"] = bson.M{"$exists": false}
	return selector
}

// UpdatedTime 更新updatedAt时间
func UpdatedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["updatedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"updatedAt": time.Now()}
	}
	return update
}

// DeletedTime 更新deletedAt时间
func DeletedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["deletedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"deletedAt": time.Now()}
	}
	return update
}