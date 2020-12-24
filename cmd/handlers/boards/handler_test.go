package boardsHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	errorWorker2 "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Board(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	outside := models.BoardOutside{
		BoardID: 4,
		Admin:   models.UserOutsideShort{},
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
		Users:   nil,
		Cards:   nil,
		Tags:    nil,
	}

	mockProfileService.EXPECT().GetBoard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Board(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	outside := models.BoardOutside{
		BoardID: 4,
		Admin:   models.UserOutsideShort{},
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
		Users:   nil,
		Cards:   nil,
		Tags:    nil,
	}

	mockProfileService.EXPECT().GetBoard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Board(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		BoardName: "vs",
		Theme:     "hy",
		Star:      false,
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	mockProfileService.EXPECT().CreateBoard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"boardName":"vs","boardTheme":"hy","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		BoardName: "vs",
		Theme:     "hy",
		Star:      false,
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	mockProfileService.EXPECT().CreateBoard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"boardName":"vs","boardTheme":"hy","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		BoardName: "vs",
		Theme:     "hy",
		Star:      false,
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	mockProfileService.EXPECT().ChangeBoard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"boardName":"vs","boardTheme":"hy","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		BoardName: "vs",
		Theme:     "hy",
		Star:      false,
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	mockProfileService.EXPECT().ChangeBoard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"boardName":"vs","boardTheme":"hy","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	mockProfileService.EXPECT().DeleteBoard(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	mockProfileService.EXPECT().DeleteBoard(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardAddMember(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
		SessionID:  "13",
		UserID:     1,
		BoardID:    4,
		MemberName: "kit",
	}

	outside := models.UserOutsideShort{
		ID:       1,
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().AddMember(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"memberUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardAddMember(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardAddMemberFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
		SessionID:  "13",
		UserID:     1,
		BoardID:    4,
		MemberName: "kit",
	}

	outside := models.UserOutsideShort{
		ID:       1,
		Email:    "mum",
		Username: "kit",
		FullName: "tik",
		Avatar:   "pig",
	}

	mockProfileService.EXPECT().AddMember(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"memberUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardAddMember(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardRemoveMember(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
		SessionID:  "13",
		UserID:     1,
		BoardID:    4,
		MemberName: "kit",
	}

	mockProfileService.EXPECT().RemoveMember(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"memberUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardRemoveMember(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardRemoveMemberFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
		SessionID:  "13",
		UserID:     1,
		BoardID:    4,
		MemberName: "kit",
	}

	mockProfileService.EXPECT().RemoveMember(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"memberUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardRemoveMember(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Card(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   1,
	}

	outside := models.CardOutside{
		CardID: 4,
		Name:   "wet",
		Order:  1,
		Tasks:  nil,
	}

	mockProfileService.EXPECT().GetCard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Card(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   1,
	}

	outside := models.CardOutside{
		CardID: 4,
		Name:   "wet",
		Order:  1,
		Tasks:  nil,
	}

	mockProfileService.EXPECT().GetCard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Card(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	outside := models.CardOutsideShort{
		CardID: 4,
		Name:   "todo",
	}

	mockProfileService.EXPECT().CreateCard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	outside := models.CardOutsideShort{
		CardID: 4,
		Name:   "todo",
	}

	mockProfileService.EXPECT().CreateCard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	outside := models.CardOutsideShort{
		CardID: 4,
		Name:   "todo",
	}

	mockProfileService.EXPECT().ChangeCard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	outside := models.CardOutsideShort{
		CardID: 4,
		Name:   "todo",
	}

	mockProfileService.EXPECT().ChangeCard(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	mockProfileService.EXPECT().DeleteCard(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		SessionID: "13",
		UserID:    1,
		CardID:    4,
		BoardID:   4,
		Name:      "todo",
		Order:     1,
	}

	mockProfileService.EXPECT().DeleteCard(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cardID":4,"cardName":"todo","cardOrder":1,"boardID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardOrder(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardsOrderInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		Cards:     []models.CardOrder{},
	}

	mockProfileService.EXPECT().CardOrderChange(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cards":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardOrder(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardOrderFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardsOrderInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
		Cards:     []models.CardOrder{},
	}

	mockProfileService.EXPECT().CardOrderChange(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"cards":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardOrder(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Task(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		BoardID:   4,
		TaskID:    int64(4),
		UserID:    1,
	}

	outside := models.TaskOutside{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
		Order:       1,
		Tags:        nil,
		Users:       nil,
		Checklists:  nil,
		Comments:    nil,
		Attachments: nil,
	}

	mockProfileService.EXPECT().GetTask(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Task(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		BoardID:   4,
		TaskID:    int64(4),
		UserID:    1,
	}

	outside := models.TaskOutside{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
		Order:       1,
		Tags:        nil,
		Users:       nil,
		Checklists:  nil,
		Comments:    nil,
		Attachments: nil,
	}

	mockProfileService.EXPECT().GetTask(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Task(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		TaskID:    int64(4),
		BoardID:   4,
		UserID:    1,
	}

	outside := models.TaskOutsideSuperShort{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
	}

	mockProfileService.EXPECT().CreateTask(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		TaskID:    int64(4),
		BoardID:   4,
		UserID:    1,
	}

	outside := models.TaskOutsideSuperShort{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
	}

	mockProfileService.EXPECT().CreateTask(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		BoardID:   4,
		TaskID:    int64(4),
		UserID:    1,
	}

	outside := models.TaskOutsideSuperShort{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
	}

	mockProfileService.EXPECT().ChangeTask(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		BoardID:   4,
		TaskID:    int64(4),
		UserID:    1,
	}

	outside := models.TaskOutsideSuperShort{
		TaskID:      4,
		Name:        "tt",
		Description: "tt tt",
	}

	mockProfileService.EXPECT().ChangeTask(input).Return(outside, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		SessionID: "13",
		BoardID:   4,
		TaskID:    int64(4),
		UserID:    1,
	}

	mockProfileService.EXPECT().DeleteTask(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		TaskID:    int64(4),
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
	}

	mockProfileService.EXPECT().DeleteTask(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskOrder(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TasksOrderInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		Tasks:     nil,
	}

	mockProfileService.EXPECT().TasksOrderChange(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tasks":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskOrder(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskOrderFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TasksOrderInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		Tasks:     nil,
	}

	mockProfileService.EXPECT().TasksOrderChange(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tasks":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskOrder(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserAdd(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AssignerName: "kit",
	}

	output := models.UserOutsideShort{
		ID:       1,
		Email:    "mz",
		Username: "kit",
		FullName: "hv",
		Avatar:   "gf",
	}

	mockProfileService.EXPECT().AssignUser(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4,"assignerUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserAdd(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserAddFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AssignerName: "kit",
	}

	output := models.UserOutsideShort{
		ID:       1,
		Email:    "mz",
		Username: "kit",
		FullName: "hv",
		Avatar:   "gf",
	}

	mockProfileService.EXPECT().AssignUser(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4,"assignerUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserAdd(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserRemove(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AssignerName: "kit",
	}

	mockProfileService.EXPECT().DismissUser(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4,"assignerUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserRemove(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserRemoveFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AssignerName: "kit",
	}

	mockProfileService.EXPECT().DismissUser(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"taskID":4,"assignerUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserRemove(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	output := models.TagOutside{
		TagID: 4,
		Color: "red",
		Name:  "fds",
	}

	mockProfileService.EXPECT().CreateTag(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	output := models.TagOutside{
		TagID: 4,
		Color: "red",
		Name:  "fds",
	}

	mockProfileService.EXPECT().CreateTag(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	output := models.TagOutside{
		TagID: 4,
		Color: "red",
		Name:  "fds",
	}

	mockProfileService.EXPECT().ChangeTag(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	output := models.TagOutside{
		TagID: 4,
		Color: "red",
		Name:  "fds",
	}

	mockProfileService.EXPECT().ChangeTag(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	mockProfileService.EXPECT().DeleteTag(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		SessionID: "13",
		UserID:    1,
		TaskID:    4,
		TagID:     4,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	mockProfileService.EXPECT().DeleteTag(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagAdd(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskTagInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		TaskID:    4,
		TagID:     4,
	}

	mockProfileService.EXPECT().AddTag(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagAdd(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagAddFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskTagInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		TaskID:    4,
		TagID:     4,
	}

	mockProfileService.EXPECT().AddTag(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagAdd(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagRemove(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskTagInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		TaskID:    4,
		TagID:     4,
	}

	mockProfileService.EXPECT().RemoveTag(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagRemove(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagRemoveFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskTagInput{
		SessionID: "13",
		BoardID:   4,
		UserID:    1,
		TaskID:    4,
		TagID:     4,
	}

	mockProfileService.EXPECT().RemoveTag(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"tagID":4,"taskID":4}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagRemove(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	output := models.CommentOutside{
		CommentID: 5,
		Message:   "gggg",
		Order:     1,
		User:      models.UserOutsideShort{},
	}

	mockProfileService.EXPECT().CreateComment(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	output := models.CommentOutside{
		CommentID: 5,
		Message:   "gggg",
		Order:     1,
		User:      models.UserOutsideShort{},
	}

	mockProfileService.EXPECT().CreateComment(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	output := models.CommentOutside{
		CommentID: 5,
		Message:   "gggg",
		Order:     1,
		User:      models.UserOutsideShort{},
	}

	mockProfileService.EXPECT().ChangeComment(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	output := models.CommentOutside{
		CommentID: 5,
		Message:   "gggg",
		Order:     1,
		User:      models.UserOutsideShort{},
	}

	mockProfileService.EXPECT().ChangeComment(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	mockProfileService.EXPECT().DeleteComment(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CommentDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CommentInput{
		SessionID: "13",
		BoardID:   4,
		CommentID: 5,
		TaskID:    4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	mockProfileService.EXPECT().DeleteComment(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/comment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CommentDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	output := models.ChecklistOutside{
		ChecklistID: 5,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().CreateChecklist(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistCreateFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	output := models.ChecklistOutside{
		ChecklistID: 5,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().CreateChecklist(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistCreate(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	output := models.ChecklistOutside{
		ChecklistID: 5,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().ChangeChecklist(input).Return(output, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistChangeFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	output := models.ChecklistOutside{
		ChecklistID: 5,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().ChangeChecklist(input).Return(output, models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistChange(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().DeleteChecklist(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_ChecklistDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.ChecklistInput{
		SessionID:   "13",
		BoardID:     4,
		UserID:      1,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	mockProfileService.EXPECT().DeleteChecklist(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/checklist/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.ChecklistDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_AttachmentCreate(t *testing.T) {

}

func TestHandler_AttachmentCreateFail(t *testing.T) {

}

func TestHandler_AttachmentDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.AttachmentInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AttachmentID: 5,
		Filename:     "ffff",
	}

	mockProfileService.EXPECT().DeleteAttachment(input).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"attachmentID":5,"taskID":4,"attachmentFileName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/attachment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.AttachmentDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_AttachmentDeleteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.AttachmentInput{
		SessionID:    "13",
		BoardID:      4,
		UserID:       1,
		TaskID:       4,
		AttachmentID: 5,
		Filename:     "ffff",
	}

	mockProfileService.EXPECT().DeleteAttachment(input).Return(models.ServeError{Codes: []string{models.ServerError}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	userJSON := `{"attachmentID":5,"taskID":4,"attachmentFileName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/attachment/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.AttachmentDelete(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Notification(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.UserInput{
		SessionID: "13",
		ID:        1,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockProfileService.EXPECT().WebSocketNotification(input, c).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	c.Set("userId", int64(1))
	c.Set("sessionID", "13")

	err := boardHandler.Notification(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_SharedURL(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	outside := "12345"

	mockProfileService.EXPECT().GetSharedURL(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.SharedURL(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_SharedURLFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		SessionID: "13",
		UserID:    1,
		BoardID:   4,
	}

	outside := "12345"

	mockProfileService.EXPECT().GetSharedURL(input).Return(outside, models.ServeError{Codes: []string{"500"}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.SharedURL(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardWS(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
		UserID:    1,
		SessionID: "13",
		BoardID:   4,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	mockProfileService.EXPECT().WebSocketBoard(input, c).Return(nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	err := boardHandler.BoardWS(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardInvite(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInviteInput{
		UserID:  1,
		BoardID: 4,
		UrlHash: "12345",
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetParamNames("ID", "url")
	c.SetParamValues("4", "12345")

	mockProfileService.EXPECT().InviteUserToBoard(input).Return(outside, nil)

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	err := boardHandler.BoardInvite(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardInviteFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInviteInput{
		UserID:  1,
		BoardID: 4,
		UrlHash: "12345",
	}

	outside := models.BoardOutsideShort{
		BoardID: 4,
		Name:    "vs",
		Theme:   "hy",
		Star:    false,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetParamNames("ID", "url")
	c.SetParamValues("4", "12345")

	mockProfileService.EXPECT().InviteUserToBoard(input).Return(outside, models.ServeError{Codes: []string{"500"}})

	boardTransport := boards.NewTransport()

	errorWorker := errorWorker2.NewErrorWorker()

	boardHandler := NewHandler(mockProfileService, boardTransport, errorWorker)

	err := boardHandler.BoardInvite(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}
