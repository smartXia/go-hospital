package response

import "devops-manage/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
