package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/csrf"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorworker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

type Middleware interface {
	CORS() echo.MiddlewareFunc
	CookieSession(next echo.HandlerFunc) echo.HandlerFunc
	AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc
	CSRFToken(next echo.HandlerFunc) echo.HandlerFunc
	Logger(next echo.HandlerFunc) echo.HandlerFunc
	CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckBoardAdminPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckCommentUpdateUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckCommentDeleteUserPermission(next echo.HandlerFunc) echo.HandlerFunc
}

type middlew struct {
	errorWorker    errorworker.ErrorWorker
	authService    auth.ServiceAuth
	authTransport  auth.TransportAuth
	boardService   boards.ServiceBoard
	internalLogger logger.Logger
}

func NewMiddleware(errorWorker errorworker.ErrorWorker, authService auth.ServiceAuth, authTransport auth.TransportAuth,
	boardService boards.ServiceBoard, internalLogger logger.Logger) Middleware {
	return middlew{
		errorWorker:    errorWorker,
		authService:    authService,
		authTransport:  authTransport,
		boardService:   boardService,
		internalLogger: internalLogger,
	}
}

func (m middlew) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://95.163.213.142:80",
			"http://95.163.213.142",
			"http://127.0.0.1:3033",
			"http://127.0.0.1:63246",
			"http://127.0.0.1",
			"http://tabutask.ru",
			"https://tabutask.ru",
			"https://127.0.0.1"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-CSRF-Token"},
		AllowCredentials: true,
	})
}

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := m.authService.CheckCookie(c)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		session, _ := c.Cookie("tabutask_id")

		c.Set("userId", userID)
		c.Set("sessionID", session.Value)

		return next(c)
	}
}

func (m middlew) AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := m.authService.CheckCookie(c)
		if err == nil {
			userInput := new(models.UserInput)
			userInput.ID = userID
			user, _ := m.authService.Auth(*userInput)
			token, _ := csrf.GenerateToken(userID)
			response, err := m.authTransport.AuthWrite(user, token)
			if err != nil {
				if err := m.errorWorker.RespError(c, err); err != nil {
					return err
				}
				return err
			}
			return c.JSON(http.StatusOK, response)
		}
		return next(c)
	}
}

func (m middlew) CSRFToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		/*token := c.Request().Header.Get("X-CSRF-Token")
		userID := c.Get("userId").(int64)

		err := csrf.CheckToken(userID, token)
		if err != nil {
			newToken, _ := csrf.GenerateToken(userID)
			if err := m.errorWorker.TokenError(c, newToken); err != nil {
				return err
			}
			return err
		}*/

		return next(c)
	}
}

func (m middlew) Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		return m.internalLogger.WriteLog(err, c)
	}
}

func (m middlew) CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("ID")
		boardID, err := strconv.ParseInt(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		err = m.boardService.CheckBoardPermission(userID, boardID, false)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}

func (m middlew) CheckBoardAdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("ID")
		boardID, err := strconv.ParseInt(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		err = m.boardService.CheckBoardPermission(userID, boardID, true)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}

func (m middlew) CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cid := c.Param("ID")
		cardID, err := strconv.ParseInt(cid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		boardID, err := m.boardService.CheckCardPermission(userID, cardID)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}

func (m middlew) CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid := c.Param("ID")
		taskID, err := strconv.ParseInt(tid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		boardID, err := m.boardService.CheckTaskPermission(userID, taskID)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}

func (m middlew) CheckCommentUpdateUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		comid := c.Param("ID")
		commentID, err := strconv.ParseInt(comid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		boardID, err := m.boardService.CheckCommentPermission(userID, commentID, false)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}

func (m middlew) CheckCommentDeleteUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		comid := c.Param("ID")
		commentID, err := strconv.ParseInt(comid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(int64)

		boardID, err := m.boardService.CheckCommentPermission(userID, commentID, true)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", boardID)

		return next(c)
	}
}
