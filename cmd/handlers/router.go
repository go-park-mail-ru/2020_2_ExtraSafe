package handlers

import (
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	boardsHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, profile profileHandler.Handler, auth authHandler.Handler, board boardsHandler.Handler,
	middle middlewares.Middleware) {

	e.Any("/", middle.CookieSession(auth.Auth))
	e.POST("/login/", middle.AuthCookieSession(auth.Login))
	e.GET("/logout/", auth.Logout)
	e.POST("/reg/", middle.AuthCookieSession(auth.Registration))

	e.GET("/profile/", middle.CookieSession(profile.Profile))
	e.GET("/accounts/", middle.CookieSession(profile.Accounts))
	e.POST("/profile/", middle.CookieSession(profile.ProfileChange))
	e.POST("/accounts/", middle.CookieSession(profile.AccountsChange))
	e.POST("/password/", middle.CookieSession(profile.PasswordChange))

	e.Static("/avatar", "../")

	e.GET("/board/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.Board)))
	e.POST("/board/", middle.CookieSession(board.BoardCreate))
	e.PUT("/board/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.BoardChange)))
	e.DELETE("/board/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.BoardDelete)))

	e.GET("/card/:ID/", middle.CookieSession(middle.CheckCardUserPermission(board.Card)))
	e.POST("/card/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.CardCreate)))
	e.PUT("/card/:ID/", middle.CookieSession(middle.CheckCardUserPermission(board.CardChange)))
	e.DELETE("/card/:ID/", middle.CookieSession(middle.CheckCardUserPermission(board.CardDelete)))

	e.GET("/task/:ID/", middle.CookieSession(middle.CheckTaskUserPermission(board.Task)))
	e.POST("/task/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.TaskCreate)))
	e.PUT("/task/:ID/", middle.CookieSession(middle.CheckTaskUserPermission(board.TaskChange)))
	e.DELETE("/task/:ID/", middle.CookieSession(middle.CheckTaskUserPermission(board.TaskDelete)))
}