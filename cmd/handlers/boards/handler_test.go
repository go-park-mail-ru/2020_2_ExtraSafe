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
		UserID:  1,
		BoardID: 4,
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Board(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardChangeInput{
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardInput{
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardAddMember(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardAddMember(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_BoardRemoveMember(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.BoardMemberInput{
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
	c.SetPath("/board/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.BoardRemoveMember(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Card(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		UserID:  1,
		CardID:  4,
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
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Card(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		UserID:  1,
		CardID:  4,
		BoardID: 4,
		Name:   "todo",
		Order:  1,
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
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		UserID:  1,
		CardID:  4,
		BoardID: 4,
		Name:   "todo",
		Order:  1,
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
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardInput{
		UserID:  1,
		CardID:  4,
		BoardID: 4,
		Name:   "todo",
		Order:  1,
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
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_CardOrder(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.CardsOrderInput{
		UserID: 1,
		Cards:  []models.CardOrder{},
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
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.CardOrder(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_Task(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		TaskID: int64(4),
		UserID: 1,
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.Task(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		TaskID: int64(4),
		UserID: 1,
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		TaskID: int64(4),
		UserID: 1,
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskInput{
		TaskID: int64(4),
		UserID: 1,
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskOrder(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TasksOrderInput{
		UserID: 1,
		Tasks:  nil,
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskOrder(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserAdd(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserAdd(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TaskUserRemove(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TaskAssignerInput{
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
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TaskUserRemove(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagCreate(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		UserID:  1,
		TaskID:  4,
		TagID:   4,
		BoardID: 1,
		Color:   "red",
		Name:    "fds",
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
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagCreate(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagChange(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		UserID:  1,
		TaskID:  4,
		TagID:   4,
		BoardID: 1,
		Color:   "red",
		Name:    "fds",
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
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagChange(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagDelete(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockProfileService := mock.NewMockServiceBoard(ctrlBoard)

	input := models.TagInput{
		UserID:  1,
		TaskID:  4,
		TagID:   4,
		BoardID: 1,
		Color:   "red",
		Name:    "fds",
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
	c.SetPath("/tag/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	err := boardHandler.TagDelete(c)

	if err != nil {
		t.Errorf("results not match")
		return
	}
}

func TestHandler_TagAdd(t *testing.T) {

}

func TestHandler_TagRemove(t *testing.T) {

}

func TestHandler_CommentCreate(t *testing.T) {

}

func TestHandler_CommentChange(t *testing.T) {

}

func TestHandler_CommentDelete(t *testing.T) {

}

func TestHandler_ChecklistCreate(t *testing.T) {

}

func TestHandler_ChecklistChange(t *testing.T) {

}

func TestHandler_ChecklistDelete(t *testing.T) {

}

func TestHandler_AttachmentCreate(t *testing.T) {

}

func TestHandler_AttachmentDelete(t *testing.T) {

}