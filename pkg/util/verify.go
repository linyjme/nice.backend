package util

var (
	LoginVerify        = Rules{"account": {Empty()}, "password": {NotEmpty()}}
	AnnouncementVerify = Rules{"Content": {NotEmpty()}, "State": {NotEmpty()}}
)
