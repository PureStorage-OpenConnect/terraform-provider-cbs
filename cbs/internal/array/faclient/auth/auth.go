package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const defaultExpTimeSec = 315360000 //10 years

type TokenConfig struct {
	Issuer   string
	ClientID string
	KeyID    string

	User string

	PrivateKey string
}

func (tc *TokenConfig) BuildToken() (string, error) {

	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(tc.PrivateKey))
	if err != nil {
		return "", err
	}

	issued := time.Now()
	expires := issued.Add(time.Second * defaultExpTimeSec)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud": tc.ClientID,
		"iss": tc.Issuer,
		"sub": tc.User,
		"iat": issued.Unix(),
		"exp": expires.Unix(),
	})
	token.Header["kid"] = tc.KeyID

	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
