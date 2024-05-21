package response

import "devops-manage/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
