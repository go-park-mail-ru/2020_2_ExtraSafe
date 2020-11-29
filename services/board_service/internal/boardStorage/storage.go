package boardStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/attachmentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
)

type Storage interface {
	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error

	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardInternal, error)
	ChangeCardOrder(cardInput models.CardsOrderInput) error
	DeleteCard(userInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
	DeleteTask(taskInput models.TaskInput) error

	//TODO tests
	AddUser(input models.BoardMember) (err error)
	//TODO tests
	RemoveUser(input models.BoardMember) (err error)

	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	GetCard(cardInput models.CardInput) (models.CardInternal, error)
	GetTask(taskInput models.TaskInput) (models.TaskInternal, []int64, error)
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (err error)
	CheckTaskPermission(userID int64, taskID int64) (err error)
	CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (err error)

	AssignUser(input models.TaskAssigner) (err error)
	DismissUser(input models.TaskAssigner) (err error)

	CreateTag(input models.TagInput) (tag models.TagOutside, err error)
	UpdateTag(input models.TagInput) (tag models.TagOutside, err error)
	DeleteTag(input models.TagInput) (err error)
	AddTag(input models.TaskTagInput) (err error)
	RemoveTag(input models.TaskTagInput) (err error)

	CreateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	DeleteComment(input models.CommentInput) (err error)

	CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	DeleteChecklist(input models.ChecklistInput) (err error)

	AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInternal) (err error)
}

type storage struct {
	db           *sql.DB
	cardsStorage CardsStorage
	tasksStorage TasksStorage
	tagStorage tagStorage.Storage
	commentStorage commentStorage.Storage
	checklistStorage checklistStorage.Storage
	attachmentStorage attachmentStorage.Storage
}

func NewStorage(db *sql.DB, cardsStorage CardsStorage, tasksStorage TasksStorage, tagStorage tagStorage.Storage, commentStorage commentStorage.Storage, checklistStorage checklistStorage.Storage, attachmentStorage attachmentStorage.Storage) Storage {
	return &storage{
		db: db,
		cardsStorage: cardsStorage,
		tasksStorage: tasksStorage,
		tagStorage: tagStorage,
		commentStorage: commentStorage,
		checklistStorage: checklistStorage,
		attachmentStorage: attachmentStorage,
	}
}

func (s *storage) GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error) {
	boards := make([]models.BoardOutsideShort, 0)

	rows, err := s.db.Query("SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B " +
									"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID WHERE B.adminID = $1 OR M.userID = $1;", userInput.ID)
	if err != nil {
		return []models.BoardOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoardsList"}
	}
	defer rows.Close()

	for rows.Next() {
		var board models.BoardOutsideShort

		err = rows.Scan(&board.BoardID, &board.Name, &board.Theme, &board.Star)
		if err != nil {
			return []models.BoardOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetBoardsList"}
		}

		boards = append(boards, board)
	}
	return boards, nil
}

func (s *storage) CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error) {
	var boardID int64

	err := s.db.QueryRow("INSERT INTO boards (adminID, boardName, theme, star) VALUES ($1, $2, $3, $4) RETURNING boardID",
		boardInput.UserID,
		boardInput.BoardName,
		boardInput.Theme,
		boardInput.Star).Scan(&boardID)

	if err != nil {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateBoard"}
	}

	board := models.BoardInternal{
		BoardID:  boardID,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardInternal{},
		UsersIDs: []int64{},
	}
	return board, nil
}

func (s *storage) GetBoard(boardInput models.BoardInput) (models.BoardInternal, error) {
	board := models.BoardInternal{}
	board.BoardID = boardInput.BoardID

	err := s.db.QueryRow("SELECT adminID, boardName, theme, star FROM boards WHERE boardID = $1", boardInput.BoardID).
				Scan(&board.AdminID, &board.Name, &board.Theme, &board.Star)

	if err != nil {
		fmt.Println(err)
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoard"}
	}

	members, err := s.getBoardMembers(boardInput)
	if err != nil {
		return models.BoardInternal{}, err
	}

	board.UsersIDs = append(board.UsersIDs, members...)

	cards, err := s.cardsStorage.GetCardsByBoard(boardInput)
	if err != nil {
		return models.BoardInternal{}, err
	}
	for _, card := range cards {
		cardInput := models.CardInput{CardID: card.CardID}

		tasks, err := s.tasksStorage.GetTasksByCard(cardInput)
		if err != nil {
			return models.BoardInternal{}, err
		}

		for i, task := range tasks {
			taskInput := models.TaskInput{TaskID: task.TaskID}

			tags, err := s.tagStorage.GetTaskTags(taskInput)
			if err != nil {
				return models.BoardInternal{}, err
			}
			tasks[i].Tags = append(tasks[i].Tags, tags...)

			users, err := s.tasksStorage.GetAssigners(taskInput)
			if err != nil {
				return models.BoardInternal{}, err
			}
			tasks[i].Users = append(tasks[i].Users, users...)

			checklists, err := s.checklistStorage.GetChecklistsByTask(taskInput)
			if err != nil {
				return models.BoardInternal{}, err
			}
			tasks[i].Checklists = append(tasks[i].Checklists, checklists...)
		}

		card.Tasks = append(card.Tasks, tasks...)
		board.Cards = append(board.Cards, card)
	}

	tags, err := s.tagStorage.GetBoardTags(boardInput)
	if err != nil {
		return models.BoardInternal{}, err
	}
	board.Tags = append(board.Tags, tags...)

	return board, nil
}

func (s *storage) getBoardMembers(boardInput models.BoardInput) ([]int64, error) {
	members := make([]int64, 0)

	rows, err := s.db.Query("SELECT userID from board_members WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		return []int64{}, models.ServeError{Codes: []string{"500"}, MethodName: "getBoardMembers"}
	}
	defer rows.Close()

	for rows.Next() {
		var userID int64

		err = rows.Scan(&userID)
		if err != nil {
			return []int64{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "getBoardMembers"}
		}

		members = append(members, userID)
	}

	return members, nil
}

func (s *storage) ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error) {
	board := models.BoardInternal{}

	err := s.db.QueryRow("UPDATE boards SET boardName = $1, theme = $2, star = $3 WHERE boardID = $4 RETURNING adminID",
								boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).Scan(&board.AdminID)
	if err != nil {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeBoard"}
	}

	board.BoardID = boardInput.BoardID
	board.Name = boardInput.BoardName
	board.Theme = boardInput.Theme
	board.Star = boardInput.Star

	return board, nil
}

func (s *storage) DeleteBoard(boardInput models.BoardInput) error {
	_, err := s.db.Exec("DELETE FROM boards WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteBoard"}
	}

	return nil
}

func (s *storage) CreateCard(cardInput models.CardInput) (models.CardOutside, error) {
	//не ищем таски, потому что при создании доски они пустые
	return s.cardsStorage.CreateCard(cardInput)
}

func (s *storage) ChangeCard(cardInput models.CardInput) (models.CardInternal, error) {
	return s.cardsStorage.ChangeCard(cardInput)
}

func (s *storage) ChangeCardOrder(cardInput models.CardsOrderInput) error {
	return s.cardsStorage.ChangeCardOrder(cardInput)
}

func (s *storage) DeleteCard(cardInput models.CardInput) error {
	return s.cardsStorage.DeleteCard(cardInput)
}

func (s *storage) CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error) {
	return s.tasksStorage.CreateTask(taskInput)
}

func (s *storage) ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error) {
	return s.tasksStorage.ChangeTask(taskInput)
}

func (s *storage) ChangeTaskOrder(taskInput models.TasksOrderInput) error {
	return s.tasksStorage.ChangeTaskOrder(taskInput)
}

func (s *storage) DeleteTask(taskInput models.TaskInput) error {
	return s.tasksStorage.DeleteTask(taskInput)
}

func (s *storage) GetCard(cardInput models.CardInput) (models.CardInternal, error) {
	card, err := s.cardsStorage.GetCardByID(cardInput)
	if err != nil {
		return models.CardInternal{}, err
	}

	tasks, err := s.tasksStorage.GetTasksByCard(cardInput)
	if err != nil {
		return models.CardInternal{}, err
	}

	//FIXME добавить сбор фич для тасок
	card.Tasks = append(card.Tasks, tasks...)
	return card, nil
}

func (s *storage) GetTask(taskInput models.TaskInput) (models.TaskInternal, []int64, error) {
	task, err := s.tasksStorage.GetTaskByID(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}

	tags, err := s.tagStorage.GetTaskTags(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}
	task.Tags = append(task.Tags, tags...)

	users, err := s.tasksStorage.GetAssigners(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}
	task.Users = append(task.Users, users...)

	checklists, err := s.checklistStorage.GetChecklistsByTask(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}
	task.Checklists = append(task.Checklists, checklists...)

	comments, userIDs, err := s.commentStorage.GetCommentsByTask(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}
	task.Comments = append(task.Comments, comments...)

	attachments, err := s.attachmentStorage.GetAttachmentsByTask(taskInput)
	if err != nil {
		return models.TaskInternal{}, nil, err
	}
	task.Attachments = append(task.Attachments, attachments...)
	return task, userIDs, nil
}

func (s *storage) AddUser(input models.BoardMember) (err error) {
	_, err = s.db.Exec("INSERT INTO board_members (boardID, userID) VALUES ($1, $2)", input.BoardID, input.MemberID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "AddUser"}
	}
	return
}

func (s *storage) RemoveUser(input models.BoardMember) (err error) {
	_, err = s.db.Exec("DELETE FROM board_members WHERE boardID = $1 AND userID = $2", input.BoardID, input.MemberID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "RemoveUser"}
	}
	return
}

func (s *storage) CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error) {
	var ok bool

	if ifAdmin {
		ok, err = s.checkBoardAdminPermission(userID, boardID)
	} else {
		ok, err = s.checkBoardUserPermission(userID, boardID)
	}

	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckBoardPermission"}
	}

	if !ok {
		return models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckBoardPermission"}
	}

	return nil
}

func (s *storage) checkBoardAdminPermission(userID int64, boardID int64) (flag bool, err error) {
	var ID int64
	err = s.db.QueryRow("SELECT boardID FROM boards WHERE boardID = $1 AND adminID = $2", boardID, userID).Scan(&ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func (s *storage) checkBoardUserPermission(userID int64, boardID int64) (flag bool, err error) {
	var ID int64
	err = s.db.QueryRow("SELECT boardID FROM boards B LEFT OUTER JOIN board_members M ON M.boardID = B.boardID WHERE B.boardID = $1 AND (M.userID = $2 OR B.adminID = $2)", boardID, userID).Scan(&ID)
	//err = s.db.QueryRow("SELECT boardID FROM board_members WHERE boardID = $1 AND userID = $2", boardID, userID).Scan(&ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func (s *storage) CheckCardPermission(userID int64, cardID int64) (err error) {
	var boardID int64

	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR M.userID = $1) AND cardID = $2", userID, cardID).Scan(&boardID)
	if err != nil && err != sql.ErrNoRows {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCardPermission"}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckCardPermission"}
	}

	return nil
}

func (s *storage) CheckTaskPermission(userID int64, taskID int64) (err error) {
	var boardID int64

	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"JOIN tasks T on T.cardID = C.cardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR M.userID = $1) AND taskID = $2", userID, taskID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckTaskPermission"}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckTaskPermission"}
	}

	return nil
}

func (s *storage) CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (err error) {
	var boardID int64
	var query string

	if ifAdmin {
		query = "SELECT B.boardID FROM boards B " +
			"JOIN cards C on C.boardID = B.boardID " +
			"JOIN tasks T on T.cardID = C.cardID " +
			"JOIN comments Com on Com.taskID = T.taskID" +
			"WHERE (B.adminID = 1 OR Com.userID = 1) AND Com.commentID = $2"
	} else {
		query = "SELECT B.boardID FROM boards B " +
			"JOIN cards C on C.boardID = B.boardID " +
			"JOIN tasks T on T.cardID = C.cardID " +
			"JOIN comments Com on Com.taskID = T.taskID" +
			"WHERE Com.userID = $1 = $1 AND Com.commentID = $2"
	}

	err = s.db.QueryRow(query, userID, commentID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCommentPermission"}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckCommentPermission"}
	}

	return nil
}