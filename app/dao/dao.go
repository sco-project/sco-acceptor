/**
    package: sco_tracers
    filename: dao
    author: diogo@gmail.com
    time: 2021/9/17 9:57
**/
package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/text/gstr"
	"go.mongodb.org/mongo-driver/bson"
	"sco-acceptor/boot"
)

// Insert 插入一条新数据
func Insert(collection string, doc interface{}) (err error) {
	ctx := context.Background()
	coll := boot.DBm.Database.Collection(collection)

	_, err = coll.InsertOne(ctx, doc)
	if err != nil {
		if gstr.Contains(err.Error(), "E1100") {
			return errors.New("duplicate key")
		} else {
			return err
		}
	}

	return nil
}

// InsertArray 插入批量数据
func InsertArray(collection string, doc []interface{}) (result interface{}, err error) {
	ctx := context.Background()
	coll := boot.DBm.Database.Collection(collection)

	result, err = coll.InsertMany(ctx, doc)

	if err != nil {
		if gstr.Contains(err.Error(), "E1100") {
			return 0, errors.New("duplicate key")
		} else {
			return 0, err
		}
	}

	return result, nil
}

// FindOne 查找一个  有固定参数.可用.否则用 All, res => 传指针
func FindOne(collection string,selector bson.M, fields bson.M, res interface{}) (err error) {
	ctx := context.Background()
	coll := boot.DBm.Database.Collection(collection)

	err = coll.Find(ctx, ExcludeDeleted(selector)).Select(fields).One(res)

	return err
}
