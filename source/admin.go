package source

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"niceBackend/common/global"
	"niceBackend/internal/api/repository/db_repo/user_repo"
	"time"
)

var Admin = new(admin)

type admin struct{}

var admins = []user_repo.User{
	{BaseMODEL: global.BaseMODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Account: "admin", Password: "cm9vdA==", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/NICE_header.jpg"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.NiceDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]user_repo.User{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			global.NiceLog.Error("migrate table failed", zap.Any("err", err))
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
