package attachmentstorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)

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
	err = s.db.QueryRow("INSERT INTO attachments (taskID, filename, filepath) VALUES ($1, $2, $3) RETURNING attachmentID, taskID", input.TaskID, input.Filename, input.Filepath).
		Scan(&attachment.AttachmentID, &attachment.TaskID)
	if err != nil {
		return attachment, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "AddAttachment"}
	}

	attachment.Filepath = input.Filepath
	attachment.Filename = input.Filename
	return
}

func (s *storage) RemoveAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error) {
	err = s.db.QueryRow("DELETE FROM attachments WHERE attachmentID = $1 RETURNING attachmentID, taskID", input.AttachmentID).
		Scan(&attachment.AttachmentID, &attachment.TaskID)
	if err != nil {
		return attachment, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "RemoveAttachment"}
	}
	return
}

func (s *storage) GetAttachmentsByTask(input models.TaskInput) (attachments []models.AttachmentOutside, err error) {
	rows, err := s.db.Query("SELECT attachmentID, filename, filepath FROM attachments WHERE taskID = $1", input.TaskID)
	if err != nil {
		return attachments, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetAttachmentsByTask"}
	}
	defer rows.Close()

	for rows.Next() {
		attachment := models.AttachmentOutside{}

		err = rows.Scan(&attachment.AttachmentID, &attachment.Filename, &attachment.Filepath)
		if err != nil {
			return attachments, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetAttachmentsByTask"}
		}

		attachments = append(attachments, attachment)
	}
	return
}
