package oauthutils

import (
	"crypto/rsa"
	"os"
	
	"github.com/golang-jwt/jwt"
)

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}
