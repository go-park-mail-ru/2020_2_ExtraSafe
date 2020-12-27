package boardStorage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/attachmentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
	"hash/adler32"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//go:generate mockgen -destination=../../../board_service/internal/service/mock/mock_boardStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage BoardStorage

type BoardStorage interface {
	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error
	CreateBoardFromTemplate(boardInput models.BoardInputTemplate) (models.BoardInternal, error)

	//GetTemplate(boardInput models.BoardInput) (, error)
	GetTemplates() ([]models.BoardTemplateOutsideShort, error)

	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardInternal, error)
	ChangeCardOrder(cardInput models.CardsOrderInput) error
	DeleteCard(userInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
	DeleteTask(taskInput models.TaskInput) (models.TaskInternalShort, error)

	AddUser(input models.BoardMember) (err error)
	RemoveUser(input models.BoardMember) (err error)

	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	GetBoardShort(boardInput models.BoardInput) (models.BoardOutsideShort, error)
	GetCard(cardInput models.CardInput) (models.CardInternal, error)
	GetTask(taskInput models.TaskInput) (models.TaskInternal, []int64, error)
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)

	GetSharedURL(boardInput models.BoardInput) (string, error)
	GetBoardByURL(boardInput models.BoardInviteInput) (models.BoardOutsideShort, error)

	CheckIfAdmin(userID int64, boardID int64) (flag bool, err error)

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (boardID int64, err error)
	CheckTaskPermission(userID int64, taskID int64) (boardID int64, err error)
	CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (boardID int64, err error)

	AssignUser(input models.TaskAssigner) (task models.TaskAssignUserOutside, err error)
	DismissUser(input models.TaskAssigner) (task models.TaskAssignUserOutside, err error)

	CreateTag(input models.TagInput) (tag models.TagOutside, err error)
	UpdateTag(input models.TagInput) (tag models.TagOutside, err error)
	DeleteTag(input models.TagInput) (err error)
	AddTag(input models.TaskTagInput) (tag models.TagOutside, err error)
	RemoveTag(input models.TaskTagInput) (tag models.TagOutside, err error)

	CreateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	DeleteComment(input models.CommentInput) (comment models.CommentOutside, err error)

	CreateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	UpdateChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)
	DeleteChecklist(input models.ChecklistInput) (checklist models.ChecklistOutside, err error)

	AddAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
	RemoveAttachment(input models.AttachmentInternal) (attachment models.AttachmentOutside, err error)
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

func (s *storage) GetTemplates() ([]models.BoardTemplateOutsideShort, error) {
	files, err := ioutil.ReadDir("../../../templates")
	if err != nil {
		return []models.BoardTemplateOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTemplates"}
	}

	templates := make([]models.BoardTemplateOutsideShort, 0)

	for _, file := range files {
		templateJsonFile, err := os.Open(file.Name())
		if err != nil {
			return []models.BoardTemplateOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetTemplates"}
		}
		templateJsonFile.Close()

		templateValue, _ := ioutil.ReadAll(templateJsonFile)

		board := models.BoardInternalTemplate{}

		err = json.Unmarshal(templateValue, &board)
		if err != nil {
			fmt.Println("cannot unmarshall", err)
			return []models.BoardTemplateOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetTemplates"}
		}

		template := models.BoardTemplateOutsideShort{
			TemplateSlug: board.Slug,
			TemplateName: board.BoardName,
			Description:  board.Description,
		}
		templates = append(templates, template)
	}

	return templates, nil
}

/*func (s *storage) getTemplate(templateSlug string) {

}*/

func (s *storage) CreateBoardFromTemplate(boardInput models.BoardInputTemplate) (models.BoardInternal, error) {
	fileTemplate := fmt.Sprintf("%s.json", boardInput.TemplateSlug)
	fmt.Println(fileTemplate)

	templateJsonFile, err := os.Open(fileTemplate)
	if err != nil {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateBoardFromTemplate"}
	}
	defer templateJsonFile.Close()

	templateValue, _ := ioutil.ReadAll(templateJsonFile)

	board := models.BoardInternalTemplate{}

	err = json.Unmarshal(templateValue, &board)
	if err != nil {
		fmt.Println("cannot unmarshall", err)
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateBoardFromTemplate"}
	}

	//create board
	boardOutside, err := s.createBoardInternal(board)
	if err != nil {
		return models.BoardInternal{}, err
	}

	//create tags
	for _, currentTag := range board.Tags {
		tag, err := s.tagStorage.CreateTag(models.TagInput{
			BoardID:   boardOutside.BoardID,
			Color:     currentTag.Color,
			Name:      currentTag.Name,
		})
		if err != nil {
			return models.BoardInternal{}, err
		}

		boardOutside.Tags = append(boardOutside.Tags, tag)
	}

	for _, currentCard := range board.Cards {
		//create cards
		card, err := s.cardsStorage.CreateCardInternal(currentCard, boardOutside.BoardID)
		if err != nil {
			return models.BoardInternal{}, err
		}

		for j, currentTask := range card.Tasks {
			inputTask := models.TaskInput{
				CardID:      card.CardID,
				Name:        currentTask.Name,
				Description: currentTask.Description,
				Order:       currentTask.Order,
			}
			//create tasks
			task, err := s.tasksStorage.CreateTask(inputTask)
			if err != nil {
				return models.BoardInternal{}, err
			}

			card.Tasks = append(card.Tasks, task)
		}

		boardOutside.Cards = append(boardOutside.Cards, card)
	}

	return boardOutside, nil
}

func (s *storage) createBoardInternal(boardInput models.BoardInternalTemplate) (models.BoardInternal, error) {
	var boardID int64

	url := createSharedUrl(boardInput.AdminID, boardInput.BoardName)
	urlString := strconv.FormatUint(uint64(url), 10)
	err := s.db.QueryRow("INSERT INTO boards (adminID, boardName, theme, star, shared_url) VALUES ($1, $2, $3, $4, $5) RETURNING boardID",
		boardInput.AdminID,
		boardInput.BoardName,
		"",
		false,
		urlString).
		Scan(&boardID)

	if err != nil {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "createBoardInternal"}
	}

	board := models.BoardInternal{
		BoardID:  boardID,
		AdminID:  boardInput.AdminID,
		Name:     boardInput.BoardName,
		Cards:    []models.CardInternal{},
		UsersIDs: []int64{},
	}
	return board, nil
}

func NewStorage(db *sql.DB, cardsStorage CardsStorage, tasksStorage TasksStorage, tagStorage tagStorage.Storage, commentStorage commentStorage.Storage, checklistStorage checklistStorage.Storage, attachmentStorage attachmentStorage.Storage) BoardStorage {
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

	url := createSharedUrl(boardInput.UserID, boardInput.BoardName)
	urlString := strconv.FormatUint(uint64(url), 10)
	err := s.db.QueryRow("INSERT INTO boards (adminID, boardName, theme, star, shared_url) VALUES ($1, $2, $3, $4, $5) RETURNING boardID",
		boardInput.UserID,
		boardInput.BoardName,
		boardInput.Theme,
		boardInput.Star,
		urlString).
		Scan(&boardID)

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

func (s *storage) GetBoardShort(boardInput models.BoardInput) (models.BoardOutsideShort, error) {
	board := models.BoardOutsideShort{}

	err := s.db.QueryRow("SELECT boardID, boardName, theme, star FROM boards WHERE boardID = $1", boardInput.BoardID).
		Scan(&board.BoardID, &board.Name, &board.Theme, &board.Star)
	if err != nil {
		return board, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoardShort"}
	}

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

func (s *storage) DeleteTask(taskInput models.TaskInput) (models.TaskInternalShort, error) {
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

func (s *storage) CheckIfAdmin(userID int64, boardID int64) (flag bool, err error) {
	return s.checkBoardAdminPermission(userID, boardID)
}

func (s *storage) checkBoardAdminPermission(userID int64, boardID int64) (flag bool, err error) {
	var ID int64
	err = s.db.QueryRow("SELECT boardID FROM boards WHERE boardID = $1 AND adminID = $2", boardID, userID).Scan(&ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func (s *storage) checkBoardUserPermission(userID int64, boardID int64) (flag bool, err error) {
	var ID int64
	err = s.db.QueryRow("SELECT B.boardID FROM boards B LEFT OUTER JOIN board_members M ON M.boardID = B.boardID WHERE B.boardID = $1 AND (M.userID = $2 OR B.adminID = $2)", boardID, userID).Scan(&ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func (s *storage) CheckCardPermission(userID int64, cardID int64) (boardID int64, err error) {
	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR M.userID = $1) AND cardID = $2", userID, cardID).Scan(&boardID)
	if err != nil && err != sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCardPermission"}
	}

	if err == sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckCardPermission"}
	}

	return boardID, nil
}

func (s *storage) CheckTaskPermission(userID int64, taskID int64) (boardID int64, err error) {
	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"JOIN tasks T on T.cardID = C.cardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR M.userID = $1) AND taskID = $2", userID, taskID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckTaskPermission"}
	}

	if err == sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckTaskPermission"}
	}

	return boardID, nil
}

func (s *storage) CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (boardID int64, err error) {
	var query string

	if ifAdmin {
		query = "SELECT B.boardID FROM boards B JOIN cards C on C.boardID = B.boardID JOIN tasks T on T.cardID = C.cardID JOIN comments Com on Com.taskID = T.taskID WHERE (B.adminID = $1 OR Com.userID = $1) AND Com.commentID = $2"
		/*query = "SELECT B.boardID FROM boards B " +
			"JOIN cards C on C.boardID = B.boardID " +
			"JOIN tasks T on T.cardID = C.cardID " +
			"JOIN comments Com on Com.taskID = T.taskID" +
			"WHERE (B.adminID = $1 OR Com.userID = $1) AND Com.commentID = $2"*/
	} else {
		query = "SELECT B.boardID FROM boards B JOIN cards C on C.boardID = B.boardID JOIN tasks T on T.cardID = C.cardID JOIN comments Com on Com.taskID = T.taskID WHERE Com.userID = $1 AND Com.commentID = $2"
	}

	err = s.db.QueryRow(query, userID, commentID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCommentPermission"}
	}

	if err == sql.ErrNoRows {
		return 0, models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckCommentPermission"}
	}

	return boardID, nil
}

func (s *storage) GetSharedURL(boardInput models.BoardInput) (string, error) {
	var url string
	err := s.db.QueryRow("SELECT shared_url FROM boards WHERE boardID = $1", boardInput.BoardID).Scan(&url)
	if err != nil {
		return "", models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "GetSharedURL"}
	}

	return url, nil
}

func createSharedUrl(adminID int64, boardName string) uint32 {
	data := []byte(strconv.FormatInt(adminID, 10) + boardName + time.Now().String())
	return adler32.Checksum(data)
}

func (s *storage) GetBoardByURL(boardInput models.BoardInviteInput) (models.BoardOutsideShort, error) {
	board := models.BoardOutsideShort{}

	err := s.db.QueryRow("SELECT boardID, boardName, theme, star FROM boards WHERE shared_url = $1", boardInput.UrlHash).
		Scan(&board.BoardID, &board.Name, &board.Theme, &board.Star)

	if err != nil {
		return models.BoardOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoardByURL"}
	}

	_, err = s.db.Exec("INSERT INTO board_members (boardID, userID) VALUES ($1, $2)", boardInput.BoardID, boardInput.UserID)
	if err != nil {
		return models.BoardOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "AddUser"}
	}

	return board, nil
}
