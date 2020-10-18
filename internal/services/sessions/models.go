package sessions

type sessionsStorage interface {
	DeleteUserSession(sessionId string)
	CreateUserSession(userId uint64, SID string)
}
