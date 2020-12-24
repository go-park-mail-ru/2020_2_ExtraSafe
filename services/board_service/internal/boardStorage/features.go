package boardStorage

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

func (s *storage) AssignUser(input models.TaskAssigner) (task models.TaskAssignUserOutside, err error) {
	err = s.tasksStorage.AssignUser(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(input.TaskID)
	if err != nil {
		return
	}
	task.TaskName, err = s.tasksStorage.GetTaskName(models.TaskInput{TaskID: input.TaskID})
	if err != nil {
		return
	}

	task.CardID = cardID
	task.TaskID = input.TaskID

	return
}

func (s *storage) DismissUser(input models.TaskAssigner) (task models.TaskAssignUserOutside, err error) {
	err = s.tasksStorage.DismissUser(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(input.TaskID)
	if err != nil {
		return
	}

	task.CardID = cardID
	task.TaskID = input.TaskID

	return
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

func (s *storage) AddTag(input models.TaskTagInput) (tag models.TagOutside, err error){
	tag, err = s.tagStorage.AddTag(input)
	if err != nil {
		return
	}
	tag, err = s.tagStorage.GetTag(tag.TagID)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(input.TaskID)
	if err != nil {
		return
	}
	tag.TaskID  = input.TaskID
	tag.CardID = cardID
	return
}

func (s *storage) RemoveTag(input models.TaskTagInput) (tag models.TagOutside, err error){
	tag, err = s.tagStorage.RemoveTag(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(tag.TaskID)
	if err != nil {
		return
	}
	tag.CardID = cardID
	return
}

func (s *storage) CreateComment(input models.CommentInput) (comment models.CommentOutside, err error){
	comment, err = s.commentStorage.CreateComment(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(comment.TaskID)
	if err != nil {
		return
	}
	comment.CardID = cardID
	return
}

func (s *storage) UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error){
	comment, err = s.commentStorage.UpdateComment(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(comment.TaskID)
	if err != nil {
		return
	}
	comment.CardID = cardID
	return
}

func (s *storage) DeleteComment(input models.CommentInput) (comment models.CommentOutside, err error){
	comment, err = s.commentStorage.DeleteComment(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(comment.TaskID)
	if err != nil {
		return
	}
	comment.CardID = cardID
	return
}

func (s *storage) CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error){
	checklist, err = s.checklistStorage.CreateChecklist(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(checklist.TaskID)
	if err != nil {
		return
	}
	checklist.CardID = cardID
	return
}

func (s *storage) UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error){
	checklist, err = s.checklistStorage.UpdateChecklist(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(checklist.TaskID)
	if err != nil {
		return
	}
	checklist.CardID = cardID
	return
}

func (s *storage) DeleteChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error){
	checklist, err = s.checklistStorage.DeleteChecklist(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(checklist.TaskID)
	if err != nil {
		return
	}
	checklist.CardID = cardID
	return
}

func (s *storage) AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error) {
	attachment, err = s.attachmentStorage.AddAttachment(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(attachment.TaskID)
	if err != nil {
		return
	}
	attachment.CardID = cardID
	return
}

func (s *storage) RemoveAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error) {
	attachment, err = s.attachmentStorage.RemoveAttachment(input)
	if err != nil {
		return
	}
	cardID, err := s.tasksStorage.GetCardIDByTask(attachment.TaskID)
	if err != nil {
		return
	}
	attachment.CardID = cardID
	return
}
