package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strings"

	"github.com/AdiKhoironHasan/bookservices-users/config"
)

type Token string

var (
	cfg        *config.Config
	secretKey  string
	serviceKey string
)

func EncodeBasicAuth(username, password string) string {
	token := base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{username, password}, ":")))

	return token
}

func DecodeBasicAuth(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func ValidateToken(existToken string) bool {
	if val, exist := os.LookupEnv("APP_SECRET_KEY"); exist {
		secretKey = val
	}

	newToken := GenerateHMACToken(secretKey)

	return newToken == existToken
}

func GenerateHMACToken(secretKey string) string {
	key := []byte(secretKey)

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(secretKey))
	hmac := mac.Sum(nil)

	return hex.EncodeToString(hmac)
}
