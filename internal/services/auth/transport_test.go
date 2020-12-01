package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestTransport_AuthRead(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	expectedUserInput := models.UserInput{ID: 1}

	transp := &transport{}

	userInput, _ := transp.AuthRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_RegRead(t *testing.T) {
	userJSON := `{"username":"kit","email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reg", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedUserInput := models.UserInputReg{
		Email:    "mariya@mail.ru",
		Username: "kit",
		Password: "1234",
	}

	transp := &transport{}

	userInput, _ := transp.RegRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_RegReadFail(t *testing.T) {
	userJSON := `{"username":1,"gmail":"mariya@mail.ru","word":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reg", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	transp := &transport{}

	_, err := transp.RegRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_LoginRead(t *testing.T) {
	userJSON := `{"email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reg", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedUserInput := models.UserInputLogin{
		Email:    "mariya@mail.ru",
		Password: "1234",
	}

	transp := &transport{}

	userInput, _ := transp.LoginRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_LoginReadFail(t *testing.T) {
	userJSON := `{"email":4,"password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reg", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	transp := &transport{}

	_, err := transp.LoginRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_AuthWrite(t *testing.T) {
	input := models.UserBoardsOutside{
		Email:    "mariya@mail.ru",
		Username: "kit",
		FullName: "kit tik",
		Avatar:   "pig.png",
		Boards:   nil,
	}

	expectedResponse:= models.ResponseUserAuth{
		Status:   200,
		Token:    "12345",
		Email:    "mariya@mail.ru",
		Username: "kit",
		FullName: "kit tik",
		Avatar:   "pig.png",
		Boards:   nil,
	}

	transp := &transport{}

	response, _ := transp.AuthWrite(input, "12345")

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_LoginWrite(t *testing.T) {
	expectedResponse:= models.ResponseToken{
		Status:   200,
		Token:    "12345",
	}

	transp := &transport{}

	response, _ := transp.LoginWrite("12345")

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_RegWrite(t *testing.T) {
	expectedResponse:= models.ResponseStatus{
		Status:   200,
	}

	transp := NewTransport()

	response, _ := transp.RegWrite()

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}