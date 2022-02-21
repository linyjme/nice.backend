package tag_repo

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type Tags struct {
	global.CustomMODEL
	TagId       uint      `json:"id" bson:"id"  gorm:"primarykey"` // 主键ID
	UUID        uuid.UUID `json:"_id "gorm:"comment:tagsUUID" bson:"_id"`
	Name        string    `json:"name" bson:"name" gorm:"type:varchar(128)"`                     //
	Slug        string    `json:"slug" bson:"slug" gorm:"type:varchar(128);not null"`            // 用户登录名
	Description string    `json:"description"  bson:"description" gorm:"type:varchar(256)"`      //
	IconName    string    `json:"icon_name" bson:"icon_name" gorm:"type:varchar(128)"`           //
	IconValue   string    `json:"icon_value" bson:"icon_name" gorm:"type:varchar(128);not null"` // 用户登录名
	V           uint8     `json:"_v" bson:"__v"`
}

func (t Tags) TableName() string {
	return "tb_tags"
}
