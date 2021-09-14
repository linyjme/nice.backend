package response

import (
	"niceBackend/common/global"
)

type JwtBlacklist struct {
	global.BaseMODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
