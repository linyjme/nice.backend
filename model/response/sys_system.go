package response

import "asyncClient/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
