package boardStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

//FIXME убрать лишние функции

func (s *storage) AssignUser(input models.TaskAssigner) (err error) {
	return s.tasksStorage.AssignUser(input)
}

func (s *storage) DismissUser(input models.TaskAssigner) (err error) {
	return s.tasksStorage.DismissUser(input)
}
func (s *storage) GetAssigners(input models.TaskInput) (assignerIDs []int64, err error){
	return s.tasksStorage.GetAssigners(input)
}

func (s *storage) CreateTag(input models.TagInput) (tag models.TagOutside, err error){
	return s.tagStorage.CreateTag(input)
}

func (s *storage) UpdateTag(input models.TagInput) (tag models.TagOutside, err error){
	return s.tagStorage.UpdateTag(input)
}

func (s *storage) DeleteTag(input models.TagInput) (err error){
	return s.tagStorage.DeleteTag(input)
}

func (s *storage) AddTag(input models.TaskTagInput) (err error){
	return s.tagStorage.AddTag(input)
}

func (s *storage) RemoveTag(input models.TaskTagInput) (err error){
	return s.tagStorage.RemoveTag(input)
}

func (s *storage) GetBoardTags(input models.BoardInput) (tags []models.TagOutside, err error){
	return s.tagStorage.GetBoardTags(input)
}

func (s *storage) GetTaskTags(input models.TaskInput) (tags []models.TagOutside, err error){
	return s.tagStorage.GetTaskTags(input)
}

func (s *storage) CreateComment(input models.CommentInput) (comment models.CommentOutside, err error){
	return s.commentStorage.CreateComment(input)
}

func (s *storage) UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error){
	return s.commentStorage.UpdateComment(input)
}

func (s *storage) DeleteComment(input models.CommentInput) (err error){
	return s.commentStorage.DeleteComment(input)
}

func (s *storage) GetCommentsByTask(input models.TaskInput) (comments []models.CommentOutside, userIDS[] int64, err error){
	return s.commentStorage.GetCommentsByTask(input)
}

func (s *storage) CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error){
	return s.checklistStorage.CreateChecklist(input)
}

func (s *storage) UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error){
	return s.checklistStorage.UpdateChecklist(input)
}

func (s *storage) DeleteChecklist(input models.ChecklistInput) (err error){
	return s.checklistStorage.DeleteChecklist(input)
}

func (s *storage) GetChecklistsByTask(input models.TaskInput) (checklists []models.ChecklistOutside, err error){
	return s.checklistStorage.GetChecklistsByTask(input)
}
