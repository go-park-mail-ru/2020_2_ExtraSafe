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

	e.Any("/api/", middle.CookieSession(auth.Auth))
	e.POST("/api/login/", middle.AuthCookieSession(auth.Login))
	e.GET("/api/logout/", auth.Logout)
	e.POST("/api/reg/", middle.AuthCookieSession(auth.Registration))

	e.GET("/api/profile/", middle.CookieSession(profile.Profile))
	e.GET("/api/boards/", middle.CookieSession(profile.Boards))
	e.POST("/api/profile/", middle.CookieSession(middle.CSRFToken(profile.ProfileChange)))
	e.POST("/api/password/", middle.CookieSession(middle.CSRFToken(profile.PasswordChange)))

	e.Static("/static/avatar", "../")
	e.Static("/static/files", "../")

	e.GET("/api/ws/notification/", middle.CookieSession(board.Notification))
	e.GET("/api/ws/board/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.BoardWS)))

	e.GET("/api/board/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.Board)))
	e.POST("/api/board/", middle.CookieSession(middle.CSRFToken(board.BoardCreate)))
	e.PUT("/api/board/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardChange))))
	e.DELETE("/api/board/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardDelete))))
	e.PUT("/api/board/:ID/user-add/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.BoardAddMember))))
	e.PUT("/api/board/:ID/user-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.BoardRemoveMember))))

	e.GET("/api/card/:ID/", middle.CookieSession(middle.CheckCardUserPermission(board.Card)))
	e.POST("/api/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.CardCreate))))
	e.PUT("/api/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCardUserPermission(board.CardChange))))
	e.DELETE("/api/card/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCardUserPermission(board.CardDelete))))

	e.GET("/api/task/:ID/", middle.CookieSession(middle.CheckTaskUserPermission(board.Task)))
	e.POST("/api/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TaskCreate))))
	e.PUT("/api/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskChange))))
	e.DELETE("/api/task/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskDelete))))

	e.POST("/api/task-order/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.TaskOrder)))
	e.POST("/api/card-order/:ID/", middle.CookieSession(middle.CheckBoardUserPermission(board.CardOrder)))

	e.POST("/api/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagCreate))))
	e.PUT("/api/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagChange))))
	e.DELETE("/api/tag/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardUserPermission(board.TagDelete))))
	e.PUT("/api/task/:ID/tag-add/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TagAdd))))
	e.PUT("/api/task/:ID/tag-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TagRemove))))
	e.PUT("/api/task/:ID/user-add/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskUserAdd))))
	e.PUT("/api/task/:ID/user-remove/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.TaskUserRemove))))

	e.POST("/api/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.CommentCreate))))
	e.PUT("/api/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCommentUpdateUserPermission(board.CommentChange))))
	e.DELETE("/api/comment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckCommentDeleteUserPermission(board.CommentDelete))))

	e.POST("/api/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistCreate))))
	e.PUT("/api/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistChange))))
	e.DELETE("/api/checklist/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.ChecklistDelete))))

	e.POST("/api/attachment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.AttachmentCreate))))
	e.DELETE("/api/attachment/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckTaskUserPermission(board.AttachmentDelete))))

	e.GET("/api/shared-url/:ID/", middle.CookieSession(middle.CSRFToken(middle.CheckBoardAdminPermission(board.SharedURL))))
	e.GET("/api/invite/board/:ID/:url/", middle.CookieSession(board.BoardInvite))
}