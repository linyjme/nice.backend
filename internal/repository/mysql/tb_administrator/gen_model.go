package tb_administrator

import "time"

// TbAdministrator
//go:generate gormgen -structs TbAdministrator -input .
type TbAdministrator struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	Uid       string    //
	Account   string    //
	Password  string    //
	NickName  string    //
	HeaderImg string    // 用户头像
	Email     string    //
	Phone     string    //
	Company   string    //
	Address   string    //
	Remark    string    //
	Status    int32     //
}
