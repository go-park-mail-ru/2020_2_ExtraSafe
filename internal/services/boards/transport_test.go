package boards

import (
	"bytes"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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
	c.Set("sessionID", "13")
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardInput{
		UserID:    1,
		BoardID:   3,
		SessionID: "13",
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
	c.Set("sessionID", "13")
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   3,
		SessionID: "13",
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
	c.Set("sessionID", "13")
	c.SetPath("/users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("3")

	expectedUserInput := models.BoardMemberInput{
		UserID:     1,
		BoardID:    3,
		SessionID:  "13",
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
		Admin: models.UserOutsideShort{
			Email:    "mariya@mail.ru",
			Username: "kit",
			FullName: "kit tik",
			Avatar:   "pig.png",
		},
		Name:  "kt",
		Theme: "dark",
		Star:  false,
		Users: nil,
		Cards: nil,
		Tags:  nil,
	}

	expectedResponse := models.ResponseBoard{
		Status:  200,
		BoardID: 3,
		Admin: models.UserOutsideShort{
			Email:    "mariya@mail.ru",
			Username: "kit",
			FullName: "kit tik",
			Avatar:   "pig.png",
		},
		Name:  "kt",
		Theme: "dark",
		Star:  false,
		Users: nil,
		Cards: nil,
		Tags:  nil,
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

	expectedResponse := models.ResponseBoardShort{
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

func TestTransport_CardChangeRead(t *testing.T) {
	userJSON := `{"cardName":"todo","cardOrder":1}`

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

	expectedUserInput := models.CardInput{
		UserID:    1,
		BoardID:   4,
		SessionID: "13",
		CardID:    4,
		Name:      "todo",
		Order:     1,
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

func TestTransport_CardChangeReadFail2(t *testing.T) {
	userJSON := `{"cardName":"dddd","cardOrder":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("a")

	transp := &transport{}

	_, err := transp.CardChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_CardWrite(t *testing.T) {
	input := models.CardOutside{
		CardID: 4,
		Name:   "todo",
		Order:  1,
		Tasks:  nil,
	}

	expectedResponse := models.ResponseCard{
		Status: 200,
		CardID: 4,
		Name:   "todo",
		Order:  1,
		Tasks:  nil,
	}

	transp := &transport{}

	response, _ := transp.CardWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_CardShortWrite(t *testing.T) {
	input := models.CardOutsideShort{
		CardID: 4,
		Name:   "todo",
	}

	expectedResponse := models.ResponseCardShort{
		Status: 200,
		CardID: 4,
		Name:   "todo",
	}

	transp := &transport{}

	response, _ := transp.CardShortWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_CardOrderRead(t *testing.T) {
	userJSON := `{"cards":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))

	expectedUserInput := models.CardsOrderInput{
		UserID:    1,
		BoardID:   4,
		SessionID: "13",
		Cards:     []models.CardOrder{},
	}

	transp := &transport{}

	userInput, _ := transp.CardOrderRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_CardOrderReadFail(t *testing.T) {
	userJSON := `{"cards":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	transp := &transport{}

	_, err := transp.CardOrderRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TaskChangeRead(t *testing.T) {
	userJSON := `{"taskID":0,"cardID":4,"taskName":"back","taskDescription":"write code","taskOrder":1}`

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

	expectedUserInput := models.TaskInput{
		UserID:      1,
		BoardID:     4,
		SessionID:   "13",
		TaskID:      4,
		CardID:      4,
		Name:        "back",
		Description: "write code",
		Order:       1,
	}

	transp := &transport{}

	userInput, _ := transp.TaskChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_TaskChangeReadFail(t *testing.T) {
	userJSON := `{"taskID":0,"cardID":4,"taskName":1,"taskDescription":"write code","taskOrder":1}`

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

	_, err := transp.TaskChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TaskChangeReadFail2(t *testing.T) {
	userJSON := `{"taskID":0,"cardID":4,"taskName":"ddd","taskDescription":"write code","taskOrder":1}`

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
	c.SetParamValues("a")

	transp := &transport{}

	_, err := transp.TaskChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TaskWrite(t *testing.T) {
	input := models.TaskOutside{
		TaskID:      5,
		Name:        "back",
		Description: "rrr rrr",
		Order:       1,
		Tags:        nil,
		Users:       nil,
		Checklists:  nil,
		Comments:    nil,
		Attachments: nil,
	}

	expectedResponse := models.ResponseTask{
		Status:      200,
		TaskID:      5,
		Name:        "back",
		Description: "rrr rrr",
		Order:       1,
		Tags:        nil,
		Users:       nil,
		Checklists:  nil,
		Comments:    nil,
		Attachments: nil,
	}

	transp := &transport{}

	response, _ := transp.TaskWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_TaskSuperShortWrite(t *testing.T) {
	input := models.TaskOutsideSuperShort{
		TaskID:      5,
		Name:        "back",
		Description: "rrr rrr",
	}

	expectedResponse := models.ResponseTaskSuperShort{
		Status:      200,
		TaskID:      5,
		Name:        "back",
		Description: "rrr rrr",
	}

	transp := &transport{}

	response, _ := transp.TaskSuperShortWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_TasksOrderRead(t *testing.T) {
	userJSON := `{"cards":[]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))

	expectedUserInput := models.TasksOrderInput{
		BoardID:   4,
		SessionID: "13",
		UserID:    1,
		Tasks:     []models.TasksOrder{},
	}

	transp := &transport{}

	userInput, _ := transp.TasksOrderRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_TasksOrderReadFail(t *testing.T) {
	userJSON := `{"cards": 1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	transp := &transport{}

	_, err := transp.TasksOrderRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TasksUserRead(t *testing.T) {
	userJSON := `{"assignerUsername":"kit","taskID":5}`

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
	c.SetParamValues("5")

	expectedUserInput := models.TaskAssignerInput{
		BoardID:      4,
		SessionID:    "13",
		UserID:       1,
		TaskID:       5,
		AssignerName: "kit",
	}

	transp := &transport{}

	userInput, _ := transp.TasksUserRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_TasksUserReadFail(t *testing.T) {
	userJSON := `{"assignerUsername":1,"taskID":5}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/task/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("5")

	transp := &transport{}

	_, err := transp.TasksUserRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TagChangeRead(t *testing.T) {
	userJSON := `{"tagID":2,"taskID":4,"boardID":1,"tagColor":"red","tagName":"fds"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")

	expectedUserInput := models.TagInput{
		UserID:    1,
		SessionID: "13",
		TaskID:    4,
		TagID:     2,
		BoardID:   1,
		Color:     "red",
		Name:      "fds",
	}

	transp := &transport{}

	userInput, _ := transp.TagChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_TagChangeReadFail(t *testing.T) {
	userJSON := `{"tagID":2,"taskID":4,"boardID":1,"tagColor":1,"tagName":"fds"}`

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

	_, err := transp.TagChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TagTaskRead(t *testing.T) {
	userJSON := `{"tagID":3,"taskID":2}`

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

	expectedUserInput := models.TaskTagInput{
		UserID:    1,
		SessionID: "13",
		BoardID:   4,
		TaskID:    2,
		TagID:     3,
	}

	transp := &transport{}

	userInput, _ := transp.TagTaskRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_TagTaskReadFail(t *testing.T) {
	userJSON := `{"tagID":3,"taskID":"uuu"}`

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

	_, err := transp.TagTaskRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_TagWrite(t *testing.T) {
	input := models.TagOutside{
		TagID: 6,
		Color: "red",
		Name:  "fsd",
	}

	expectedResponse := models.ResponseTag{
		Status:  200,
		TagID:   6,
		Color:   "red",
		TagName: "fsd",
	}

	transp := &transport{}

	response, _ := transp.TagWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_CommentChangeRead(t *testing.T) {
	userJSON := `{"commentID":5,"taskID":4,"commentMessage":"gggg","commentOrder":1}`

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

	expectedUserInput := models.CommentInput{
		CommentID: 5,
		TaskID:    4,
		SessionID: "13",
		BoardID:   4,
		Message:   "gggg",
		Order:     1,
		UserID:    1,
	}

	transp := &transport{}

	userInput, _ := transp.CommentChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_CommentChangeReadFail(t *testing.T) {
	userJSON := `{"commentID":5,"taskID":4,"commentMessage":1,"commentOrder":1}`

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

	_, err := transp.CommentChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_CommentWrite(t *testing.T) {
	input := models.CommentOutside{
		CommentID: 7,
		Message:   "gdfv df",
		Order:     0,
		User:      models.UserOutsideShort{},
	}

	expectedResponse := models.ResponseComment{
		Status:    200,
		CommentID: 7,
		Message:   "gdfv df",
		Order:     0,
		User:      models.UserOutsideShort{},
	}

	transp := &transport{}

	response, _ := transp.CommentWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_ChecklistChangeRead(t *testing.T) {
	userJSON := `{"checklistID":5,"taskID":4,"checklistName":"ffff"}`

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

	expectedUserInput := models.ChecklistInput{
		UserID:      1,
		SessionID:   "13",
		BoardID:     4,
		ChecklistID: 5,
		TaskID:      4,
		Name:        "ffff",
		Items:       nil,
	}

	transp := &transport{}

	userInput, _ := transp.ChecklistChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_ChecklistChangeReadFail(t *testing.T) {
	userJSON := `{"checklistID":5,"taskID":4,"checklistName":1}`

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

	_, err := transp.ChecklistChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_ChecklistWrite(t *testing.T) {
	input := models.ChecklistOutside{
		ChecklistID: 6,
		Name:        "check",
		Items:       nil,
	}

	expectedResponse := models.ResponseChecklist{
		Status:      200,
		ChecklistID: 6,
		Name:        "check",
		Items:       nil,
	}

	transp := &transport{}

	response, _ := transp.ChecklistWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_AttachmentAddRead(t *testing.T) {
	params := make(map[string]string, 0)
	params["attachmentFileName"] = "ttt"
	params["taskID"] = "2"

	e := echo.New()
	body, writer, err := fileUploadRequest(params, "file", "../../../default/default_avatar.png")
	if err != nil {
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))

	expectedUserInput := models.AttachmentInput{
		UserID:    1,
		SessionID: "13",
		BoardID:   4,
		TaskID:    2,
		Filename:  "ttt",
	}

	file, _ := os.Open("../../../default/default_avatar.png")
	f, _ := file.Stat()
	var byteContainer []byte
	byteContainer = make([]byte, f.Size())
	byteContainer, _ = ioutil.ReadAll(file)
	expectedUserInput.File = byteContainer

	transp := &transport{}

	userInput, err := transp.AttachmentAddRead(c)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_AttachmentDeleteRead(t *testing.T) {
	userJSON := `{"attachmentID":5,"taskID":4,"attachmentFileName":"ffff"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetPath("/card/:ID")
	c.SetParamNames("ID")
	c.SetParamValues("4")
	c.Set("sessionID", "13")
	c.Set("boardID", int64(4))

	expectedUserInput := models.AttachmentInput{
		UserID:       1,
		SessionID:    "13",
		BoardID:      4,
		TaskID:       4,
		AttachmentID: 5,
		Filename:     "ffff",
		File:         nil,
	}
	transp := &transport{}

	userInput, _ := transp.AttachmentDeleteRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_AttachmentDeleteReadFail(t *testing.T) {
	userJSON := `{"attachmentID":5,"taskID":4,"attachmentFileName":1}`

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

	_, err := transp.AttachmentDeleteRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_AttachmentWrite(t *testing.T) {
	input := models.AttachmentOutside{
		AttachmentID: 11,
		Filename:     "ttt",
		Filepath:     "ppp",
	}

	expectedResponse := models.ResponseAttachment{
		Status:       200,
		AttachmentID: 11,
		Filename:     "ttt",
		Filepath:     "ppp",
	}

	transp := &transport{}

	response, _ := transp.AttachmentWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_UserShortWrite(t *testing.T) {
	input := models.UserOutsideShort{
		ID:       0,
		Email:    "mam",
		Username: "kit",
		FullName: "tik",
		Avatar:   "fdf",
	}

	expectedResponse := models.ResponseUser{
		Status:   200,
		Email:    "mam",
		Username: "kit",
		FullName: "tik",
		Avatar:   "fdf",
	}

	transp := &transport{}

	response, _ := transp.UserShortWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_URLRead(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetParamNames("ID", "url")
	c.SetParamValues("4", "12345")

	expectedUserInput := models.BoardInviteInput{
		UserID:  1,
		BoardID: 4,
		UrlHash: "12345",
	}

	transp := &transport{}

	userInput, _ := transp.URLRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_URLReadFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))
	c.SetParamNames("ID", "url")
	c.SetParamValues("J", "12345")

	transp := &transport{}

	_, err := transp.URLRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_URLWrite(t *testing.T) {
	input := "12345"

	expectedResponse := models.ResponseURL{
		Status: 200,
		URL:    "12345",
	}

	transp := &transport{}

	response, _ := transp.URLWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func fileUploadRequest(params map[string]string, paramName, path string) (io.Reader, *multipart.Writer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, nil, err
	}

	return body, writer, nil
}
