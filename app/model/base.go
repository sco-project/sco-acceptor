/**
    package: sco_tracers
    filename: model
    author: diogo@gmail.com
    time: 2021/9/16 11:10
**/
package model

import (
	"github.com/gogf/gf/os/gtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// PublicFields 公共字段 BaseModel to be emmbered to other struct as audit trail perpurse
type PublicFields struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`                                  // 唯一ID
	CreatedAt time.Time          `bson:"CreatedAt" json:"CreatedAt"`                     // 创建时间 ISODate
	UpdatedAt *time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"` // 修改时间 ISODate
	DeletedAt *time.Time         `bson:"DeletedAt,omitempty" json:"DeletedAt,omitempty"` // 删除时间 ISODate
	//Ctime     int64              `bson:"ctime,omitempty" json:"ctime,omitempty"`         // 创建时间 TimestampMilli
}

// SetFieldsValue 设置公共字段值，在插入数据时使用
func (p *PublicFields) SetFieldsValue() {
	//now := time.Now()
	//glog.Println(p.ID.String())
	// FIXME: 这里加8小时
	now := gtime.Now().UTC().Add(8 * time.Hour).Time
	if p.ID.IsZero() {
		p.ID = primitive.NewObjectID()
	}

	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
}

// PublicFieldsInt 设置公共字段值，自定id，在插入数据时使用
type PublicFieldsInt struct {
	ID        int64      `bson:"_id" json:"id,string"`                           // 唯一ID
	CreatedAt time.Time  `bson:"CreatedAt" json:"CreatedAt"`                     // 创建时间 ISODate
	UpdatedAt time.Time  `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"` // 修改时间 ISODate
	DeletedAt *time.Time `bson:"DeletedAt,omitempty" json:"DeletedAt,omitempty"` // 删除时间 ISODate
}

// SetFieldsValue 初始化
func (m *PublicFieldsInt) SetFieldsValue(newID int64) {
	//t := time.Now()
	t := gtime.Now().UTC().Add(8 * time.Hour).Time

	if m.ID == 0 {
		m.ID = newID
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = t
	}
}
