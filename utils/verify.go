package utils

var (
	LoginVerify    = Rules{"Account": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"Account": {NotEmpty()}, "Password": {NotEmpty()}}
)
