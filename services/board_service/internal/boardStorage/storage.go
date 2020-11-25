package boardStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)

	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error

	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardOutside, error)
	ChangeCardOrder(cardInput models.CardsOrderInput) error
	DeleteCard(userInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
	DeleteTask(taskInput models.TaskInput) error

	AddUser(input models.BoardMember) (err error)
	RemoveUser(input models.BoardMember) (err error)

	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	GetCard(cardInput models.CardInput) (models.CardOutside, error)
	GetTask(taskInput models.TaskInput) (models.TaskOutside, error)

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (err error)
	CheckTaskPermission(userID int64, taskID int64) (err error)
}

type storage struct {
	db           *sql.DB
	cardsStorage CardsStorage
	tasksStorage TasksStorage
}

func NewStorage(db *sql.DB, cardsStorage CardsStorage, tasksStorage TasksStorage) Storage {
	return &storage{
		db: db,
		cardsStorage: cardsStorage,
		tasksStorage: tasksStorage,
	}
}

func (s *storage) GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error) {
	boards := make([]models.BoardOutsideShort, 0)

	rows, err := s.db.Query("SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B " +
									"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID WHERE B.adminID = $1 OR  M.userID = $1;", userInput.ID)
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
		fmt.Println(err)
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateBoard"}
	}

	board := models.BoardInternal{
		BoardID:  boardID,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardOutside{},
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

		card.Tasks = append(card.Tasks, tasks...)
		board.Cards = append(board.Cards, card)
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
		fmt.Println(err)
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeBoard"}
	}

	members, err := s.getBoardMembers(models.BoardInput{ BoardID: boardInput.BoardID })
	if err != nil {
		return models.BoardInternal{}, err
	}

	cards, err := s.cardsStorage.GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID})
	if err != nil {
		return models.BoardInternal{}, err
	}

	for _, card := range cards {
		cardInput := models.CardInput{CardID: card.CardID}

		tasks, err := s.tasksStorage.GetTasksByCard(cardInput)
		if err != nil {
			return models.BoardInternal{}, err
		}

		card.Tasks = append(card.Tasks, tasks...)
	}

	board.BoardID = boardInput.BoardID
	board.Name = boardInput.BoardName
	board.Theme = boardInput.Theme
	board.Star = boardInput.Star
	board.UsersIDs = append(board.UsersIDs, members...)
	board.Cards = append(board.Cards, cards...)

	return board, nil
}

func (s *storage) DeleteBoard(boardInput models.BoardInput) error {
	_, err := s.db.Exec("DELETE FROM boards WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteBoard"}
	}

	return nil
}

func (s *storage) CreateCard(cardInput models.CardInput) (models.CardOutside, error) {
	//card, err :=
	//не ищем таски, потому что при создании доски они пустые
	return s.cardsStorage.CreateCard(cardInput)
}

func (s *storage) ChangeCard(cardInput models.CardInput) (models.CardOutside, error) {
	card, err := s.cardsStorage.ChangeCard(cardInput)
	if err != nil {
		return models.CardOutside{}, err
	}

	tasks, err := s.tasksStorage.GetTasksByCard(cardInput)
	if err != nil {
		return models.CardOutside{}, err
	}

	card.Tasks = append(card.Tasks, tasks...)
	return card, nil
}

func (s *storage) ChangeCardOrder(cardInput models.CardsOrderInput) error {
	return s.cardsStorage.ChangeCardOrder(cardInput)
}

func (s *storage) DeleteCard(cardInput models.CardInput) error {
	return s.cardsStorage.DeleteCard(cardInput)
}

func (s *storage) CreateTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	return s.tasksStorage.CreateTask(taskInput)
}

func (s *storage) ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	return s.tasksStorage.ChangeTask(taskInput)
}

func (s *storage) ChangeTaskOrder(taskInput models.TasksOrderInput) error {
	return s.tasksStorage.ChangeTaskOrder(taskInput)
}

func (s *storage) DeleteTask(taskInput models.TaskInput) error {
	return s.tasksStorage.DeleteTask(taskInput)
}

func (s *storage) GetCard(cardInput models.CardInput) (models.CardOutside, error) {
	card, err := s.cardsStorage.GetCardByID(cardInput)
	if err != nil {
		return models.CardOutside{}, err
	}

	tasks, err := s.tasksStorage.GetTasksByCard(cardInput)
	if err != nil {
		return models.CardOutside{}, err
	}

	card.Tasks = append(card.Tasks, tasks...)
	return card, nil
}

func (s *storage) GetTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	return s.tasksStorage.GetTaskByID(taskInput)
}

func (s *storage) CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error) {
	var ok bool

	if ifAdmin {
		ok, err = s.checkBoardAdminPermission(userID, boardID)
	} else {
		ok, err = s.checkBoardUserPermission(userID, boardID)
	}

	if err != nil {
		fmt.Println(err)	//log
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
	err = s.db.QueryRow("SELECT boardID FROM board_members WHERE boardID = $1 AND userID = $2", boardID, userID).Scan(&ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
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
								"WHERE (B.adminID = $1 OR  M.userID = $1) AND cardID = $2", userID, cardID).Scan(&boardID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
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
								"WHERE (B.adminID = $1 OR  M.userID = $1) AND taskID = $2", userID, taskID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckTaskPermission"}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}, Descriptions: []string{"Permissions denied"},
			MethodName: "CheckTaskPermission"}
	}

	return nil
}
