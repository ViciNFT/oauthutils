package oauthutils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
)

var verifyKey *rsa.PublicKey

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	verifyKey = loadRSAPublicKeyFromDisk(location)
	return verifyKey
}

func loadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
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

func GetKey(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
