package response

import "devops-manage/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}
