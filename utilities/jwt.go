package utilities

import (
	"backend-mental-guardians/constants"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func parseToken(token string) (*jwt.Token, error) {

	tokenString := strings.Split(token, " ")[1]

	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return nil, constants.ErrInvalidToken
	}
	return result, nil
}

func GetUserIdFromToken(token string) (int, error) {
	result, err := parseToken(token)
	if err != nil {
		return 0, err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := int(claims["userId"].(float64))

		return userID, err
	}
	return 0, constants.ErrInvalidToken
}