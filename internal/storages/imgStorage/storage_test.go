package imgStorage

import (
	"bytes"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStorage_UploadAvatar(t *testing.T) {
	path := "../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	_, err = file.Stat()
	if err != nil {
		t.Error("cannot get stat of test file")
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("avatar", "someName")
	if err != nil {
		t.Error("cannot create form file")
		return
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		t.Error("cannot close writer")
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/profile/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // <<< important part
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID: 1,
		Avatar: "default/default_avatar.png",
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		t.Error("cannot form file from test request")
		return
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
		return
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}

	err = os.Remove("../../../" + userAvatar.Avatar)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}

func TestStorage_UploadAvatarFailOnCreate(t *testing.T) {
	path := "../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	_, err = file.Stat()
	if err != nil {
		t.Error("cannot get stat of test file")
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("avatar", "someName")
	if err != nil {
		t.Error("cannot create form file")
		return
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		t.Error("cannot close writer")
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/profile/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // <<< important part
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID: 1,
		Avatar: "default/default_avatar.png",
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		t.Error("cannot form file from test request")
		return
	}

	err = storage.UploadAvatar(avatar, &userAvatar, false)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestStorage_UploadAvatarPngDecode(t *testing.T) {
	path := "../../../test/testPic.png"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	_, err = file.Stat()
	if err != nil {
		t.Error("cannot get stat of test file")
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("avatar", "someName")
	if err != nil {
		t.Error("cannot create form file")
		return
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		t.Error("cannot close writer")
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/profile/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // <<< important part
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID: 1,
		Avatar: "default/default_avatar.png",
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		t.Error("cannot form file from test request")
		return
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}

	err = os.Remove("../../../" + userAvatar.Avatar)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}

func TestStorage_UploadAvatarFailOnDecode(t *testing.T) {
	path := "../../../test/testPic.webp"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	_, err = file.Stat()
	if err != nil {
		t.Error("cannot get stat of test file")
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("avatar", "someName")
	if err != nil {
		t.Error("cannot create form file")
		return
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		t.Error("cannot close writer")
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/profile/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // <<< important part
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID: 1,
		Avatar: "default/default_avatar.png",
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		t.Error("cannot form file from test request")
		return
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err == nil {
		t.Error("expected error")
		return
	}
}
