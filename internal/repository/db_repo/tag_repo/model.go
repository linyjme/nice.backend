package tag_repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	UUID    uuid.UUID `json:"_id "gorm:"comment:announcementUUID;column:_id" bson:"_id"`
	Name        string `json:"name" bson:"name" gorm:"type:varchar(128)"`                                  //
	Slug        string `json:"slug" bson:"slug" gorm:"type:varchar(128);not null"`                         //
	Description string `json:"description"  bson:"description" gorm:"type:varchar(256)"`                   //
	IconName    string `json:"name" bson:"icon_name" gorm:"type:varchar(128);column:icon_name"`            //
	IconValue   string `json:"value" bson:"icon_name" gorm:"type:varchar(128);column:icon_value;not null"` //
	Extends     string `json:"extends"  bson:"extends" gorm:"type:varchar(256);column:extends"` //
}

func (t Tags) TableName() string {
	return "tb_tags"
}
