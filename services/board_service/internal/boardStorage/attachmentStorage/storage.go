package attachmentStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Storage interface {
	AddAttachment(input models.AttachmentInput) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInput) (err error)

	GetAttachmentsByTask(input models.TaskInput) (attachments []models.AttachmentOutside, err error)
}