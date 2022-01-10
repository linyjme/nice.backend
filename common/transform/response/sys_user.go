package response

import (
	"niceBackend/internal/api/repository/db_repo/user_repo"
)

type SysUserResponse struct {
	User user_repo.SysUser `json:"user"`
}

type LoginResponse struct {
	User      user_repo.SysUser `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
