package logger

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/rs/zerolog"
	"time"
)

type Logger interface {
	WriteLog(err error, c echo.Context) error
}

type logger struct {
	zeroLogger *zerolog.Logger
}

func NewLogger(zeroLogger *zerolog.Logger) Logger {
	return &logger{
		zeroLogger: zeroLogger,
	}
}

func (l logger) WriteLog(err error, c echo.Context) error {
	if err == nil {
		infoLog := l.zeroLogger.Info()
		infoLog.
			Time("Request time",time.Now()).
			Str("URL", c.Request().RequestURI).
			Int("Status", c.Response().Status)
		infoLog.Send()
		return err
	}

	servError, ok := err.(models.ServeError)
	if ok != true {
		errLog := l.zeroLogger.Error()
		errLog.
			Time("Request time",time.Now()).
			Str("URL", c.Request().RequestURI).
			Int("Status", c.Response().Status).
			Str("Error ", err.Error())
		errLog.Send()
		return err
	}

	for i, code := range servError.Codes {
		errLog := l.zeroLogger.Error()
		if len(err.(models.ServeError).Descriptions) == 0 {
			errLog.
				Time("Request time",time.Now()).
				Str("URL", c.Request().RequestURI).
				Int("Status", c.Response().Status).
				Str("Error code", code).
				Str("Error ", err.Error()).
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
