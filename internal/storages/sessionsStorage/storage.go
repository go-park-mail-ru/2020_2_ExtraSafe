package sessionsStorage

type Storage interface {
	DeleteUserSession(sessionId string)
	CreateUserSession(userId uint64, SID string)
	CheckUserSession(sessionId string) (uint64, bool)
}

type storage struct {
	Sessions *map[string]uint64
}

func NewStorage(sessions *map[string]uint64) Storage {
	return &storage{
		Sessions: sessions,
	}
}

func (s *storage) DeleteUserSession(sessionId string)  {
	delete(*s.Sessions, sessionId)
}

func (s *storage) CreateUserSession(userId uint64, SID string) {
	(*s.Sessions)[SID] = userId
}

func (s *storage) CheckUserSession(sessionId string) (uint64, bool) {
	userId, authorized := (*s.Sessions)[sessionId]
	return userId, authorized
}