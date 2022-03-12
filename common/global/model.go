package global

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseMODEL struct {
	gorm.Model
	ID       uint      `json:"id" bson:"id"  gorm:"primarykey"`
	UUID     uuid.UUID `json:"_id "gorm:"comment:announcementUUID;column:_id" bson:"_id"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at;not null"` // 创建时间
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at;not null"` // 更新时间
}
