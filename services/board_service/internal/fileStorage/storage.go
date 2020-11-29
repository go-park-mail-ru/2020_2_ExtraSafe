package fileStorage

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"image"
	"os"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	UploadFile(file models.AttachmentFileInternal, attachment *models.AttachmentInternal, isTest bool) error
	DeleteFile(filepath string, isTest bool) error
}

type storage struct {}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) UploadFile(file models.AttachmentFileInternal, attachment *models.AttachmentInternal, isTest bool) error {
	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatInt(file.UserID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID+file.Filename)))

	filename, err := saveFile(file.File, name, false)

	if err != nil {
		return models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"},
			MethodName: "UploadFile"}
	}

	attachment.Filepath = "attachments/" + filename
	return nil
}

func saveFile(src []byte, name string, isTest bool) (filename string, err error) {
	r := bytes.NewReader(src)
	_, fmtName, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	filename = name + "." + fmtName

	var dst *os.File
	if isTest {
		dst, err = os.Create("../../../../attachments/" + name)
	} else {
		dst, err = os.Create("../../../attachments/" + name)
	}

	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = dst.Write(src)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *storage) DeleteFile(filepath string, isTest bool) error {
	if isTest {
		os.Remove("../../../../" + filepath)
	} else {
		os.Remove("../../../" + filepath)
	}
	return nil
}
