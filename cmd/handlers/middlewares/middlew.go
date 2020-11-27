package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/csrf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"time"
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
}

type middlew struct {
	errorWorker errorWorker
	authService auth.Service
	authTransport authTransport
	boardService boards.Service
	logger *zerolog.Logger
}

func NewMiddleware(errorWorker errorWorker, authService auth.Service, authTransport authTransport,
	boardService boards.Service, logger *zerolog.Logger) Middleware {
	return middlew{
		errorWorker: errorWorker,
		authService: authService,
		authTransport: authTransport,
		boardService: boardService,
		logger: logger,
	}
}

func (m middlew) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://95.163.213.142:80",
			"http://95.163.213.142",
			"http://127.0.0.1:3033",
			"http://127.0.0.1",
			"http://tabutask.ru"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-CSRF-Token"},
		AllowCredentials: true,
	})
}

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.authService.CheckCookie(c)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}
		c.Set("userId", userId)
		return next(c)
	}
}

func (m middlew) AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.authService.CheckCookie(c)
		if err == nil {
			userInput := new(models.UserInput)
			userInput.ID = userId
			user, _ := m.authService.Auth(*userInput)
			token, _ := csrf.GenerateToken(userId)
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

		if err == nil {
			infoLog := m.logger.Info()
			infoLog.
				Time("Request time",time.Now()).
				Str("URL", c.Request().RequestURI).
				Int("Status", c.Response().Status)
			infoLog.Send()
			return err
		}

		for i, code := range err.(models.ServeError).Codes {
			errLog := m.logger.Error()
			//TODO - исправить, если ошибка не из микросервиса, нет дескрипшена
			if code == models.ServerError {
				errLog.
					Time("Request time",time.Now()).
					Str("URL", c.Request().RequestURI).
					Int("Status", c.Response().Status).
					Str("Error code", code).
					Str("Error ", err.(models.ServeError).Descriptions[i]).
					Str("In function", err.(models.ServeError).MethodName)
				errLog.Send()
			} else {
				errLog.
					Time("Request time", time.Now()).
					Str("URL", c.Request().RequestURI).
					Int("Status", c.Response().Status).
					Str("Error code", code).
					Str("Error ", err.(models.ServeError).Descriptions[i]).
					Str("In function", err.(models.ServeError).MethodName)
				errLog.Send()
			}
		}
		return err
	}
}

func (m middlew) CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("ID")
		boardID, err := strconv.ParseInt(bid, 10, 64)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		userID := c.Get("userId").(int64)

		err = m.boardService.CheckBoardPermission(userID, boardID, false)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		c.Set("boardID", bid)

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

		c.Set("boardID", bid)

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

		err = m.boardService.CheckCardPermission(userID, cardID)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

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

		err = m.boardService.CheckTaskPermission(userID, taskID)
		if err != nil {
			if err := m.errorWorker.RespError(c, err); err != nil {
				return err
			}
			return err
		}

		return next(c)
	}
}