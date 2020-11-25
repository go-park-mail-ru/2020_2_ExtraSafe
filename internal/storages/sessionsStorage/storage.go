package sessionsStorage

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/tarantool/go-tarantool"
)

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
	sessionDataSlice, _ := data.([]interface{})
	userData, _ := sessionDataSlice[1].(uint64)

	return userData, nil
}