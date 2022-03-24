package service

import (
	"errors"
	"niceBackend/config"
	"time"

	"niceBackend/common/global"
	"niceBackend/common/transform/response"

	"gorm.io/gorm"
)

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList response.JwtBlacklist) (err error) {
	err = global.NiceDb.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	err := global.NiceDb.Where("jwt = ?", jwt).First(&response.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.NiceRedis.Get(userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(config.GetConf().JWT.ExpiresTime) * time.Second
	err = global.NiceRedis.Set(userName, jwt, timer).Err()
	return err
}
