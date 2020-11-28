package cardsStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(cardInput models.CardInput) (models.CardInternal, error)
	DeleteCard(cardInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardInternal, error)
	GetCardByID(cardInput models.CardInput) (models.CardInternal, error)
	ChangeCardOrder(taskInput models.CardsOrderInput) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateCard(cardInput models.CardInput) (models.CardOutside, error) {
	var cardID int64

	err := s.db.QueryRow("INSERT INTO cards (boardID, cardName, cardOrder) VALUES ($1, $2, $3) RETURNING cardID",
								cardInput.BoardID, cardInput.Name, cardInput.Order).Scan(&cardID)

	if err != nil {
		fmt.Println(err)
		return models.CardOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateCard"}
	}
	
	card := models.CardOutside{
		CardID: cardID,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskOutsideShort{},
	}
	return card, nil
}

func (s *storage) ChangeCard(cardInput models.CardInput) (models.CardInternal, error) {
	_, err := s.db.Exec("UPDATE cards SET cardName = $1, cardOrder = $2 WHERE cardID = $3", cardInput.Name, cardInput.Order, cardInput.CardID)
	if err != nil {
		fmt.Println(err)
		return models.CardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeCard"}
	}

	card := models.CardInternal{
		CardID: cardInput.CardID,
		Name: cardInput.Name,
		Order: cardInput.Order,
	}
	return card, nil
}

func (s *storage) DeleteCard(cardInput models.CardInput) error {
	_, err := s.db.Exec("DELETE FROM cards WHERE cardID = $1", cardInput.CardID)
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteCard"}
	}

	return nil
}

func (s *storage) GetCardsByBoard(boardInput models.BoardInput) ([]models.CardInternal, error) {
	cards := make([]models.CardInternal, 0)

	rows, err := s.db.Query("SELECT cardID, cardName, cardOrder FROM cards WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		return []models.CardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetCardsByBoard"}
	}
	defer rows.Close()

	for rows.Next() {
		var card models.CardInternal

		err = rows.Scan(&card.CardID, &card.Name, &card.Order)
		if err != nil {
			return []models.CardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetCardsByBoard"}
		}

		cards = append(cards, card)
	}
	return cards, nil
}

func (s *storage) GetCardByID(cardInput models.CardInput) (models.CardInternal, error) {
	card := models.CardInternal{}
	card.CardID = cardInput.CardID

	err := s.db.QueryRow("SELECT cardName, cardOrder FROM cards WHERE cardID = $1", cardInput.CardID).
					Scan(&card.Name, &card.Order)

	if err != nil {
		fmt.Println(err)
		return models.CardInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetCardByID"}
	}

	return card, nil
}

func (s *storage) ChangeCardOrder(cardInput models.CardsOrderInput) error {
	for _, card := range cardInput.Cards {
			_, err := s.db.Exec("UPDATE cards SET cardOrder = $1 WHERE cardID = $2",
				card.Order, card.CardID)

			if err != nil {
				fmt.Println(err)
				return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "ChangeCardOrder"}
		}
	}

	return nil
}