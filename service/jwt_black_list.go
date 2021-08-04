package service

import (
	"errors"
	"gorm.io/gorm"
	"niceBackend/common/global"
	"niceBackend/transform/response"
	"time"
)

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList response.JwtBlacklist) (err error) {
	err = global.NICE_DB.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	err := global.NICE_DB.Where("jwt = ?", jwt).First(&response.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.NICE_REDIS.Get(userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.NICE_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.NICE_REDIS.Set(userName, jwt, timer).Err()
	return err
}
