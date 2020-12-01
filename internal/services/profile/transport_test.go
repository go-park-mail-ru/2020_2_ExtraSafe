package profile

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

func TestTransport_ProfileRead(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	expectedUserInput := models.UserInput{ID: 1}

	transp := &transport{}

	userInput, _ := transp.ProfileRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_BoardsWrite(t *testing.T) {
	input := []models.BoardOutsideShort{models.BoardOutsideShort{BoardID: 1, Name: "ex", Theme: "dark", Star: true}}

	expectedResponse:= models.ResponseBoards{
		Status:   200,
		Boards:   []models.BoardOutsideShort{models.BoardOutsideShort{BoardID: 1, Name: "ex", Theme: "dark", Star: true}},
	}

	transp := NewTransport()

	response, _ := transp.BoardsWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_ProfileWrite(t *testing.T) {
	input := models.UserOutside{
		Email:    "mariya@mail.ru",
		Username: "kit",
		FullName: "kit tik",
		Avatar:   "pig.png",
	}

	expectedResponse:= models.ResponseUser{
		Status:   200,
		Email:    "mariya@mail.ru",
		Username: "kit",
		FullName: "kit tik",
		Avatar:   "pig.png",
	}

	transp := &transport{}

	response, _ := transp.ProfileWrite(input)

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestTransport_PasswordChangeRead(t *testing.T) {
	userJSON := `{"oldpassword":"1234","password":"12345"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	expectedUserInput := models.UserInputPassword{
		ID:          1,
		OldPassword: "1234",
		Password:    "12345",
	}

	transp := &transport{}

	userInput, _ := transp.PasswordChangeRead(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_PasswordChangeReadFail(t *testing.T) {
	userJSON := `{"oldpassword":1234,"password":"12345"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	transp := &transport{}

	_, err := transp.PasswordChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
		return
	}
}

func TestTransport_ProfileChangeRead(t *testing.T) {
	params := make(map[string]string, 0)
	params["email"] = "mam"
	params["username"] = "kit"
	params["fullName"] = "gop"

	e := echo.New()
	body, writer, err := fileUploadRequest(params, "avatar", "../../../default/default_avatar.png")
	if err != nil{
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	expectedUserInput := models.UserInputProfile{
		ID:       1,
		Email:    "mam",
		Username: "kit",
		FullName: "gop",
		Avatar:   nil,
	}

	file, _ := os.Open("../../../default/default_avatar.png")
	f, _ := file.Stat()
	var byteContainer []byte
	byteContainer = make([]byte, f.Size())
	byteContainer, _ = ioutil.ReadAll(file)
	expectedUserInput.Avatar = byteContainer

	transp := &transport{}

	userInput, err := transp.ProfileChangeRead(c)

	if !reflect.DeepEqual(userInput, expectedUserInput) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserInput, userInput)
		return
	}
}

func TestTransport_ProfileChangeReadFail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", int64(1))

	transp := &transport{}

	_, err := transp.ProfileChangeRead(c)

	if err == nil {
		t.Errorf("results not match")
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