package service

import (
	"errors"
	"gorm.io/gorm"
	"niceBackend/common/global"
	"niceBackend/core/model"
	"niceBackend/pkg"
)

//@author: yjLin
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.NICE_DB.Where("account = ? AND password = ?", u.Account, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.NICE_DB.Where("username = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = pkg.MD5V([]byte(u.Password))
	err = global.NICE_DB.Create(&u).Error
	return err, u
}

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.NICE_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
