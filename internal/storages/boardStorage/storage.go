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
	DeleteCard(userInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error

	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	GetCard(cardInput models.CardInput) (models.CardOutside, error)
	GetTask(taskInput models.TaskInput) (models.TaskOutside, error)

	CheckBoardPermission(userID uint64, boardID uint64, ifAdmin bool) (err error)
	CheckCardPermission(userID uint64, cardID uint64) (err error)
	CheckTaskPermission(userID uint64, taskID uint64) (err error)

	checkBoardAdminPermission(userID uint64, boardID uint64) (flag bool, err error)
	checkBoardUserPermission(userID uint64, boardID uint64) (flag bool, err error)
	getBoardMembers(boardInput models.BoardInput) ([]uint64, error)
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
		return []models.BoardOutsideShort{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var board models.BoardOutsideShort

		err = rows.Scan(&board.BoardID, &board.Name, &board.Theme, &board.Star)
		if err != nil {
			return []models.BoardOutsideShort{}, err
		}

		boards = append(boards, board)
	}
	return boards, nil
}

func (s *storage) CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error) {
	var boardID uint64

	err := s.db.QueryRow("INSERT INTO boards (adminID, boardName, theme, star) VALUES ($1, $2, $3, $4) RETURNING boardID",
		boardInput.UserID,
		boardInput.BoardName,
		boardInput.Theme,
		boardInput.Star).Scan(&boardID)

	if err != nil {
		fmt.Println(err) //TODO error
		return models.BoardInternal{} ,models.ServeError{Codes: []string{"500"}}
	}

	board := models.BoardInternal{
		BoardID:  boardID,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardOutside{},
		UsersIDs: []uint64{},
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
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}}
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

	}

	board.Cards = append(board.Cards, cards...)
	return board, nil
	/*if err != nil && err != sql.ErrNoRows {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}}
	}

	if err == sql.ErrNoRows {
		return models.BoardInternal{}, models.ServeError{Codes: []string{"not existing board"}}
	}*/
}

func (s *storage) getBoardMembers(boardInput models.BoardInput) ([]uint64, error) {
	members := make([]uint64, 0)

	rows, err := s.db.Query("SELECT userID from board_members WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		return []uint64{}, models.ServeError{Codes: []string{"500"}}
	}
	defer rows.Close()

	for rows.Next() {
		var userID uint64

		err = rows.Scan(&userID)
		if err != nil {
			return []uint64{}, models.ServeError{Codes: []string{"500"}}
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
		return models.BoardInternal{}, models.ServeError{Codes: []string{"500"}}
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
		return models.ServeError{Codes: []string{"500"}}
	}

	return nil
}

func (s *storage) CreateCard(cardInput models.CardInput) (models.CardOutside, error) {
	card, err := s.cardsStorage.CreateCard(cardInput)
	if err != nil {
		return models.CardOutside{}, err
	}

	//не ищем таски, потому что при создании доски они пустые
	return card, nil
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

func (s *storage) DeleteCard(cardInput models.CardInput) error {
	err := s.cardsStorage.DeleteCard(cardInput)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) CreateTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	task, err := s.tasksStorage.CreateTask(taskInput)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return task, nil
}

func (s *storage) ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	task, err := s.tasksStorage.ChangeTask(taskInput)
	if err != nil {
		return models.TaskOutside{}, nil
	}
	return task, nil
}

func (s *storage) DeleteTask(taskInput models.TaskInput) error {
	err := s.tasksStorage.DeleteTask(taskInput)
	if err != nil {
		return err
	}

	return nil
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
	task, err := s.tasksStorage.GetTaskByID(taskInput)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return task, nil
}

func (s *storage) CheckBoardPermission(userID uint64, boardID uint64, ifAdmin bool) (err error) {
	var ok bool

	if ifAdmin {
		ok, err = s.checkBoardAdminPermission(userID, boardID)
	} else {
		ok, err = s.checkBoardUserPermission(userID, boardID)
	}

	if err != nil {
		fmt.Println(err)	//log
		return models.ServeError{Codes: []string{"500"}}
	}

	if !ok {
		return models.ServeError{Codes: []string{"403"}}
	}

	return nil
}

func (s *storage) checkBoardAdminPermission(userID uint64, boardID uint64) (flag bool, err error) {
	var ID uint64
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

func (s *storage) checkBoardUserPermission(userID uint64, boardID uint64) (flag bool, err error) {
	var ID uint64
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

func (s *storage) CheckCardPermission(userID uint64, cardID uint64) (err error) {
	var boardID uint64

	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR  M.userID = $1) AND cardID = $2", userID, cardID).Scan(&boardID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}}
	}

	return nil
}

func (s *storage) CheckTaskPermission(userID uint64, taskID uint64) (err error) {
	var boardID uint64

	err = s.db.QueryRow("SELECT B.boardID FROM boards B " +
								"JOIN cards C on C.boardID = B.boardID " +
								"JOIN tasks T on T.cardID = C.cardID " +
								"LEFT OUTER JOIN board_members M ON B.boardID = M.boardID " +
								"WHERE (B.adminID = $1 OR  M.userID = $1) AND taskID = $2", userID, taskID).Scan(&boardID)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}}
	}

	if err == sql.ErrNoRows {
		return models.ServeError{Codes: []string{"403"}}
	}

	return nil
}