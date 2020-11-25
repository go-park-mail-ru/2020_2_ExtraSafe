package sessions

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/sessions/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestService_CheckCookie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	mockSessionStorage.EXPECT().CheckUserSession(SID).Return(uint64(1), nil)


	userID, _ := service.CheckCookie(ctx)
	assert.Equal(t, uint64(1), userID)
}

func TestService_CheckCookieInternalFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	mockSessionStorage.EXPECT().CheckUserSession(SID).Return(uint64(1), errors.New(""))

	_, err := service.CheckCookie(ctx)
	assert.Error(t, err)
}

func TestService_CheckCookieFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	_, err := service.CheckCookie(ctx)
	assert.Error(t, err)
}

func TestService_SetCookie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	mockSessionStorage.EXPECT().CreateUserSession(uint64(1), gomock.Any()).Return(nil)
	err := service.SetCookie(ctx, uint64(1))
	assert.Equal(t, nil, err)
}

func TestService_SetCookieFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)
	resultErr := errors.New("internal error while creating")

	mockSessionStorage.EXPECT().CreateUserSession(uint64(1), gomock.Any()).Return(resultErr)
	err := service.SetCookie(ctx, uint64(1))
	assert.Error(t, err)
}

func TestService_DeleteCookie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	mockSessionStorage.EXPECT().DeleteUserSession(SID).Return(nil)

	err := service.DeleteCookie(ctx)
	assert.Equal(t, nil, err)
}

func TestService_DeleteCookieBadName(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	err := service.DeleteCookie(ctx)
	assert.Error(t, err)
}

func TestService_DeleteCookieFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStorage := mock.NewMockSessionStorage(ctrl)

	service := NewService(mockSessionStorage)

	mockSessionStorage.EXPECT().DeleteUserSession(SID).Return(errors.New("no such cookie"))

	err := service.DeleteCookie(ctx)
	assert.Error(t, err)
}
