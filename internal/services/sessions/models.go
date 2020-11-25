package sessions

//go:generate mockgen -destination=./mock/mock_sessionStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/sessions SessionStorage
////go:generate mockgen -destination=./mock/mock_boardStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile BoardStorage
type SessionStorage interface {
	DeleteUserSession(sessionId string) error
	CreateUserSession(userId int64, SID string) error
	CheckUserSession(sessionId string) (int64, error)
}
