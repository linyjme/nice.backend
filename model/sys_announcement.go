package model

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type Announcement struct {
	global.BASE_MODEL
	UUID    uuid.UUID `json:"_id "gorm:"comment:announcementUUID" bson:"_id"`
	Content string    `json:"content" bson:"content" gorm:"type:varchar(128)"` //
	State   uint8     `json:"state" bson:"state" gorm:"default:1;not null"`    //
	V       uint8     `json:"_v" bson:"__v" gorm:"default:0;not null"`
}

func (a Announcement) TableName() string {
	return "sys_announcement"
}
