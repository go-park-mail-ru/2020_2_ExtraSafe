package sessionsStorage

type Storage interface {
	DeleteUserSession(sessionId string)
	CreateUserSession(userId uint64, SID string)
}

type storage struct {
	Sessions *map[string]uint64
}

func NewStorage(sessions map[string]uint64) Storage {
	return &storage{
		Sessions: &sessions,
	}
}

func (s *storage) DeleteUserSession(sessionId string)  {
	delete(*s.Sessions, sessionId)
}

func (s *storage) CreateUserSession(userId uint64, SID string) {
	(*s.Sessions)[SID] = userId
}