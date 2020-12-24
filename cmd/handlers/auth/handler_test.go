package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	errorWorker2 "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Auth(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInput{ID: int64(1)}

	outside := models.UserBoardsOutside{
		Email:    "mam",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig.png",
		Boards:   nil,
	}

	mockAuthService.EXPECT().Auth(input).Return(outside, nil)

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := authHandler.Auth(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_AuthFail(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInput{ID: int64(1)}

	outside := models.UserBoardsOutside{
		Email:    "mam",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig.png",
		Boards:   nil,
	}

	mockAuthService.EXPECT().Auth(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := authHandler.Auth(c)

	if err == nil {
		t.Errorf("results not match %s", err)
		return
	}
}

func TestHandler_Login(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInputLogin{
		Email:    "mariya@mail.ru",
		Password: "1234",
	}

	outside := models.UserSession{
		SessionID: "4321",
		UserID:    1,
	}

	mockAuthService.EXPECT().Login(input).Return(outside, nil)

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	userJSON := `{"email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := authHandler.Login(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_LoginFail(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInputLogin{
		Email:    "mariya@mail.ru",
		Password: "1234",
	}

	outside := models.UserSession{
		SessionID: "4321",
		UserID:    1,
	}

	mockAuthService.EXPECT().Login(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	userJSON := `{"email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := authHandler.Login(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Registration(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInputReg{
		Email:    "mariya@mail.ru",
		Username: "kit",
		Password: "1234",
	}

	outside := models.UserSession{
		SessionID: "4321",
		UserID:    1,
	}

	mockAuthService.EXPECT().Registration(input).Return(outside, nil)

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	userJSON := `{"username":"kit","email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := authHandler.Registration(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_RegistrationFail(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	input := models.UserInputReg{
		Email:    "mariya@mail.ru",
		Username: "kit",
		Password: "1234",
	}

	outside := models.UserSession{
		SessionID: "4321",
		UserID:    1,
	}

	mockAuthService.EXPECT().Registration(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	userJSON := `{"username":"kit","email":"mariya@mail.ru","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := authHandler.Registration(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Logout(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	authTransport := auth.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	mockAuthService.EXPECT().Logout(c).Return(nil)

	authHandler := NewHandler(mockAuthService, authTransport, errorWorker)

	err := authHandler.Logout(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}