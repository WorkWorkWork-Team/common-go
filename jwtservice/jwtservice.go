package jwtservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(citizenID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	CitizenID string `json:"CitizenID"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
	ttl       time.Duration
}

func NewJWTService(secretKey string, issuer string, ttl time.Duration) JWTService {
	return &jwtServices{
		secretKey: secretKey,
		issuer:    issuer,
		ttl:       ttl,
	}
}

func (service *jwtServices) GenerateToken(citizenID string) (string, error) {
	claims := &authCustomClaims{
		citizenID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(service.ttl).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	return t, err
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New(fmt.Sprint("Invalid token", token.Header["alg"]))
		}
		return []byte(service.secretKey), nil
	})

}
