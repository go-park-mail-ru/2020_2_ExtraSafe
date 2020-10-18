package handlers

import (
	authHandler "./auth"
	profileHandler "./profile"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, profile profileHandler.Handler, auth authHandler.Handler) {
	e.Any("/", auth.Auth)
	e.POST("/login/", auth.Login)
	e.GET("/logout/", auth.Logout)
	e.POST("/reg/", auth.Registration)
	e.GET("/profile/", profile.Profile)
	e.GET("/accounts/", profile.Accounts)
	e.Static("/avatar", "")
	e.POST("/profile/", profile.ProfileChange)
	e.POST("/accounts/", profile.AccountsChange)
	e.POST("/password/", profile.PasswordChange)
}