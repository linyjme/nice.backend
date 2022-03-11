package response

import (
	"niceBackend/internal/repository/db_repo/admin_repo"
)

type SysUserResponse struct {
	Admin admin_repo.Admin `json:"user"`
}

type LoginResponse struct {
	Admin     admin_repo.Admin `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}
