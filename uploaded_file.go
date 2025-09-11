package model

import "github.com/google/uuid"

type UploadedFile struct {
	Base
	OriginalName string `json:"originalName"`
	Size         uint64 `json:"size"`
}

func NewUploadedFile(size uint64) *UploadedFile {
	id := uuid.New()
	return &UploadedFile{
		Base: Base{ID: id},
		Size: size,
	}
}
