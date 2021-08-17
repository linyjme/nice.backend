package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Announcement struct {
	Id        uint32    `json:"id "gorm:"comment:announcementID" bson:"id"`
	UUID      uuid.UUID `json:"_id "gorm:"comment:announcementUUID" bson:"_id"`
	Content   string    `json:"content" bson:"content" gorm:"type:varchar(128)"` //
	State     int8      `json:"state" bson:"state" gorm:"type:varchar(128)"`     //
	CreatedAt time.Time `json:"create_at" bson:"create_at"`                      // 创建时间
	UpdatedAt time.Time `json:"update_at" bson:"update_at"`                      // 更新时间
	V         uint8     `json:"_v" bson:"__v"`
}
