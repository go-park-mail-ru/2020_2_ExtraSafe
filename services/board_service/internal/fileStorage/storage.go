package fileStorage

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//go:generate mockgen -destination=../../../board_service/internal/service/mock/mock_fileStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/fileStorage FileStorage

type FileStorage interface {
	UploadFile(file models.AttachmentFileInternal, attachment *models.AttachmentInternal, isTest bool) error
	DeleteFile(filepath string, isTest bool) error
}

type storage struct {}

func NewStorage() FileStorage {
	return &storage{}
}

func (s *storage) UploadFile(file models.AttachmentFileInternal, attachment *models.AttachmentInternal, isTest bool) error {
	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatInt(file.UserID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID+file.Filename)))

	filename, err := saveFile(file.File, name, file.Filename, isTest)

	if err != nil {
		return models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"},
			MethodName: "UploadFile"}
	}

	attachment.Filepath = "attachments/" + filename
	return nil
}

func saveFile(src []byte, name string, initialName string, isTest bool) (filename string, err error) {
	fmtName := filepath.Ext(initialName)
	filename = name + fmtName

	var dst *os.File
	if isTest {
		dst, err = os.Create("../../../../attachments/" + filename)
	} else {
		dst, err = os.Create("../../../attachments/" + filename)
	}

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer dst.Close()

	_, err = dst.Write(src)
	if err != nil {
		fmt.Println(err)
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
