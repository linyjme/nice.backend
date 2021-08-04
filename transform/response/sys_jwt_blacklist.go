package response

import (
	"niceBackend/common/global"
)

type JwtBlacklist struct {
	global.BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
