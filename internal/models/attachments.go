package models

type Attachment struct {
	AttachmentID int64
	TaskID int64
	Filename string
}

type AttachmentInput struct {
	TaskID int64
	Filename string
	Filepath string
}

type AttachmentFileInput struct {
	UserID   int64
	Filename string
	File     []byte
}

type AttachmentOutside struct {
	AttachmentID int64
	Filename string
}