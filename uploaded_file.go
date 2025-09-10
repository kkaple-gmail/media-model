package model

import "github.com/google/uuid"

type UploadedFile struct {
	Base
	OriginalName string `json:"originalName"`
	Size         uint64 `json:"size"`
}

func NewUploadedFile(size uint64) *UploadFile {
	id = uuid.New()
	return &UploadFile{ID: id, Size: size}
}
