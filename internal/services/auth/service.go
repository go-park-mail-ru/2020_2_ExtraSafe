package auth

import (
	"../../../internal/models"
)

type Service interface {
	Auth(request models.UserInput) (err error)
	Login(request models.UserInputLogin) (response models.User, err error)
	Logout(request models.UserInput) (err error)
	Registration(request models.UserInputReg) (response models.User, err error)
}

type service struct {
	userStorage userStorage
}

func NewService(userStorage userStorage) Service {
	return &service{
		userStorage: userStorage,
	}
}

func (s *service)Auth(request models.UserInput) (err error) {
	return err
}

func (s *service)Login(request models.UserInputLogin) (response models.User, err error) {
	var user models.User
	user, err = s.userStorage.CheckUser(request)
	if err != nil {
		return models.User{}, err
	}

	//setCookie(c, user.ID)

	return user, err
}

func (s *service)Logout(request models.UserInput) (err error) {
	/*cc := c.(*Handlers)

	session, err := c.Cookie("tabutask_id")
	if err != nil {
		return err
	}
	sessionID := session.Value

	delete(*cc.Sessions, sessionID)
	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(session)*/
	return nil
}

func (s *service)Registration(request models.UserInputReg) (response models.User, err error) {
	var user models.User
	user, err = s.userStorage.CreateUser(request)
	if err != nil {
		return models.User{}, err
	}

	//setCookie(c, user.ID)

	return user, err
}
