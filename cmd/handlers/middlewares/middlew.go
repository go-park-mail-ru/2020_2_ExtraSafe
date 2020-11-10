package middlewares

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/csrf"
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
	CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckBoardAdminPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc
}

type middlew struct {
	sessionsService sessionsService
	errorWorker errorWorker
	authService authService
	authTransport authTransport
	boardStorage boardStorage
}

func NewMiddleware(sessionsService sessionsService, errorWorker errorWorker, authService authService,
	authTransport authTransport, boardStorage boardStorage) Middleware {
	return middlew{
		sessionsService: sessionsService,
		errorWorker: errorWorker,
		authService: authService,
		authTransport: authTransport,
		boardStorage: boardStorage,
	}
}

func (m middlew) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://95.163.213.142:80",
			"http://95.163.213.142",
			"http://127.0.0.1:3033",
			"http://tabutask.ru"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
}

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err != nil {
			return m.errorWorker.TransportError(c)
		}
		c.Set("userId", userId)
		return next(c)
	}
}

func (m middlew) AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err == nil {
			userInput := new(models.UserInput)
			userInput.ID = userId
			user, _ := m.authService.Auth(*userInput)
			response, err := m.authTransport.AuthWrite(user)
			if err != nil {
				return m.errorWorker.TransportError(c)
			}
			return c.JSON(http.StatusOK, response)
		}
		return next(c)
	}
}

func (m middlew) CSRFToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("X-CSRF-Token")
		fmt.Println(token)
		userID := c.Get("userId").(uint64)

		err := csrf.CheckToken(userID, token)
		if err != nil {
			newToken, _ := csrf.GenerateToken(userID)
			return m.errorWorker.TokenError(c, newToken)
		}

		return next(c)
	}
}

/*func (m middlew) RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		result := next(c)
		if result != nil {
			logger.Error("%s %s %d %s",
				c.Request().Method,
				c.Request().RequestURI,
				c.Response().Status,
				time.Since(start))
		} else {
			logger.Infof("%s %s %d %s",
				c.Request().Method,
				c.Request().RequestURI,
				c.Response().Status,
				time.Since(start))
		}
		return result
	}
}*/


func (m middlew) CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("ID")
		boardID, err := strconv.ParseUint(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckBoardPermission(userID, boardID, false)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", bid)

		return next(c)
	}
}

func (m middlew) CheckBoardAdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("ID")
		boardID, err := strconv.ParseUint(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckBoardPermission(userID, boardID, true)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", bid)

		return next(c)
	}
}

func (m middlew) CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cid := c.Param("ID")
		cardID, err := strconv.ParseUint(cid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckCardPermission(userID, cardID)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		//c.Set("boardID", cid)

		return next(c)
	}
}

func (m middlew) CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid := c.Param("ID")
		taskID, err := strconv.ParseUint(tid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckTaskPermission(userID, taskID)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		//c.Set("boardID", tid)

		return next(c)
	}
}