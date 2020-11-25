package tagStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Storage interface {
	CreateTag(input models.TagInput) (tag models.TagOutside, err error)
	UpdateTag(input models.TagInput) (tag models.TagOutside, err error)
	DeleteTag(input models.TagInput) (err error)

	AddTag(input models.TaskTagInput) (err error)
	RemoveTag(input models.TaskTagInput) (err error)

	GetBoardTags(input models.BoardInput) (tags []models.TagOutside, err error)
	GetTaskTags(input models.TaskInput) (tags []models.TagOutside, err error)
}
