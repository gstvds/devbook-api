package authentication

import (
	"api/src/utils/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GenerateToken generates an JSON Web Token for a User
func GenerateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SECRET_KEY)
}

// ValidateToken validates if a Token is valid
func ValidateToken(request *http.Request) error {
	tokenString := extractToken(request)

	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// ExtractUserId extract a userId from a given token
func ExtractUserId(request *http.Request) (uint64, error) {
	tokenString := extractToken(request)

	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f",claims["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}

// extractToken extracts a JSON Web Token from the request Headers
func extractToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// getValidationKey checks if the signing method is valid and then returns the SECRET_KEY
func getValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("invalid signature method %v", token.Header["alg"])
	}

	return config.SECRET_KEY, nil
}
