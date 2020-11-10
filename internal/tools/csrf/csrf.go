package csrf

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("Token is not valid ")
	}
	UID := strconv.FormatUint(userID, 10)
	tokenUID := strconv.FormatFloat(claims["uid"].(float64), 'f', 0, 64)
	fmt.Printf("%T\n", tokenUID)
	fmt.Println(tokenUID)
	if tokenUID != UID {
		fmt.Println(1111111)
		return errors.New("Invalid user in token ")
	}
	return nil
}