package profileHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	errorWorker2 "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Profile(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInput{ID: int64(1)}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().Profile(input).Return(outside, nil)

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.Profile(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ProfileFail(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInput{ID: int64(1)}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().Profile(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.Profile(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ProfileChange(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInputProfile{
		ID:       1,
		Email:    "",
		Username: "",
		FullName: "",
		Avatar:   nil,
	}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().ProfileChange(input).Return(outside, nil)

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	userJSON := `{"username":"kit","email":"mum","fullName":"tik"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.ProfileChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ProfileChangeFail(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInputProfile{
		ID:       1,
		Email:    "",
		Username: "",
		FullName: "",
		Avatar:   nil,
	}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().ProfileChange(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	userJSON := `{"username":"kit","email":"mum","fullName":"tik"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.ProfileChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_PasswordChange(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInputPassword{
		ID:          1,
		OldPassword: "12345",
		Password:    "1234",
	}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().PasswordChange(input).Return(outside, nil)

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	userJSON := `{"oldpassword":"12345","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.PasswordChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_PasswordChangeFail(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInputPassword{
		ID:          1,
		OldPassword: "12345",
		Password:    "1234",
	}

	outside := models.UserOutside{
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().PasswordChange(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	userJSON := `{"oldpassword":"12345","password":"1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.PasswordChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Boards(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInput{ID: int64(1)}

	mockProfileService.EXPECT().Boards(input).Return(nil, nil)

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.Boards(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardsFail(t *testing.T) {
	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := mock.NewMockServiceProfile(ctrlProfile)

	input := models.UserInput{ID: int64(1)}

	mockProfileService.EXPECT().Boards(input).Return(nil, models.ServeError{Codes: []string{models.ServerError}})

	profileTransport := profile.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	profileHandler := NewHandler(mockProfileService, profileTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	err := profileHandler.Boards(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}