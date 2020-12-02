package fileStorage

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"io/ioutil"
	"os"
	"testing"
)

func TestStorage_UploadFile(t *testing.T) {
	path := "../../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileb, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()

	fileInput := models.AttachmentFileInternal{
		UserID:   1,
		Filename: "filename",
		File:     fileb,
	}

	attachment := models.AttachmentInternal{
		TaskID:       1,
		AttachmentID: 0,
		Filename:     fileInput.Filename,
		Filepath:     "",
	}

	err = storage.UploadFile(fileInput, &attachment, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
		return
	}

	err = os.Remove("../../../../" + attachment.Filepath)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}

func TestStorage_DeleteFile(t *testing.T) {
	path := "../../../../test/testPic.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error("cannot open test file")
		return
	}
	defer file.Close()

	fileb, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("cannot read test file")
		return
	}

	storage := NewStorage()

	fileInput := models.AttachmentFileInternal{
		UserID:   1,
		Filename: "filename",
		File:     fileb,
	}

	attachment := models.AttachmentInternal{
		TaskID:       1,
		AttachmentID: 0,
		Filename:     fileInput.Filename,
		Filepath:     "",
	}

	err = storage.UploadFile(fileInput, &attachment, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
		return
	}

	err = storage.DeleteFile(attachment.Filepath, true)
	if err != nil {
		t.Errorf("unexpected error: %s ", err)
	}
	return
}