package checklistStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Storage interface {
	CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	DeleteChecklist(input models.ChecklistInput) (err error)

	GetChecklistsByTask(input models.TaskInput) (checklists []models.ChecklistOutside, err error)
}