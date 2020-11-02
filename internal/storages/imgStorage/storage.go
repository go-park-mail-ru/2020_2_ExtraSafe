package imgStorage

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	UploadAvatar(file *multipart.FileHeader, user *models.UserAvatar) error
}

type storage struct {
	Users    *[]models.User
}

func NewStorage(someUsers *[]models.User) Storage {
	return &storage{
		Users: someUsers,
	}
}

func (s *storage) UploadAvatar(file *multipart.FileHeader, user *models.UserAvatar) error {
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"401"}}
	}
	defer src.Close()

	oldAvatar := user.Avatar
	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatUint(user.ID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID)))

	filename, err := saveImage(&src, name)
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"402"}}
	}

	user.Avatar = "avatars/" + filename
	if oldAvatar != "default/default_avatar.png" {
		os.Remove("../" + oldAvatar)
	}

	return nil
}

func saveImage(src *multipart.File, name string) (string, error) {
	img, fmtName, err := image.Decode(*src)
	if err != nil {
		return "", err
	}

	filename := name + "." + fmtName

	dst, err := os.Create("../avatars/" + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	switch fmtName {
	case "png":
		{
			enc := png.Encoder{
				CompressionLevel: png.BestCompression, //еще не реализовано
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
		return "", errors.New("Неверный формат файла ")
	}
	return filename, nil
}


