package model

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type Article struct {
	global.BaseMODEL
	UUID        uuid.UUID `json:"_id "gorm:"comment:tagsUUID" bson:"_id"`
	Password    string    `json:"password" bson:"password" gorm:"type:varchar(64)"`           //
	Keywords    string    `json:"keywords" bson:"keywords" gorm:"type:varchar(128);not null"` // 用户登录名
	State       uint8     `json:"state" bson:"state" gorm:"default:1;not null"`               //
	Public      uint8     `json:"public" bson:"public" gorm:"default:1;not null"`             //
	Origin      uint8     `json:"origin" bson:"origin" gorm:"default:0;not null"`             //
	Tag         Tags      `json:"tag" gorm:"foreignKey:TagId;references:TagId;comment:文字标签"`
	Category    Category  `json:"category" bson:"category" gorm:"foreignKey:CategoryId;references:CategoryId;comment:文章分类"` //
	Description string    `json:"description"  bson:"description" gorm:"type:varchar(256)"`                                 //
	Title       string    `json:"title" bson:"title" gorm:"type:varchar(64)"`                                               //
	Content     string    `json:"content" bson:"content" gorm:"type:varchar(1024)"`                                         //
	Thumb       string    `json:"thumb" bson:"thumb" gorm:"type:varchar(128)"`                                              //
	Meta        uint8     `json:"meta" bson:"meta" gorm:"type:varchar(128)"`                                                //
	Extends     string    `json:"extends" bson:"extends" gorm:"type:varchar(256)"`                                          //
	V           uint8     `json:"_v" bson:"__v"`
}

func (A Article) TableName() string {
	return "articles"
}
