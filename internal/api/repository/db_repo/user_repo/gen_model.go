package user_repo

import (
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/global"
)

type User struct {
	global.BaseMODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Account   string    `json:"account" gorm:"type:varchar(128);not null"`                                     // 用户登录名
	Password  string    `json:"-"  gorm:"type:varchar(256)"`                                                   //
	NickName  string    `json:"nick_name" gorm:"type:varchar(128)"`                                            //
	HeaderImg string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Email     string    `json:"email" gorm:"type:varchar(256)"`                                                //
	Phone     string    `json:"phone" gorm:"type:varchar(32)"`                                                 //
	Company   string    `json:"company" gorm:"type:varchar(256)"`                                              //
	Address   string    `json:"address" gorm:"type:varchar(256)"`                                              //
	Remark    string    `json:"remark" gorm:"type:varchar(512)"`                                               //
	Status    uint8     `json:"status" gorm:"default:0"`                                                       //
}

func (s User) TableName() string {
	return "user"
}
