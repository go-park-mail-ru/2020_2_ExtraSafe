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

	e.Use(middle.Logger)

	e.Any("/", middle.CookieSession(auth.Auth))
	e.POST("/login/", middle.AuthCookieSession(auth.Login))
	e.GET("/logout/", auth.Logout)
	e.POST("/reg/", middle.AuthCookieSession(auth.Registration))

	e.GET("/profile/", middle.CookieSession(profile.Profile))
	e.GET("/accounts/", middle.CookieSession(profile.Accounts))
	e.GET("/boards/", middle.CookieSession(profile.Boards))
	e.POST("/profile/", middle.CookieSession(middle.CSRFToken(profile.ProfileChange)))
	e.POST("/accounts/", middle.CookieSession(middle.CSRFToken(profile.AccountsChange)))
	e.POST("/password/", middle.CookieSession(middle.CSRFToken(profile.PasswordChange)))

	e.Static("/avatar", "../")
	e.Static("/files", "../")

	e.GET("/board/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.Board)))
	e.POST("/board/", middle.CookieSession(middle.CSRFToken(board.BoardCreate)))
	e.PUT("/board/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardChange))))
	e.DELETE("/board/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardDelete))))
	e.PUT("/board/:ID/user-add/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardAddMember))))
	e.PUT("/board/:ID/user-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardRemoveMember))))

	e.GET("/card/:ID/", middle.CookieSession(middle.CheckCardUserPermission(board.Card)))
	e.POST("/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.CardCreate))))
	e.PUT("/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCardUserPermission(board.CardChange))))
	e.DELETE("/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCardUserPermission(board.CardDelete))))

	e.GET("/task/:ID/", middle.CookieSession(middle.CheckTaskUserPermission(board.Task)))
	e.POST("/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TaskCreate))))
	e.PUT("/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskChange))))
	e.DELETE("/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskDelete))))

	e.POST("/task-order/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.TaskOrder)))
	e.POST("/card-order/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.CardOrder)))

	e.POST("/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagCreate))))
	e.PUT("/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagChange))))
	e.DELETE("/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagDelete))))
	e.PUT("/task/:ID/tag-add/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TagAdd))))
	e.PUT("/task/:ID/tag-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TagRemove))))
	e.PUT("/task/:ID/user-add/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskUserAdd))))
	e.PUT("/task/:ID/user-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskUserRemove))))

	e.POST("/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.CommentCreate))))
	e.PUT("/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCommentUpdateUserPermission(board.CommentChange))))
	e.DELETE("/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCommentDeleteUserPermission(board.CommentDelete))))

	e.POST("/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistCreate))))
	e.PUT("/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistChange))))
	e.DELETE("/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistDelete))))

	e.POST("/attachment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.AttachmentCreate))))
	e.DELETE("/attachment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.AttachmentDelete))))
}