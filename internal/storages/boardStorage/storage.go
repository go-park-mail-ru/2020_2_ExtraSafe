package boardStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
)

type Storage interface {
	GetBoardsList(userInput models.UserInput) ([] models.BoardOutside, error)

	//FIXME сменить userInput -> boardInput
	CreateBoard(userInput models.BoardChangeInput) (models.BoardOutside, error)
	GetBoard(userInput models.BoardInput) (models.BoardOutside, error)
	ChangeBoard(userInput models.BoardChangeInput) (models.BoardOutside, error)
	DeleteBoard(userInput models.BoardInput) error

	CreateColumn(userInput models.CardInput) (models.CardOutside, error)
	ChangeColumn(userInput models.CardInput) (models.CardOutside, error)
	DeleteColumn(userInput models.CardInput) error

	CreateTask(userInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(userInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(userInput models.TaskInput) error
}

type storage struct {
	db *sql.DB
	cardsStorage cardsStorage
	tasksStorage tasksStorage
}

func NewStorage(db *sql.DB, cardsStorage cardsStorage, tasksStorage tasksStorage) Storage {
	return &storage{
		db: db,
		cardsStorage: cardsStorage,
		tasksStorage: tasksStorage,
	}
}

func (s *storage) GetBoardsList(userInput models.UserInput) ([] models.BoardOutside, error) {
	// взять по айди пользователя список таблиц (из members и admins)
	// узнать, надо ли возвращать по что-то кроме айди таблицы (наверное еще имя)
	// запрос с JOIN

	boards := make([]models.BoardOutside, 0)

	rows, err := s.db.Query("SELECT boardID FROM boards_members WHERE userID = $1", userInput.ID)
	if err != nil {
		return []models.BoardOutside{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var boardID uint64

		err = rows.Scan(&boardID)
		if err != nil {
			return []models.BoardOutside{}, err
		}

		//TODO заполнить все поля структуры ?
		//TODO переделать на BoardsOutputShort
		currentBoard := models.BoardOutside{
			BoardID: boardID,
			Admin:   models.User{},
			Name:    "",
			Theme:   "",
			Star:    false,
			Users:   nil,
			Columns: nil,
		}

		boards = append(boards, currentBoard)
	}

	rows, err = s.db.Query("SELECT boardID FROM boards WHERE adminID = $1", userInput.ID)
	if err != nil {
		return []models.BoardOutside{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var boardID uint64

		err = rows.Scan(&boardID)
		if err != nil {
			return []models.BoardOutside{}, err
		}

		//TODO заполнить все поля структуры
		//TODO найти admins ?
		currentBoard := models.BoardOutside{
			BoardID: boardID,
			Admin:   models.User{},
			Name:    "",
			Theme:   "",
			Star:    false,
			Users:   nil,
			Columns: nil,
		}

		boards = append(boards, currentBoard)
	}
	return boards, nil
}

func (s *storage) CreateBoard(userInput models.BoardChangeInput) (models.BoardOutside, error) {
	var boardID uint64 = 0
	var quantityBoards uint64 = 0
	//TODO сделать по-другому поиск последнего ID
	err := s.db.QueryRow("SELECT COUNT(*) FROM boards").Scan(&quantityBoards)
	boardID = quantityBoards + 1

	_, err = s.db.Exec("INSERT INTO boards (boardID, adminID, boardName, theme, star) VALUES ($1, $2, $3, $4, $5)",
		boardID,
		userInput.UserID,
		userInput.BoardName,
		userInput.Theme,
		userInput.Star)

	if err != nil {
		//TODO разработать код ошибок на стороне БД
		fmt.Println(err)
		return models.BoardOutside{} ,models.ServeError{Codes: []string{"500"}}
	}

	board := models.BoardOutside{
		BoardID: boardID,
		//TODO заполнение структуры админа?
		//Admin:   userInput.UserID,
		Name:    userInput.BoardName,
		Theme:   userInput.Theme,
		Star:    userInput.Star,
		Users:   []models.UserOutside{},
		Columns: []models.CardOutside{},
	}

	return board, nil
}

func (s *storage) GetBoard(userInput models.BoardInput) (models.BoardOutside, error) {
	board := models.BoardOutside{}
	var adminID uint64
	rows, err := s.db.Query("SELECT adminID, boardName, theme, star FROM boards JOIN board_members USING (boardID) WHERE userID = $1", userInput.BoardID)
	if err != nil {
		return models.BoardOutside{}, err
	}
	defer rows.Close()

		//Scan(&adminID, &board.Name, &board.Theme, &board.Star)


	for rows.Next() {
		var networkName, link string

		err = rows.Scan(&networkName, &link)
		if err != nil {
			return models.UserOutside{}, err
		}

		//TODO поиграться с рефлектами
		reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkName).SetString(link)
	}
	return models.BoardOutside{}, nil
}

func (s *storage) ChangeBoard(userInput models.BoardChangeInput) (models.BoardOutside, error) {
	return models.BoardOutside{}, nil
}

func (s *storage) DeleteBoard(userInput models.BoardInput) error { return nil }

func (s *storage) CreateColumn(userInput models.CardInput) (models.CardOutside, error) {
	return models.CardOutside{}, nil
}
func (s *storage) ChangeColumn(userInput models.CardInput) (models.CardOutside, error) {
	return models.CardOutside{}, nil
}
func (s *storage) DeleteColumn(userInput models.CardInput) error {
	return nil
}

func (s *storage) CreateTask(userInput models.TaskInput) (models.TaskOutside, error) {
	return models.TaskOutside{}, nil
}
func (s *storage) ChangeTask(userInput models.TaskInput) (models.TaskOutside, error) {
	return models.TaskOutside{}, nil
}
func (s *storage) DeleteTask(userInput models.TaskInput) error {
	return nil
}