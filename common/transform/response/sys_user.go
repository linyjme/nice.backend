package response

import (
	"niceBackend/internal/repository/db_repo/user_repo"
)

type SysUserResponse struct {
	Admin user_repo.Admin `json:"user"`
}

type LoginResponse struct {
	Admin      user_repo.Admin `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
