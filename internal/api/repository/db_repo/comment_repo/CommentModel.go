package comment_repo

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type Comments struct {
	global.BaseMODEL
	UUID       uuid.UUID `json:"_id "gorm:"comment:tagsUUID" bson:"_id"`
	Pid        uint32    `json:"pid" bson:"pid" gorm:""`                                 //
	State      uint8     `json:"state" bson:"state" gorm:"default:1;not null"`           //
	IsTop      bool      `json:"is_top" bson:"is_top" gorm:"default:false"`              //
	Slug       string    `json:"slug" bson:"slug" gorm:"type:varchar(128);not null"`     // 用户登录名
	Likes      uint8     `json:"state" bson:"state" gorm:"default:0;not null"`           //
	PostId     uint8     `json:"post_id" bson:"post_id" gorm:""`                         //
	Content    string    `json:"content"  bson:"content" gorm:"type:varchar(256)"`       //
	Agent      string    `json:"agent" bson:"agent" gorm:"type:varchar(64)"`             //
	Author     string    `json:"author" bson:"author" gorm:"type:varchar(64)"`           //
	Ip         string    `json:"ip" bson:"ip" gorm:"type:varchar(32)"`                   //
	IpLocation string    `json:"ip_location" bson:"ip_location" gorm:"type:varchar(64)"` //
	Extends    string    `json:"extends" bson:"extends" gorm:"type:varchar(256)"`        // 用户登录名
	V          uint8     `json:"_v" bson:"__v"`
}

func (C Comments) TableName() string {
	return "comments"
}
