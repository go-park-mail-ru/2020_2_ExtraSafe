package imgstorage

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:generate mockgen -destination=../../../profile_service/internal/service/mock/mock_avatarStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgStorage AvatarStorage

type AvatarStorage interface {
	UploadAvatar(file []byte, user *models.UserAvatar, isTest bool) error
}

type storage struct{}

func NewStorage() AvatarStorage {
	return &storage{}
}

func (s *storage) UploadAvatar(file []byte, user *models.UserAvatar, isTest bool) error {
	oldAvatar := user.Avatar
	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatInt(user.ID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID)))

	filename, err := saveImage(file, name, isTest)
	if err != nil {
		return models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"},
			MethodName: "UploadAvatar"}
	}

	user.Avatar = "avatars/" + filename
	if oldAvatar != "default/default_avatar.png" {
		if isTest {
			_ = os.Remove("../../../../" + oldAvatar)
		} else {
			_ = os.Remove("../../../" + oldAvatar)
		}
	}

	return nil
}

func saveImage(src []byte, name string, isTest bool) (string, error) {
	r := bytes.NewReader(src)
	img, fmtName, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	filename := name + "." + fmtName

	var dst *os.File
	if isTest {
		dst, err = os.Create("../../../../avatars/" + filename)
	} else {
		dst, err = os.Create("../../../avatars/" + filename)
	}

	if err != nil {
		return "", err
	}
	defer dst.Close()

	switch fmtName {
	case "png":
		{
			enc := png.Encoder{
				CompressionLevel: png.BestCompression,
			}
			err := enc.Encode(dst, img)
			if err != nil {
				return "", err
			}
		}
	case "jpeg":
		{
			opt := jpeg.Options{
				Quality: 60,
			}

			err = jpeg.Encode(dst, img, &opt)
			if err != nil {
				return "", err
			}
		}
	default:
		return "", errors.New("Wrong file type ")
	}
	return filename, nil
}
