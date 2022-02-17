package response

import (
	"niceBackend/internal/repository/db_repo/user_repo"
)

type SysUserResponse struct {
	User user_repo.User `json:"user"`
}

type LoginResponse struct {
	User      user_repo.User `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
