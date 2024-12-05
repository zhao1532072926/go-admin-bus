package utils

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	// The secret string used to generate the token
	Secret string
	// The token expiration time
	TimeoutSecond int64
}

// GenerateToken generates a JWT token
func (j *JWT) GenerateToken(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(time.Duration(j.TimeoutSecond) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

// ValidateToken validates a JWT token
func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	return token, err
}
