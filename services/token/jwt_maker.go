package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 12

type JWTMaker struct {
	secretKey string
}

func NewJwtMaker(jwtsecret string) (Maker, error) {
	// secretKey:=  os.Getenv("JWT_SECRET")
	if len(jwtsecret) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secretkey : must be at least %d characters long", minSecretKeySize)
	}

	return &JWTMaker{secretKey: jwtsecret}, nil
}

func (jwtMaker *JWTMaker) CreateToken(username, userId string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, userId, duration)
	if err != nil {
		return "", nil, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(jwtMaker.secretKey))
	if err != nil {
		return "", nil, err
	}
	return token, payload, err

}
func (jwtMaker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtMaker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil

}
