package errorWorker

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"net/http"
)

type ErrorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}

type errorWorker struct {
}

func NewErrorWorker() ErrorWorker {
	return &errorWorker{}
}

type ResponseError struct {
	Status        int      `json:"status"`
	Codes         []string `json:"codes"`
}

func (e errorWorker) RespError(c echo.Context, serveError error) (err error) {
	responseErr := new(ResponseError)
	responseErr.Codes = serveError.(models.ServeError).Codes
	responseErr.Status = 500
	return c.JSON(http.StatusBadRequest, responseErr)
}

func (e errorWorker) TransportError(c echo.Context) (err error) {
	responseErr := new(ResponseError)
	responseErr.Codes = nil
	responseErr.Status = 500
	return c.JSON(http.StatusBadRequest, responseErr)
}