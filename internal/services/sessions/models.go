package sessions

type sessionsStorage interface {
	DeleteUserSession(sessionId string) error
	CreateUserSession(userId int64, SID string) error
	CheckUserSession(sessionId string) (int64, error)
}
