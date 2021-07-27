package service

import (
	"asyncClient/global"
	"asyncClient/model"
	"asyncClient/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.RAY_DB.Where("username = ? AND password = ?", u.Account, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}
