package service

import (
	"errors"
	"go-learn/global"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims structure
type CustomClaims struct {
	ID         uint
	Username   string
	NickName   string
	BufferTime int64
	jwt.StandardClaims
}

// JWT struct
type JWT struct {
	SigningKey []byte
}

// NewJWT generate
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SigningKey),
	}
}

// CreateToken function
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken function
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("Token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Token is invalid")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("Token is invalid")
}
