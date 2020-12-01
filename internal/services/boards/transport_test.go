package boards

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

func TestTransport_BoardRead(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardInput{
		UserID:  1,
		BoardID: 3,
	}

	transp := &transport{}

	userInput, _ := transp.BoardRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_BoardReadFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("j")

	transp := NewTransport()

	_, err := transp.BoardRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_BoardChangeRead(t *testing.T) {
	userJSON := `{"boardName":"tx","boardTheme":"dark","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   3,
		BoardName: "tx",
		Theme:     "dark",
		Star:      false,
	}

	transp := &transport{}

	userInput, _ := transp.BoardChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_BoardChangeReadFail(t *testing.T) {
	userJSON := `{"boardName":1,"boardTheme":"dark","boardStar":false}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	transp := &transport{}

	_, err := transp.BoardChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_BoardMemberRead(t *testing.T) {
	userJSON := `{"memberUsername":"kit"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardMemberInput{
		UserID:    1,
		BoardID:   3,
		MemberName: "kit",
	}

	transp := &transport{}

	userInput, _ := transp.BoardMemberRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_BoardMemberReadFail(t *testing.T) {
	userJSON := `{"memberUsername":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	transp := &transport{}

	_, err := transp.BoardMemberRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_BoardWrite(t *testing.T) {
	input := models.BoardOutside{
		BoardID: 3,
		Admin:   models.UserOutsideShort{
			Email:    "mariya@mail.ru",
			Username: "kit",
			FullName: "kit tik",
			Avatar:   "pig.png",
		},
		Name:    "kt",
		Theme:   "dark",
		Star:    false,
		Users:   nil,
		Cards:   nil,
		Tags:    nil,
	}

	expectedResponse:= models.ResponseBoard{
		Status:  200,
		BoardID: 3,
		Admin:   models.UserOutsideShort{
			Email:    "mariya@mail.ru",
			Username: "kit",
			FullName: "kit tik",
			Avatar:   "pig.png",
		},
		Name:    "kt",
		Theme:   "dark",
		Star:    false,
		Users:   nil,
		Cards:   nil,
		Tags:    nil,
	}

	transp := &transport{}

	response, _ := transp.BoardWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_BoardShortWrite(t *testing.T) {
	input := models.BoardOutsideShort{
		BoardID: 3,
		Name:    "kt",
		Theme:   "dark",
		Star:    false,

	}

	expectedResponse:= models.ResponseBoardShort{
		Status:  200,
		BoardID: 3,
		Name:    "kt",
		Theme:   "dark",
		Star:    false,
	}

	transp := &transport{}

	response, _ := transp.BoardShortWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}



/*
CardChangeRead(c echo.Context) (request models.CardInput, err error)
CardWrite(card models.CardOutside) (response models.ResponseCard, err error)
CardShortWrite(card models.CardOutsideShort) (response models.ResponseCardShort, err error)
CardOrderRead(c echo.Context) (tasksOrder models.CardsOrderInput, err error)

TaskChangeRead(c echo.Context) (request models.TaskInput, err error)
TaskWrite(task models.TaskOutside) (response models.ResponseTask, err error)
TaskSuperShortWrite(task models.TaskOutsideSuperShort) (response models.ResponseTaskSuperShort, err error)
TasksOrderRead(c echo.Context) (tasksOrder models.TasksOrderInput, err error)
TasksUserRead(c echo.Context) (request models.TaskAssignerInput, err error)

TagChangeRead(c echo.Context) (request models.TagInput, err error)
TagTaskRead(c echo.Context) (request models.TaskTagInput, err error)
TagWrite(tag models.TagOutside) (response models.ResponseTag, err error)

CommentChangeRead(c echo.Context) (request models.CommentInput, err error)
CommentWrite(comment models.CommentOutside) (response models.ResponseComment, err error)

ChecklistChangeRead(c echo.Context) (request models.ChecklistInput, err error)
ChecklistWrite(checklist models.ChecklistOutside) (response models.ResponseChecklist, err error)

AttachmentAddRead(c echo.Context) (request models.AttachmentInput, err error)
AttachmentDeleteRead(c echo.Context) (request models.AttachmentInput, err error)
AttachmentWrite(attachment models.AttachmentOutside) (response models.ResponseAttachment, err error)

UserShortWrite(user models.UserOutsideShort) (response models.ResponseUser, err error)
*/

func TestTransport_CardChangeRead(t *testing.T) {
	userJSON := `{"cardName":"todo","cardOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	expectedUserInput := models.CardInput{
		UserID:  1,
		CardID:  4,
		Name:    "todo",
		Order:   1,
	}

	transp := &transport{}

	userInput, _ := transp.CardChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_CardChangeReadFail(t *testing.T) {
	userJSON := `{"cardName":1,"cardOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	transp := &transport{}

	_, err := transp.CardChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_CardWrite(t *testing.T) {

}

func TestTransport_CardShortWrite(t *testing.T) {

}

func TestTransport_CardOrderRead(t *testing.T) {
	userJSON := `{"cardName":"todo","cardOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	expectedUserInput := models.CardInput{
		UserID:  1,
		CardID:  4,
		Name:    "todo",
		Order:   1,
	}

	transp := &transport{}

	userInput, _ := transp.CardChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_CardOrderReadFail(t *testing.T) {

}