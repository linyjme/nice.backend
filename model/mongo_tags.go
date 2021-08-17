package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Extends struct {
	Name  string `json:"name" bson:"name" gorm:"type:varchar(128)"`            //
	Value string `json:"value" bson:"value" gorm:"type:varchar(128);not null"` // 用户登录名
}

type Tags struct {
	UUID        uuid.UUID `json:"_id "gorm:"comment:tagsUUID" bson:"_id"`
	Id          uint32    `json:"id "gorm:"comment:tagsID" bson:"id"`
	Name        string    `json:"name" bson:"name" gorm:"type:varchar(128)"`                //
	Slug        string    `json:"slug" bson:"slug" gorm:"type:varchar(128);not null"`       // 用户登录名
	Description string    `json:"description"  bson:"description" gorm:"type:varchar(256)"` //
	Extends     Extends   `json:"extends" bson:"extends"`
	CreatedAt   time.Time `json:"create_at" bson:"create_at"` // 创建时间
	UpdatedAt   time.Time `json:"update_at" bson:"update_at"` // 更新时间
	V           uint8     `json:"_v" bson:"__v"`
}
