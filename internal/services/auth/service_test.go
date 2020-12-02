package auth

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/validation"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/mock"
	protoAuth "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestService_CheckCookie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := "ALMoijiomIUHNbgyuygfuyefgKAJmiejcuierhhiauwdh"

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockAuthClient(ctrlAuth)

	service := service{authService: mockAuthService}

	input := &protoAuth.UserSession{
		SessionID: SID,
		UserID:    -1,
	}

	mockAuthService.EXPECT().CheckCookie(context.Background(), input).Return(&protoProfile.UserID{ID: int64(1)}, nil)

	userID, err := service.CheckCookie(ctx)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(userID, int64(1)) {
		t.Errorf("results not match, want %v, have %v", int64(1), userID)
		return
	}
}

func TestService_Auth(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockAuthClient(ctrlAuth)

	service := service{authService: mockAuthService}
	request := models.UserInput{ID: 1}
	input := &protoProfile.UserID{ID: request.ID}

	boards := make([]*protoProfile.BoardOutsideShort, 0)
	boards = append(boards, &protoProfile.BoardOutsideShort{
		BoardID: 1,
		Name:    "name",
		Theme:   "dark",
		Star:    false,
	})

	internal := &protoProfile.UserBoardsOutside{
		Email:    "epridius@gmail",
		Username: "pkaterinaa",
		FullName: "lalala",
		Avatar:   "default",
		Boards:   boards,
	}

	responseBoards := make([]models.BoardOutsideShort, 0)
	responseBoards = append(responseBoards, models.BoardOutsideShort{
		BoardID: boards[0].BoardID,
		Name:    boards[0].Name,
		Theme:   boards[0].Theme,
		Star:    boards[0].Star,
	})

	response := models.UserBoardsOutside{
		Email:    internal.Email,
		Username: internal.Username,
		FullName: internal.FullName,
		Avatar:   internal.Avatar,
		Boards:   responseBoards,
	}

	mockAuthService.EXPECT().Auth(context.Background(), input).Return(internal, nil)
	output, err := service.Auth(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, response) {
		t.Errorf("results not match, want %v, have %v", response, output)
		return
	}

}

func TestService_Registration(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockAuthClient(ctrlAuth)

	validator := validation.NewService()
	service := NewService(mockAuthService, validator)

	request := models.UserInputReg{
		Email:    "pridiuskate@gmail.com",
		Username: "pkaterinaa",
		Password: "12212121",
	}

	input := &protoProfile.UserInputReg{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	session := &protoAuth.UserSession{
		SessionID: "ansfbaoufnguwqgddobwuyifq",
		UserID:    1,
	}

	expect := models.UserSession{
		SessionID: session.SessionID,
		UserID:    session.UserID,
	}

	mockAuthService.EXPECT().Registration(context.Background(), input).Return(session, nil)
	output, err := service.Registration(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_Login(t *testing.T) {
	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockAuthClient(ctrlAuth)

	validator := validation.NewService()
	service := NewService(mockAuthService, validator)

	request := models.UserInputLogin{
		Email:    "pridiuskate@gmail.com",
		Password: "12212121",
	}

	input := &protoProfile.UserInputLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	session := &protoAuth.UserSession{
		SessionID: "ansfbaoufnguwqgddobwuyifq",
		UserID:    1,
	}

	expect := models.UserSession{
		SessionID: session.SessionID,
		UserID:    session.UserID,
	}

	mockAuthService.EXPECT().Login(context.Background(), input).Return(session, nil)
	output, err := service.Login(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_Logout(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	cookie := new(http.Cookie)
	SID := "ALMoijiomIUHNbgyuygfuyefgKAJmiejcuierhhiauwdh"

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.Request().AddCookie(cookie)

	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockAuthClient(ctrlAuth)

	service := service{authService: mockAuthService}

	input := &protoAuth.UserSession{
		SessionID: SID,
		UserID:    -1,
	}

	mockAuthService.EXPECT().DeleteCookie(context.Background(), input).Return(&protoAuth.Nothing{Dummy: true}, nil)

	err := service.Logout(ctx)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
