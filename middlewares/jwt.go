package middlewares

import (
	"backend-mental-guardians/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint) (string, error) {
	//membuat payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	//membuat header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//verify signature
	return token.SignedString([]byte(configs.InitConfigJWT()))
}
