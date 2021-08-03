package model

type User struct {
	BASE_MODEL
	ID          uint   `gorm:"primarykey;AUTO_INCREMENT"`                 // 主键ID
	Account     string `json:"account" gorm:"type:varchar(128);not null"` // 用户登录名
	Password    string `json:"-"  gorm:"type:varchar(256)"`               //
	Name        string `json:"name" gorm:"type:varchar(128)"`             //
	Email       string `json:"email" gorm:"type:varchar(256)"`            //
	Phone       string `json:"phone" gorm:"type:varchar(32)"`             //
	Company     string `json:"company" gorm:"type:varchar(256)"`          //
	Address     string `json:"address" gorm:"type:varchar(256)"`          //
	Remark      string `json:"remark" gorm:"type:varchar(512)"`           //
	Status      uint8  `json:"status" gorm:"default:0"`                   //
	AccountType uint8  `json:"accountType" gorm:"default:2;not null"`     //
	Extend_A    string `json:"extend_a"`                                  //
}
