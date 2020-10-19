package handlers

import (
	authHandler "./auth"
	profileHandler "./profile"
	"github.com/labstack/echo"
	"./middlewares"
)

func Router(e *echo.Echo, profile profileHandler.Handler, auth authHandler.Handler, middle middlewares.Middleware) {

	e.Any("/", middle.CookieSession(auth.Auth))
	e.POST("/login/", auth.Login)
	e.GET("/logout/", auth.Logout)
	e.POST("/reg/", auth.Registration)
	e.GET("/profile/", middle.CookieSession(profile.Profile))
	e.GET("/accounts/", middle.CookieSession(profile.Accounts))
	e.Static("/avatar", "")
	e.POST("/profile/", middle.CookieSession(profile.ProfileChange))
	e.POST("/accounts/", middle.CookieSession(profile.AccountsChange))
	e.POST("/password/", middle.CookieSession(profile.PasswordChange))
}