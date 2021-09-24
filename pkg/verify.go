package pkg

var (
	LoginVerify        = Rules{"Account": {Empty()}, "Password": {NotEmpty()}}
	RegisterVerify     = Rules{"Account": {NotEmpty()}, "Password": {NotEmpty()}}
	AnnouncementVerify = Rules{"Content": {NotEmpty()}, "State": {NotEmpty()}}
)
