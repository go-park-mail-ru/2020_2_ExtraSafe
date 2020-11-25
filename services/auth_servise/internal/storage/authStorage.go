package storage

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/tarantool/go-tarantool"
)

//go:generate mockgen -destination=./mock/mock_sessionsStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/sessionsStorage Storage

type Storage interface {
	DeleteUserSession(sessionId string) error
	CreateUserSession(userId uint64, SID string) error
	CheckUserSession(sessionId string) (uint64, error)
}

type storage struct {
	Sessions *tarantool.Connection
}

func NewStorage(sessions *tarantool.Connection) Storage {
	return &storage{
		Sessions: sessions,
	}
}

func (s *storage) DeleteUserSession(sessionId string)  error {
	_, err := s.Sessions.Delete("sessions", "primary", []interface{}{sessionId})
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteUserSession"}
	}
	return nil
}

func (s *storage) CreateUserSession(userId uint64, SID string) error {
	_, err := s.Sessions.Insert("sessions", []interface{}{SID, userId})
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CreateUserSession"}
	}
	return nil
}

func (s *storage) CheckUserSession(sessionId string) (uint64, error) {

	resp, err := s.Sessions.Select("sessions", "primary", 0, 1, tarantool.IterEq, []interface{}{sessionId})

	if err != nil {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckUserSession"}
	}

	if len(resp.Tuples()) == 0 {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("Empty sessions DB "),
			MethodName: "CheckUserSession"}
	}

	data := resp.Data[0]
	sessionDataSlice, ok := data.([]interface{})
	if !ok {
		return 0, models.ServeError{Codes: []string{"500"},
			OriginalError: fmt.Errorf("cannot cast data: %v", sessionDataSlice),
			MethodName: "CheckUserSession"}
	}

	userData, ok := sessionDataSlice[1].(uint64)
	if !ok {
		return 0, models.ServeError{Codes: []string{"500"},
			OriginalError: fmt.Errorf("cannot cast data: %v", sessionDataSlice[1]),
			MethodName: "CheckUserSession"}
	}

	return userData, nil
}