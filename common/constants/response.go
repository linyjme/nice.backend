package constants

var ResultCode = map[int]string{
	200:   "OK",
	201:   "获取管理员信息成功",
	4000:  "Parameter error",
	4001:  "Bad required",
	4002:  "用户名不存在或者密码错误",
	5001:  "Admin server abnormal",
	10086: "invalid code",
	1000:  "Unknow exception"}
