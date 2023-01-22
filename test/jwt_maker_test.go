package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fxmbx/go_practice_pr/services/token"
	"github.com/fxmbx/go_practice_pr/utils"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := token.NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)
	user := CreateRandomUser(t)
	username := user.UserName
	userId := user.ID
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payloadt, err := maker.CreateToken(username, fmt.Sprintf("%d", userId), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotZero(t, payloadt.ID)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.ID)
	require.Greater(t, payload.ExpiredAt, time.Now())

	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, payload.IssuedAt, issuedAt, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, expiredAt, time.Second)
}

func TestExpiredJwtToken(t *testing.T) {
	maker, err := token.NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)
	user := CreateRandomUser(t)

	username := user.UserName
	userId := user.ID
	duration := -time.Minute
	tokenstring, payloadt, err := maker.CreateToken(username, fmt.Sprintf("%d", userId), duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenstring)
	require.NotZero(t, payloadt.ID)

	payload, err := maker.VerifyToken(tokenstring)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidJWT(t *testing.T) {
	payload, err := token.NewPayload(utils.RandomString(6), utils.RandomString(1), time.Minute)
	require.NoError(t, err)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenString, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	// require.NotEmpty(t, token)

	maker, err := token.NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)
	payload, err = maker.VerifyToken(tokenString)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
