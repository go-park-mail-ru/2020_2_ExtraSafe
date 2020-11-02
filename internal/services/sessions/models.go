package sessions

type sessionsStorage interface {
	DeleteUserSession(sessionId string) error
	CreateUserSession(userId uint64, SID string) error
	CheckUserSession(sessionId string) (uint64, error)
}
