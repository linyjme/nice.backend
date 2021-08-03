package service

import (
	"niceBackend/common/global"
	"niceBackend/model"
	"niceBackend/utils"
	"errors"
	"gorm.io/gorm"
)

//@author: yjLin
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.NICE_DB.Where("username = ? AND password = ?", u.Account, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func Register(u model.User) (err error, userInter model.User) {
	var user model.User
	if !errors.Is(global.NICE_DB.Where("username = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.NICE_DB.Create(&u).Error
	return err, u
}
