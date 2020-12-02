package tagStorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateTag(input models.TagInput) (tag models.TagOutside, err error)
	UpdateTag(input models.TagInput) (tag models.TagOutside, err error)
	DeleteTag(input models.TagInput) (err error)

	AddTag(input models.TaskTagInput) (err error)
	RemoveTag(input models.TaskTagInput) (err error)

	GetBoardTags(input models.BoardInput) (tags []models.TagOutside, err error)
	GetTaskTags(input models.TaskInput) (tags []models.TagOutside, err error)
}
type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateTag(input models.TagInput) (tag models.TagOutside, err error) {
	err = s.db.QueryRow("INSERT INTO tags (boardID, name, color) VALUES ($1, $2, $3) RETURNING tagID", input.BoardID, input.Name, input.Color).
				Scan(&tag.TagID)
	if err != nil {
		return tag, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateTag"}
	}

	tag.Color = input.Color
	tag.Name = input.Name
	return
}

func (s *storage) UpdateTag(input models.TagInput) (tag models.TagOutside, err error) {
	_, err = s.db.Exec("UPDATE tags SET name = $1, color = $2 WHERE tagID = $3", input.Name, input.Color, input.TagID)
	if err != nil {
		return tag, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "UpdateTag"}
	}

	tag.TagID = input.TagID
	tag.Color = input.Color
	tag.Name = input.Name
	return
}
func (s *storage) DeleteTag(input models.TagInput) (err error) {
	_, err = s.db.Exec("DELETE FROM tags WHERE tagID = $1", input.TagID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "UpdateTag"}
	}

	return
}

// assign tag to task
func (s *storage) AddTag(input models.TaskTagInput) (err error) {
	_, err = s.db.Exec("INSERT INTO task_tags (taskID, tagID) VALUES ($1, $2)", input.TaskID, input.TagID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "AddTag"}
	}
	return
}

func (s *storage) RemoveTag(input models.TaskTagInput) (err error) {
	_, err = s.db.Exec("DELETE FROM task_tags WHERE taskID = $1 AND tagID = $2", input.TaskID, input.TagID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "RemoveTag"}
	}
	return
}

func (s *storage) GetBoardTags(input models.BoardInput) (tags []models.TagOutside, err error) {
	rows, err := s.db.Query("SELECT tagID, name, color FROM tags WHERE boardID = $1", input.BoardID)
	if err != nil {
		return tags, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoardTags"}
	}
	defer rows.Close()

	for rows.Next() {
		tag := models.TagOutside{}

		err = rows.Scan(&tag.TagID, &tag.Name, &tag.Color)
		if err != nil {
			return tags, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetBoardTags"}
		}

		tags = append(tags, tag)
	}
	return
}
func (s *storage) GetTaskTags(input models.TaskInput) (tags []models.TagOutside, err error) {
	rows, err := s.db.Query("SELECT DISTINCT T.tagID, T.name, T.color FROM task_tags LEFT JOIN tags T ON task_tags.tagID = T.tagID WHERE task_tags.taskID  = $1;", input.TaskID)
	if err != nil {
		return tags, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTaskTags"}
	}
	defer rows.Close()

	for rows.Next() {
		tag := models.TagOutside{}

		err = rows.Scan(&tag.TagID, &tag.Name, &tag.Color)
		if err != nil {
			return tags, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetTaskTags"}
		}

		tags = append(tags, tag)
	}

	return
}