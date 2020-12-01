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

}

func TestHandler_BoardChange(t *testing.T) {

}

func TestHandler_BoardDelete(t *testing.T) {

}

func TestHandler_BoardAddMember(t *testing.T) {

}

func TestHandler_BoardRemoveMember(t *testing.T) {

}

func TestHandler_Card(t *testing.T) {

}

func TestHandler_CardCreate(t *testing.T) {

}

func TestHandler_CardChange(t *testing.T) {

}

func TestHandler_CardDelete(t *testing.T) {

}

func TestHandler_CardOrder(t *testing.T) {

}

func TestHandler_Task(t *testing.T) {

}

func TestHandler_TaskCreate(t *testing.T) {

}

func TestHandler_TaskChange(t *testing.T) {

}

func TestHandler_TaskDelete(t *testing.T) {

}

func TestHandler_TaskOrder(t *testing.T) {

}

func TestHandler_TaskUserAdd(t *testing.T) {

}

func TestHandler_TaskUserRemove(t *testing.T) {

}

func TestHandler_TagCreate(t *testing.T) {

}

func TestHandler_TagChange(t *testing.T) {

}

func TestHandler_TagDelete(t *testing.T) {

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