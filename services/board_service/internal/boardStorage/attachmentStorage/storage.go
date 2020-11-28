package attachmentStorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInternal) (err error)

	GetAttachmentsByTask(input models.TaskInput) (attachments []models.AttachmentOutside, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error) {
	err = s.db.QueryRow("INSERT INTO attachments (taskID, filename, filepath) VALUES ($1, $2, $3) RETURNING attachmentID", input.TaskID, input.Filename, input.Filepath).
				Scan(&attachment.AttachmentID)
	if err != nil {
		return attachment, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "AddAttachments"}
	}

	attachment.Filepath = input.Filepath
	attachment.Filename = input.Filename
	return
}

func (s *storage) RemoveAttachment(input models.AttachmentInternal) (err error) {
return
}

func (s *storage) GetAttachmentsByTask(input models.TaskInput) (attachments []models.AttachmentOutside, err error) {
return
}