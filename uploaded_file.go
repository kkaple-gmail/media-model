package model

type UploadedFile struct {
	Base
	OriginalName string `json:"originalName"`
	Size         uint64 `json:"size"`
}
