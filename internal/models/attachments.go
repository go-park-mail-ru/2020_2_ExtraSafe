package models

type Attachment struct {
	AttachmentID int64
	TaskID       int64
	Filename     string
}

type AttachmentInput struct {
	BoardID      int64  `json:"-"`
	UserID       int64  `json:"-"`
	TaskID       int64  `json:"taskID"`
	AttachmentID int64  `json:"attachmentID"`
	Filename     string `json:"attachmentFileName"`
	File         []byte `json:"-"`
}

type AttachmentInternal struct {
	TaskID       int64
	AttachmentID int64
	Filename     string
	Filepath     string
}

type AttachmentFileInternal struct {
	UserID   int64
	Filename string
	File     []byte
}

type AttachmentOutside struct {
	AttachmentID int64  `json:"attachmentID"`
	Filename     string `json:"attachmentFileName"`
	Filepath     string `json:"attachmentFilePath"`
}