package cardsStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(cardInput models.CardInput) (models.CardOutside, error)
	DeleteCard(cardInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardOutside, error)
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
	var cardID uint64 = 0
	var quantityCards uint64 = 0

	//FIXME сделать по-другому поиск последнего ID
	err := s.db.QueryRow("SELECT COUNT(*) FROM cards WHERE boardID = $1", cardInput.BoardID).Scan(&quantityCards)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return models.CardOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	cardID = quantityCards + 1

	_, err = s.db.Exec("INSERT INTO cards (cardID, boardID, cardName, cardOrder) VALUES ($1, $2, $3, $4)", cardID, cardInput.BoardID, cardInput.Name, cardInput.Order)

	if err != nil {
		//TODO в таких местах надо возвращать internal error (или как-то так), и записывать ошибку в лог
		fmt.Println(err) 
		return models.CardOutside{} ,models.ServeError{Codes: []string{"500"}}
	}
	
	card := models.CardOutside{
		CardID: cardID,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskOutside{},
	}
	return card, nil
}

func (s *storage) ChangeCard(cardInput models.CardInput) (models.CardOutside, error) {
	_, err := s.db.Exec("UPDATE cards SET cardName = $1, cardOrder = $2 WHERE cardID = $3", cardInput.Name, cardInput.Order)
	if err != nil {
		fmt.Println(err)
		return models.CardOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	card := models.CardOutside{
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
		return models.ServeError{Codes: []string{"500"}}
	}

	return nil
}

func (s *storage) GetCardsByBoard(boardInput models.BoardInput) ([]models.CardOutside, error) {
	cards := make([]models.CardOutside, 0)

	rows, err := s.db.Query("SELECT cardID, cardName, cardOrder FROM cards WHERE boardID = $1", boardInput.BoardID)
	if err != nil {
		return []models.CardOutside{}, models.ServeError{Codes: []string{"500"}}
	}
	defer rows.Close()

	for rows.Next() {
		var card models.CardOutside

		err = rows.Scan(&card.CardID, &card.Name, &card.Order)
		if err != nil {
			return []models.CardOutside{}, models.ServeError{Codes: []string{"500"}}
		}

		cards = append(cards, card)
	}
	return cards, nil
}