package imgstorage

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"io/ioutil"
	"os"
	"testing"
)

func TestStorage_UploadAvatar(t *testing.T) {
	path := "../../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	avatar, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
		return
	}

	err = os.Remove("../../../../" + userAvatar.Avatar)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}

func TestStorage_UploadAvatarFailOnCreate(t *testing.T) {
	path := "../../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	avatar, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	err = storage.UploadAvatar(avatar, &userAvatar, false)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestStorage_UploadAvatarPngDecode(t *testing.T) {
	path := "../../../../test/testPic.png"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	avatar, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}

	err = os.Remove("../../../../" + userAvatar.Avatar)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}

func TestStorage_UploadAvatarFailOnDecode(t *testing.T) {
	path := "../../../../test/testPic.webp"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	avatar, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	err = storage.UploadAvatar(avatar, &userAvatar, true)
	if err == nil {
		t.Error("expected error")
		return
	}
}
