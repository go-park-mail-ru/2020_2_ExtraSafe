package commentStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Storage interface {
	CreateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	DeleteComment(input models.CommentInput) (err error)

	GetCommentsByTask(input models.TaskInput) (comments []models.CommentOutside, err error)
}
