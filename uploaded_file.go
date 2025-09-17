package model

import "github.com/google/uuid"

type UploadedFileState string

const (
	New        UploadedFileState = "new"
	Duplicate  UploadedFileState = "duplicate"
	Processing UploadedFileState = "processing"
	Processed  UploadedFileState = "processed"
	Error      UploadedFileState = "error"
)

type UploadedFile struct {
	Base
	OriginalName string            `json:"original_name"`
	Size         uint64            `json:"size"`
	MIME         string            `json:"mime"`
	SHA256       []byte            `json:"sha256"`
	State        UploadedFileState `json:"state"`
}

func NewUploadedFile(
	origName string,
	size uint64,
	mime string,
	sha256 []byte) *UploadedFile {
	id := uuid.New()
	return &UploadedFile{
		Base:         Base{ID: id},
		OriginalName: origName,
		Size:         size,
		MIME:         mime,
		SHA256:       sha256,
	}
}
