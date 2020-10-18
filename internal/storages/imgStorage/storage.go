package imgStorage

import (
	"../../../internal/models"
	"../../../internal/errorWorker"
	"crypto/sha256"
	"errors"
	"fmt"
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
	UploadAvatar(file *multipart.FileHeader, userID uint64) (err error, filename string)
}

type storage struct {
	Sessions *map[string]uint64
	Users    *[]models.User
}

func NewStorage(someUsers *[]models.User, sessions *map[string]uint64) Storage {
	return &storage{
		Sessions: sessions,
		Users: someUsers,
	}
}

func (s *storage) UploadAvatar(file *multipart.FileHeader, userID uint64) (err error, filename string) {
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return errorWorker.ResponseError{Codes: []string{"401"}, Status: 500}, ""
	}
	defer src.Close()

	oldAvatar := (*s.Users)[userID].Avatar

	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatUint(userID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID)))

	filename, err = saveImage(&src, name)
	if err != nil {
		fmt.Println(err)
		return errorWorker.ResponseError{Codes: []string{"402"}, Status: 500}, ""
	}

	(*s.Users)[userID].Avatar = "avatars/" + filename

	if oldAvatar != "default/default_avatar.png" {
		os.Remove("./" + oldAvatar)
	}

	return nil, filename
}

func saveImage(src *multipart.File, name string) (string, error) {
	img, fmtName, err := image.Decode(*src)
	if err != nil {
		return "", err
	}

	filename := name + "." + fmtName

	//TODO - нормальный путь
	dst, err := os.Create("/home/keith/tehnopark/go/2020_2_ExtraSafe/avatars/" + filename)
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


