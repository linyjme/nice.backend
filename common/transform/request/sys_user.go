package request

// Admin register structure

type Announcements struct {
	Content string `json:"content"` // 用户 账号密码
	State   uint8  `json:"state"`   // 用户 账号密码
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}
