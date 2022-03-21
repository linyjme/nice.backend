package admin_repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UUID      uuid.UUID `json:"_id "gorm:"comment:;column:uid" bson:"uid"`
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

func (Admin) TableName() string {
	return "tb_administrator"
}
