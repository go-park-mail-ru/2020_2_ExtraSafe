package service

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage"
	fStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/fileStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
)

type service struct {
	boardStorage boardStorage.Storage
	fileStorage fStorage.Storage
	profileService protoProfile.ProfileClient
}

var NameService = "BoardService"

func NewService(boardStorage boardStorage.Storage, fileStorage fStorage.Storage, profileService protoProfile.ProfileClient) *service {
	return &service{
		boardStorage: boardStorage,
		fileStorage: fileStorage,
		profileService: profileService,
	}
}

func (s *service) GetBoardsList(_ context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error) {
	userInput := models.UserInput{ID: input.ID}

	boardsList, err := s.boardStorage.GetBoardsList(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = new(protoProfile.BoardsOutsideShort)

	for _, board := range boardsList {
		outputBoard := protoProfile.BoardOutsideShort{
			BoardID: board.BoardID,
			Name:    board.Name,
			Theme:   board.Theme,
			Star:    board.Star,
		}
		output.Boards = append(output.Boards, &outputBoard)
	}
	return output, nil
}

func (s *service) CreateBoard(_ context.Context, input *protoBoard.BoardChangeInput) (output *protoProfile.BoardOutsideShort, err error) {
	userInput := models.BoardChangeInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		BoardName: input.BoardName,
		Theme:     input.Theme,
		Star:      input.Star,
	}

	boardInternal, err := s.boardStorage.CreateBoard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoProfile.BoardOutsideShort{
		BoardID: boardInternal.BoardID,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
	}
	return output, nil
}

func (s *service) GetBoard(c context.Context, input *protoBoard.BoardInput) (output *protoBoard.BoardOutside, err error) {
	userInput := models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}

	boardInternal, err := s.boardStorage.GetBoard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)
	membersIDs = append(membersIDs, boardInternal.UsersIDs...)

	members, err := s.profileService.GetUsersByIDs(c, &protoProfile.UserIDS{UserIDs: membersIDs})
	if err != nil {
		return output, err
	}

	cards := make([]*protoBoard.CardOutside, 0)
	for _, card := range boardInternal.Cards {
		tasks := make([]*protoBoard.TaskOutsideShort, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, &protoBoard.TaskOutsideShort{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
				Tags: convertTags(task.Tags),
				Users: getUsersForTask(task.Users, members),
				Checklists: convertChecklists(task.Checklists),
			})
		}
		cards = append(cards, &protoBoard.CardOutside{
			CardID: card.CardID,
			Name:   card.Name,
			Order:  card.Order,
			Tasks:  tasks,
		})
	}

	output = &protoBoard.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   members.Users[0],
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   members.Users[1:],
		Cards:   cards,
		Tags: 	 convertTags(boardInternal.Tags),
	}

	return output, nil
}

func convertChecklists(checklists []models.ChecklistOutside) (output []*protoBoard.ChecklistOutside) {
	for _, checklist := range checklists {
		output = append(output, &protoBoard.ChecklistOutside{
			ChecklistID: checklist.ChecklistID,
			Items: checklist.Items,
			Name:  checklist.Name,
		})
	}
	return output
}

func convertTags(tags []models.TagOutside) (output []*protoBoard.TagOutside) {
	for _, tag := range tags {
		output = append(output, &protoBoard.TagOutside{
			TagID: tag.TagID,
			Color: tag.Color,
			Name:  tag.Name,
		})
	}
	return output
}

func getUsersForTask(userIDs []int64, users *protoProfile.UsersOutsideShort) []*protoProfile.UserOutsideShort {
	output := make([]*protoProfile.UserOutsideShort, 0)
	for _, id := range userIDs {
		for _, user := range users.Users{
			if id == user.ID {
				output = append(output, user)
			}
		}
	}

	return output
}

func (s *service) ChangeBoard(_ context.Context, input *protoBoard.BoardChangeInput) (output *protoBoard.BoardOutside, err error) {
	userInput := models.BoardChangeInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		BoardName: input.BoardName,
		Theme:     input.Theme,
		Star:      input.Star,
	}

	boardInternal, err := s.boardStorage.ChangeBoard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   nil,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   nil,
		Cards:   nil,
	}
	return output, nil
}

func (s *service) DeleteBoard(_ context.Context, input *protoBoard.BoardInput) (*protoBoard.Nothing, error) {
	userInput := models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}

	err := s.boardStorage.DeleteBoard(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) AddUserToBoard(c context.Context, input *protoBoard.BoardMemberInput) (*protoBoard.Nothing, error) {
	user, err := s.profileService.GetUserByUsername(c, &protoProfile.UserName{UserName: input.MemberName})
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	userInput := models.BoardMember{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		MemberID:  user.ID,
	}

	err = s.boardStorage.AddUser(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) RemoveUserToBoard(c context.Context, input *protoBoard.BoardMemberInput) (*protoBoard.Nothing, error) {
	user, err := s.profileService.GetUserByUsername(c, &protoProfile.UserName{UserName: input.MemberName})
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	userInput := models.BoardMember{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		MemberID:  user.ID,
	}

	err = s.boardStorage.RemoveUser(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CreateCard(_ context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.CreateCard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) GetCard(_ context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.GetCard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) ChangeCard(_ context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.ChangeCard(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) DeleteCard(_ context.Context, input *protoBoard.CardInput) (*protoBoard.Nothing, error) {
	userInput := models.CardInput{
		UserID:    input.UserID,
		CardID:   input.CardID,
		BoardID: input.BoardID,
	}

	err := s.boardStorage.DeleteCard(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CardOrderChange(_ context.Context, input *protoBoard.CardsOrderInput) (*protoBoard.Nothing, error) {
	cardOrder := make([]models.CardOrder, 0)
	for _, card := range input.Cards {
		cardOrder = append(cardOrder, models.CardOrder{
			CardID: card.CardID,
			Order:  card.Order,
		})
	}

	userInput := models.CardsOrderInput{
		UserID:  input.UserID,
		Cards: cardOrder,
	}

	err := s.boardStorage.ChangeCardOrder(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CreateTask(_ context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, err := s.boardStorage.CreateTask(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
	}

	return output, nil
}

func (s *service) GetTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, userIDs, err := s.boardStorage.GetTask(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	taskAssigners, err := s.profileService.GetUsersByIDs(c, &protoProfile.UserIDS{UserIDs: task.Users})
	if err != nil {
		return output, err
	}

	commentUsers, err := s.profileService.GetUsersByIDs(c, &protoProfile.UserIDS{UserIDs: userIDs})
	if err != nil {
		return output, err
	}

	comments := make([]*protoBoard.CommentOutside, 0)
	for i, comment := range task.Comments{
		comments = append(comments, &protoBoard.CommentOutside{
			CommentID: comment.CommentID,
			Message:   comment.Message,
			Order:     comment.Order,
			User:      commentUsers.Users[i],
		})
	}

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
		Tags: convertTags(task.Tags),
		Users: getUsersForTask(task.Users, taskAssigners),
		Checklists: convertChecklists(task.Checklists),
		Comments: comments,
	}

	return output, nil
}

func (s *service) ChangeTask(_ context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, _, err := s.boardStorage.ChangeTask(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	/*users, err := s.profileService.GetUsersByIDs(c, &protoProfile.UserIDS{UserIDs: task.Users})
	if err != nil {
		return output, err
	}*/

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
	}

	return output, nil
}

func (s *service) DeleteTask(_ context.Context, input *protoBoard.TaskInput) (*protoBoard.Nothing, error) {
	userInput := models.TaskInput{
		UserID:    input.UserID,
		TaskID:   input.TaskID,
		CardID: input.CardID,
	}

	err := s.boardStorage.DeleteTask(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) TasksOrderChange(_ context.Context, input *protoBoard.TasksOrderInput) (*protoBoard.Nothing, error) {
	tasksOrder := make([]models.TasksOrder, 0)
	for _, tasks := range input.Tasks {
		taskOrder := make([]models.TaskOrder, 0)
		for _, task := range tasks.Tasks {
			taskOrder = append(taskOrder, models.TaskOrder{
				TaskID: task.TaskID,
				Order:  task.Order,
			})
		}
		tasksOrder = append(tasksOrder, models.TasksOrder{
			CardID: tasks.CardID,
			Tasks:  taskOrder,
		})
	}

	userInput := models.TasksOrderInput{
		UserID:  input.UserID,
		Tasks: tasksOrder,
	}

	err := s.boardStorage.ChangeTaskOrder(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) AssignUser(c context.Context, input *protoBoard.TaskAssignerInput) (*protoBoard.Nothing, error) {
	user, err := s.profileService.GetUserByUsername(c, &protoProfile.UserName{UserName: input.AssignerName})
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	userInput := models.TaskAssigner{
		UserID:    input.UserID,
		TaskID:   input.TaskID,
		AssignerID: user.ID,
	}

	err = s.boardStorage.AssignUser(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) DismissUser(c context.Context, input *protoBoard.TaskAssignerInput) (*protoBoard.Nothing, error) {
	user, err := s.profileService.GetUserByUsername(c, &protoProfile.UserName{UserName: input.AssignerName})
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	userInput := models.TaskAssigner{
		UserID:    input.UserID,
		TaskID:   input.TaskID,
		AssignerID: user.ID,
	}

	err = s.boardStorage.DismissUser(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) GetAssigners(_ context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskAssignerIDs, err error) {
	userInput := models.TaskInput{
		UserID:    input.UserID,
		TaskID:   input.TaskID,
		CardID: input.CardID,
	}

	assigners, err := s.boardStorage.GetAssigners(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.TaskAssignerIDs{AssignerIDs: assigners}

	return output, nil
}

func (s *service) CreateTag(_ context.Context, input *protoBoard.TagInput) (output *protoBoard.TagOutside, err error) {
	userInput := models.TagInput{
		UserID:  input.UserID,
		BoardID: input.BoardID,
		Color:   input.Color,
		Name:    input.Name,
	}

	tag, err := s.boardStorage.CreateTag(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.TagOutside{
		TagID: tag.TagID,
		Color: tag.Color,
		Name:  tag.Name,
	}

	return output, nil
}

func (s *service) ChangeTag(_ context.Context, input *protoBoard.TagInput) (output *protoBoard.TagOutside, err error) {
	userInput := models.TagInput{
		UserID:  input.UserID,
		TagID: input.TagID,
		BoardID: input.BoardID,
		Color:   input.Color,
		Name:    input.Name,
	}

	tag, err := s.boardStorage.UpdateTag(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.TagOutside{
		TagID: tag.TagID,
		Color: tag.Color,
		Name:  tag.Name,
	}

	return output, nil
}

func (s *service) DeleteTag(_ context.Context, input *protoBoard.TagInput) (*protoBoard.Nothing, error) {
	userInput := models.TagInput{
		UserID:  input.UserID,
		TagID: input.TagID,
		BoardID: input.BoardID,
	}

	err := s.boardStorage.DeleteTag(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) AddTag(_ context.Context, input *protoBoard.TaskTagInput) (*protoBoard.Nothing, error) {
	userInput := models.TaskTagInput{
		UserID:  input.UserID,
		TagID: input.TagID,
		TaskID: input.TaskID,
	}

	err := s.boardStorage.AddTag(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) RemoveTag(_ context.Context, input *protoBoard.TaskTagInput) (*protoBoard.Nothing, error) {
	userInput := models.TaskTagInput{
		UserID:  input.UserID,
		TagID: input.TagID,
		TaskID: input.TaskID,
	}

	err := s.boardStorage.RemoveTag(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CreateComment(ctx context.Context, input *protoBoard.CommentInput) (output *protoBoard.CommentOutside, err error) {
	userInput := models.CommentInput{
		CommentID: input.CommentID,
		UserID:  input.UserID,
		TaskID: input.TaskID,
		Message: input.Message,
		Order: input.Order,
	}

	comment, err := s.boardStorage.CreateComment(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	//TODO - менить на GetUserByID
	user, err := s.profileService.GetUsersByIDs(ctx, &protoProfile.UserIDS{UserIDs: []int64{input.UserID}})
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.CommentOutside{
		CommentID: comment.CommentID,
		Message:   comment.Message,
		Order:     comment.Order,
		User:      user.Users[0],
	}

	return output, nil
}

func (s *service) ChangeComment(ctx context.Context, input *protoBoard.CommentInput) (output *protoBoard.CommentOutside, err error) {
	userInput := models.CommentInput{
		CommentID: input.CommentID,
		UserID:  input.UserID,
		TaskID: input.TaskID,
		Message: input.Message,
		Order: input.Order,
	}

	comment, err := s.boardStorage.UpdateComment(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	//TODO - менить на GetUserByID
	user, err := s.profileService.GetUsersByIDs(ctx, &protoProfile.UserIDS{UserIDs: []int64{input.UserID}})
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.CommentOutside{
		CommentID: comment.CommentID,
		Message:   comment.Message,
		Order:     comment.Order,
		User:      user.Users[0],
	}

	return output, nil
}

func (s *service) DeleteComment(_ context.Context, input *protoBoard.CommentInput) (*protoBoard.Nothing, error) {
	userInput := models.CommentInput{
		CommentID: input.CommentID,
		UserID:  input.UserID,
		TaskID: input.TaskID,
		Message: input.Message,
		Order: input.Order,
	}

	err := s.boardStorage.DeleteComment(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CreateChecklist(_ context.Context, input *protoBoard.ChecklistInput) (output *protoBoard.ChecklistOutside, err error) {
	userInput := models.ChecklistInput{
		UserID: input.UserID,
		ChecklistID: input.ChecklistID,
		TaskID: input.TaskID,
		Name: input.Name,
		Items: input.Items,
	}

	checklist, err := s.boardStorage.CreateChecklist(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.ChecklistOutside{
		ChecklistID: checklist.ChecklistID,
		Name:   checklist.Name,
		Items: checklist.Items,
	}

	return output, nil
}

func (s *service) ChangeChecklist(_ context.Context, input *protoBoard.ChecklistInput) (output *protoBoard.ChecklistOutside, err error) {
	userInput := models.ChecklistInput{
		UserID: input.UserID,
		ChecklistID: input.ChecklistID,
		TaskID: input.TaskID,
		Name: input.Name,
		Items: input.Items,
	}

	checklist, err := s.boardStorage.UpdateChecklist(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	output = &protoBoard.ChecklistOutside{
		ChecklistID: checklist.ChecklistID,
		Name:   checklist.Name,
		Items: checklist.Items,
	}

	return output, nil
}

func (s *service) DeleteChecklist(_ context.Context, input *protoBoard.ChecklistInput) (*protoBoard.Nothing, error) {
	userInput := models.ChecklistInput{
		UserID: input.UserID,
		ChecklistID: input.ChecklistID,
		TaskID: input.TaskID,
		Name: input.Name,
		Items: input.Items,
	}

	err := s.boardStorage.DeleteChecklist(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) AddAttachment(_ context.Context, input *protoBoard.AttachmentInput) (output *protoBoard.AttachmentOutside, err error) {
	fileInput := models.AttachmentFileInternal{
		UserID:   input.UserID,
		Filename: input.Filename,
		File:     input.File,
	}

	userInput := &models.AttachmentInternal{
		TaskID:       input.TaskID,
		Filename:     input.Filename,
	}

	err = s.fileStorage.UploadFile(fileInput, userInput, false)
	if err != nil {
		return output, err
	}

	attachment, err := s.boardStorage.AddAttachment(*userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.AttachmentOutside{
		AttachmentID: attachment.AttachmentID,
		Filename:     attachment.Filename,
		Filepath:     attachment.Filepath,
	}

	return output, nil
}

func (s *service) RemoveAttachment(_ context.Context, input *protoBoard.AttachmentInput) (*protoBoard.Nothing, error) {
	userInput := models.AttachmentInternal{
		TaskID:       input.TaskID,
		Filename:     input.Filename,
		AttachmentID: input.AttachmentID,
	}

	err := s.boardStorage.RemoveAttachment(userInput)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, err
	}

	err = s.fileStorage.DeleteFile(userInput.Filepath, false)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, err
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CheckBoardPermission(_ context.Context, input *protoBoard.CheckPermissions) (*protoBoard.Nothing, error) {
	err := s.boardStorage.CheckBoardPermission(input.UserID, input.ElementID, input.IfAdmin)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CheckCardPermission(_ context.Context, input *protoBoard.CheckPermissions) (*protoBoard.Nothing, error) {
	err := s.boardStorage.CheckCardPermission(input.UserID, input.ElementID)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CheckTaskPermission(_ context.Context, input *protoBoard.CheckPermissions) (*protoBoard.Nothing, error) {
	err := s.boardStorage.CheckTaskPermission(input.UserID, input.ElementID)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}

func (s *service) CheckCommentPermission(_ context.Context, input *protoBoard.CheckPermissions) (*protoBoard.Nothing, error) {
	err := s.boardStorage.CheckCommentPermission(input.UserID, input.ElementID, input.IfAdmin)
	if err != nil {
		return &protoBoard.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), NameService)
	}

	return &protoBoard.Nothing{Dummy: true}, nil
}