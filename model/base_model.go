package model

import (
	"time"
)

type BASE_MODEL struct {
	CreatedAt time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdatedAt time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}
