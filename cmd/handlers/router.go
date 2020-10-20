package handlers

import (
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, profile profileHandler.Handler, auth authHandler.Handler, middle middlewares.Middleware) {

	e.Any("/", middle.CookieSession(auth.Auth))
	e.POST("/login/", middle.AuthCookieSession(auth.Login))
	e.GET("/logout/", auth.Logout)
	e.POST("/reg/", middle.AuthCookieSession(auth.Registration))
	e.GET("/profile/", middle.CookieSession(profile.Profile))
	e.GET("/accounts/", middle.CookieSession(profile.Accounts))
	e.Static("/avatar", "../")
	e.POST("/profile/", middle.CookieSession(profile.ProfileChange))
	e.POST("/accounts/", middle.CookieSession(profile.AccountsChange))
	e.POST("/password/", middle.CookieSession(profile.PasswordChange))
}