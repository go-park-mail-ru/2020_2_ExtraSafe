package models

type Attachment struct {
	AttachmentID int64
	TaskID int64
	Filename string
}

//TODO AttachmentInput

type AttachmentOutside struct {
	AttachmentID int64
	Filename string
}