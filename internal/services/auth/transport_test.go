package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTransport_AuthRead(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", 1)

	expectedUserInput := models.UserInput{ID: 1}

	transp := &transport{}

	userInput, _ := transp.AuthRead(c)

	assert.Equal(t, http.StatusCreated, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}

}