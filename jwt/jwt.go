package jwt

import (
	"errors"
	"net/http"
	"rest/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte("SecretYouShouldHide")

func GenerateJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":  time.Now().Add(time.Minute * 10).Unix(),
			"user": user.ID,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(r *http.Request) (uuid.UUID, error) {
	var userId uuid.UUID

	if r.Header["Authorization"] != nil {
		token, _ := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if token == nil {
			return userId, errors.New("incorrect token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userId, _ := uuid.Parse(claims["user"].(string))
			return userId, nil
		}
		return userId, errors.New("invalid token")
	}
	return userId, errors.New("unknown token \"Authorization\"")
}
