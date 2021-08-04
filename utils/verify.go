package utils

var (
	LoginVerify    = Rules{"Account": {Empty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"Account": {NotEmpty()}, "Password": {NotEmpty()}}
)
