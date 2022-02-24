package announcement_repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	UUID    uuid.UUID `json:"_id "gorm:"comment:announcementUUID;column:_id" bson:"_id"`
	Content string    `json:"content" bson:"content" gorm:"type:varchar(128)"` //
	State   uint8     `json:"state" bson:"state" gorm:"default:1;not null"`    //
}

func (a Announcement) TableName() string {
	return "tb_announcement"
}
