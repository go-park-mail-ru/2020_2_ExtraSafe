package sources

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Handlers struct {
	echo.Context
	Users    *[]User
	Sessions *map[string]uint64 //map[sessionID]userID
}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
}

type UserLinks struct {
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

type UserInputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInputReg struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInputProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserInputPassword struct {
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

func (h *Handlers) checkUserAuthorized(c echo.Context) (ResponseUser, error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return ResponseUser{}, ResponseError{Status: 500, Codes: nil}
	}
	sessionID := session.Value
	userID, authorized := (*h.Sessions)[sessionID]

	if authorized {
		for _, user := range *h.Users {
			if user.ID == userID {
				response := new(ResponseUser)
				response.WriteResponse(user)
				return *response, nil
			}
		}
	}
	return ResponseUser{}, ResponseError{Status: 500, Codes: nil}
}

func (h *Handlers) CheckUser(userInput UserInputLogin) (ResponseUser, uint64, error) {
	response := new(ResponseUser)
	for _, user := range *h.Users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			response.WriteResponse(user)
			return *response, user.ID, nil
		}
	}

	errorCodes := []string{"101"}
	return ResponseUser{}, 0, ResponseError{Codes: errorCodes, Status: 500}
}

func (h *Handlers) CreateUser(userInput UserInputReg) (ResponseUser, uint64, error) {
	errorCodes := make([]string, 0)
	for _, user := range *h.Users {
		if userInput.Email == user.Email {
			errorCodes = append(errorCodes, "201")
		}

		if userInput.Username == user.Username {
			errorCodes = append(errorCodes, "202")
		}
	}

	if len(errorCodes) != 0 {
		return ResponseUser{}, 0, ResponseError{Codes: errorCodes, Status: 500}
	}

	var id uint64 = 0
	if len(*h.Users) > 0 {
		id = (*h.Users)[len(*h.Users)-1].ID + 1
	}

	newUser := User{
		ID:       id,
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	}

	*h.Users = append(*h.Users, newUser)

	response := new(ResponseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (h *Handlers) ChangeUserProfile(userInput *UserInputProfile, userExist *User) (ResponseUser, error) {
	errorCodes := make([]string, 0)
	for _, user := range *h.Users {
		if (userInput.Email == user.Email) && (user.ID != userExist.ID) {
			errorCodes = append(errorCodes, "301")
		}

		if (userInput.Username == user.Username) && (user.ID != userExist.ID) {
			errorCodes = append(errorCodes, "302")
		}
	}

	if len(errorCodes) != 0 {
		return ResponseUser{}, ResponseError{Codes: errorCodes, Status: 500}
	}

	response := new(ResponseUser)

	userExist.Username = userInput.Username
	userExist.Email = userInput.Email
	userExist.FullName = userInput.FullName

	response.WriteResponse(*userExist)
	return *response, nil
}

func (h *Handlers) ChangeUserAccounts(userInput *UserLinks, userExist *User) (ResponseUserLinks, error) {
	userExist.Links.Bitbucket = userInput.Bitbucket
	userExist.Links.Github = userInput.Github
	userExist.Links.Instagram = userInput.Instagram
	userExist.Links.Telegram = userInput.Telegram
	userExist.Links.Facebook = userInput.Facebook
	userExist.Links.Vk = userInput.Vk

	response := new(ResponseUserLinks)
	response.WriteResponse(userExist.Username, *userExist.Links, userExist.Avatar)

	return *response, nil
}

func (h *Handlers) ChangeUserPassword(userInput *UserInputPassword, userExist *User) (ResponseUser, error) {
	if userInput.OldPassword != userExist.Password {
		errorCodes := []string{"501"}
		return ResponseUser{}, ResponseError{Codes: errorCodes, Status: 500}
	}

	userExist.Password = userInput.Password

	response := new(ResponseUser)
	response.WriteResponse(*userExist)

	return *response, nil
}

func getFormParams(params url.Values) (userInput *UserInputProfile) {
	userInput = new(UserInputProfile)
	userInput.Username = params.Get("username")
	userInput.Email = params.Get("email")
	userInput.FullName = params.Get("fullName")

	return
}

func (h *Handlers) UploadAvatar(file *multipart.FileHeader, userID uint64) (err error, filename string) {
	src, err := file.Open()
	if err != nil {
		return ResponseError{Codes: []string{"401"}, Status: 500}, ""
	}
	defer src.Close()

	oldAvatar := (*h.Users)[userID].Avatar

	hash := sha256.New()

	formattedTime := strings.Join(strings.Split(time.Now().String(), " "), "")
	formattedID := strconv.FormatUint(userID, 10)
	name := fmt.Sprintf("%x", hash.Sum([]byte(formattedTime+formattedID)))

	filename, err = saveImage(&src, name)
	if err != nil {
		return ResponseError{Codes: []string{"402"}, Status: 500}, ""
	}

	(*h.Users)[userID].Avatar = "avatars/" + filename

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

	dst, err := os.Create("./avatars/" + filename)
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
