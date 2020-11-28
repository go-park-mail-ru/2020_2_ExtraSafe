package fileStorage

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"os"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	UploadFile(file models.AttachmentFileInput, attachment *models.AttachmentInput, isTest bool) error
}

type storage struct {}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) UploadFile(file models.AttachmentFileInput, attachment *models.AttachmentInput, isTest bool) error {
	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatInt(file.UserID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID+file.Filename)))

	var dst *os.File
	var err error
	if isTest {
		dst, err = os.Create("../../../../attachments/" + name)
	} else {
		dst, err = os.Create("../../../attachments/" + name)
	}

	if err != nil {
		return models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"},
			MethodName: "UploadFile"}
	}
	defer dst.Close()

	attachment.Filepath = "attachments/" + file.Filename
	return nil
}
