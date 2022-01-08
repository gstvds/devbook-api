package authentication

import (
	"api/src/utils/config"
	"github.com/dgrijalva/jwt-go"
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
