package checkliststorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	DeleteChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)

	GetChecklistsByTask(input models.TaskInput) (checklists []models.ChecklistOutside, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error) {
	err = s.db.QueryRow("INSERT INTO checklists (taskID, name, items) VALUES ($1, $2, $3) RETURNING checklistID",
		input.TaskID, input.Name, input.Items).Scan(&checklist.ChecklistID)
	if err != nil {
		return checklist, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateChecklist"}
	}
	checklist.Items = input.Items
	checklist.Name = input.Name
	checklist.TaskID = input.TaskID
	return
}

func (s *storage) UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error) {
	err = s.db.QueryRow("UPDATE checklists SET name = $1, items = $2 WHERE checklistID = $3 RETURNING taskID", input.Name, input.Items, input.ChecklistID).
		Scan(&checklist.TaskID)
	if err != nil {
		return checklist, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "UpdateChecklist"}
	}
	checklist.ChecklistID = input.ChecklistID
	checklist.Items = input.Items
	checklist.Name = input.Name
	return
}

func (s *storage) DeleteChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error) {
	err = s.db.QueryRow("DELETE FROM checklists WHERE checklistID = $1 RETURNING taskID", input.ChecklistID).
		Scan(&checklist.TaskID)
	if err != nil {
		return checklist, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "DeleteChecklist"}
	}
	checklist.ChecklistID = input.ChecklistID
	return
}

func (s *storage) GetChecklistsByTask(input models.TaskInput) (checklists []models.ChecklistOutside, err error) {
	rows, err := s.db.Query("SELECT checklistID, name, items FROM checklists WHERE taskID = $1", input.TaskID)
	if err != nil {
		return checklists, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetChecklistsByTask"}
	}
	defer rows.Close()

	for rows.Next() {
		checklist := models.ChecklistOutside{}

		err = rows.Scan(&checklist.ChecklistID, &checklist.Name, &checklist.Items)
		if err != nil {
			return checklists, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetChecklistsByTask"}
		}

		checklists = append(checklists, checklist)
	}
	return
}
