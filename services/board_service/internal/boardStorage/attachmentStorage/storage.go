package attachmentStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Storage interface {
	AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInternal) (err error)

	GetAttachmentsByTask(input models.TaskInput) (attachments []models.AttachmentOutside, err error)
}

