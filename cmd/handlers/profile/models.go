package profileHandler


import (
	"../../../internal/models"
	"github.com/labstack/echo"
)

type profileService interface {
	Profile(request models.UserInput) (user models.User, err error)
	Accounts(request models.UserInput) (user models.User, err error)
	ProfileChange(request models.UserInputProfile) (user models.User, err error)
	AccountsChange(request models.UserInputLinks) (user models.User, err error)
	PasswordChange(request models.UserInputPassword) (user models.User, err error)
}

type profileTransport interface {
	ProfileRead(c echo.Context) (request models.UserInput, err error)

	ProfileChangeRead(c echo.Context) (request models.UserInputProfile, err error)
	AccountsChangeRead(c echo.Context) (request models.UserInputLinks, err error)
	PasswordChangeRead(c echo.Context) (request models.UserInputPassword, err error)

	AccountsWrite(user models.User) (response models.ResponseUserLinks, err error)
	ProfileWrite(user models.User) (response models.ResponseUser, err error)
}

/*type errorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) (err error)
	ServeFatalError(ctx *fasthttp.RequestCtx)
}*/
