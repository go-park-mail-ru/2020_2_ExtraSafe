package csrf

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"strconv"
	"time"
)

type Service interface {
	GenerateToken(userID uint64) (token string, err error)
	CheckToken(userID uint64, tokenString string) (err error)
}

var mySigningKey = []byte("3ndec")

func GenerateToken(userID uint64) (token string, err error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["uid"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	return jwtToken.SignedString(mySigningKey)
}

func CheckToken(userID uint64, tokenString string) (err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"777"}, Descriptions: []string{"Token is not valid"},
			MethodName: "CheckToken"}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.ServeError{Codes: []string{"777"}, Descriptions: []string{"Token is not valid"},
			MethodName: "CheckToken"}
	}

	UID := strconv.FormatUint(userID, 10)
	tokenUID := strconv.FormatFloat(claims["uid"].(float64), 'f', 0, 64)

	if tokenUID != UID {
		return models.ServeError{Codes: []string{"777"}, Descriptions: []string{"Invalid user in token"},
			MethodName: "CheckToken"}
	}
	return nil
}