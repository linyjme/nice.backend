package user_service

import (
	"errors"
	"gorm.io/gorm"
	"niceBackend/common/global"
	"niceBackend/internal/repository/db_repo/user_repo"
	"niceBackend/pkg"
)

//@author: yjLin
//@function: Login
//@description: 用户登录
//@param: u *model.User
//@return: err error, userInter *model.User

func Login(u *user_repo.User) (err error, userInter *user_repo.User) {
	var user user_repo.User
	err = global.NiceDb.Where("account = ? AND password = ?", u.Account, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func Register(u user_repo.User) (err error, userInter user_repo.User) {
	var user user_repo.User
	if !errors.Is(global.NiceDb.Where("username = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = pkg.MD5V([]byte(u.Password))
	err = global.NiceDb.Create(&u).Error
	return err, u
}

func FindUserByUuid(uuid string) (err error, user *user_repo.User) {
	var u user_repo.User
	if err = global.NiceDb.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
