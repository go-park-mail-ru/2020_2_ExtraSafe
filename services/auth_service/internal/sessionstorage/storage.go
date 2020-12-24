package sessionstorage

//go:generate mockgen -destination=../../../mock/mock_sessionsStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_service/internal/sessionsStorage SessionStorage

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/tarantool/go-tarantool"
)

type SessionStorage interface {
	DeleteUserSession(sessionID string) error
	CreateUserSession(userID int64, SID string) error
	CheckUserSession(sessionID string) (int64, error)
}

type storage struct {
	Sessions *tarantool.Connection
}

func NewStorage(sessions *tarantool.Connection) SessionStorage {
	return &storage{
		Sessions: sessions,
	}
}

func (s *storage) DeleteUserSession(sessionID string) error {
	_, err := s.Sessions.Delete("sessions", "primary", []interface{}{sessionID})
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteUserSession"}
	}
	return nil
}

func (s *storage) CreateUserSession(userId int64, sID string) error {
	_, err := s.Sessions.Insert("sessions", []interface{}{sID, userId})
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CreateUserSession"}
	}
	return nil
}

func (s *storage) CheckUserSession(sessionID string) (int64, error) {

	resp, err := s.Sessions.Select("sessions", "primary", 0, 1, tarantool.IterEq, []interface{}{sessionID})

	if err != nil {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckUserSession"}
	}

	if len(resp.Tuples()) == 0 {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("No such session "),
			MethodName: "CheckUserSession"}
	}

	data := resp.Data[0]
	sessionDataSlice, _ := data.([]interface{})
	userData, _ := sessionDataSlice[1].(uint64)

	return int64(userData), nil
}
