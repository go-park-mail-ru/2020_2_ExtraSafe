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

	e.GET("/board/:boardID/", middle.CookieSession(board.Board))
	e.POST("/board/", middle.CookieSession(board.BoardCreate))
	e.PUT("/board/:boardID/", middle.CookieSession(board.BoardChange))
	e.DELETE("/board/:boardID/", middle.CookieSession(board.BoardDelete))

	e.GET("/card/:cardID/", middle.CookieSession(board.Board))
	e.POST("/card/", middle.CookieSession(board.CardCreate))
	e.PUT("/card/:cardID/", middle.CookieSession(board.CardChange))
	e.DELETE("/card/:cardID/", middle.CookieSession(board.CardDelete))

	e.GET("/task/:taskID/", middle.CookieSession(board.Board))
	e.POST("/task/", middle.CookieSession(board.TaskCreate))
	e.PUT("/task/:taskID/", middle.CookieSession(board.TaskChange))
	e.DELETE("/task/:taskID/", middle.CookieSession(board.TaskDelete))
}