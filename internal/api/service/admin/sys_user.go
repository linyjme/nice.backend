package admin

import (
	"errors"

	"niceBackend/common/global"
	"niceBackend/internal/repository/db_repo/admin_repo"
	"niceBackend/pkg"

	"gorm.io/gorm"
)

//@author: yjLin
//@function: Login
//@description: 用户登录
//@param: u *model.Admin
//@return: err error, userInter *model.Admin

func Login(u *admin_repo.Admin) (err error, userInter *admin_repo.Admin) {
	var user admin_repo.Admin
	err = global.NiceDb.Where("account = ? AND password = ?", u.Account, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func Register(u admin_repo.Admin) (err error, userInter admin_repo.Admin) {
	var user admin_repo.Admin
	if !errors.Is(global.NiceDb.Where("username = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = pkg.MD5V([]byte(u.Password))
	err = global.NiceDb.Create(&u).Error
	return err, u
}

func FindUserByUuid(uuid string) (err error, user *admin_repo.Admin) {
	var u admin_repo.Admin
	if err = global.NiceDb.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
