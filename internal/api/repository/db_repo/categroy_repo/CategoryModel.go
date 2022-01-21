package categroy_repo

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type Category struct {
	global.CustomMODEL
	UUID        uuid.UUID `json:"_id "gorm:"comment:tagsUUID" bson:"_id"`
	CategoryId  uint      `json:"category_id" bson:"category_id"  gorm:"primarykey"`        // 主键ID
	Pid         uint32    `json:"pid" bson:"pid" gorm:""`                                   //
	Name        string    `json:"name" bson:"name" gorm:"type:varchar(64)"`                 //
	Slug        string    `json:"slug" bson:"slug" gorm:"type:varchar(128);not null"`       // 用户登录名
	Description string    `json:"description"  bson:"description" gorm:"type:varchar(256)"` //
	IconName    string    `json:"icon_name" bson:"icon_name" gorm:"type:varchar(64)"`       //
	IconValue   string    `json:"icon_value" bson:"icon_name" gorm:"type:varchar(64)"`      //
	Extends     string    `json:"extends" bson:"extends" gorm:"type:varchar(256)"`          //
	V           uint8     `json:"_v" bson:"__v"`
}

func (C Category) TableName() string {
	return "categories"
}
