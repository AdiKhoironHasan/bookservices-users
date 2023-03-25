package assembler

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = "secret"
)

// Function to generate a JWT token
func GenerateToken(email string) (string, error) {
	// Set the expiration time for the token
	tokenExpires := time.Now().Add(time.Hour * 2).Unix()

	// Create the JWT claims, which includes the email and token expiration time
	claims := jwt.MapClaims{
		"email": email,
		"exp":   tokenExpires,
	}

	// Create the token using the claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Function to refresh a JWT token
func RefreshToken(tokenString string) (string, error) {
	// Parse the token string to get the claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	// Check if the token is valid and hasn't expired yet
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Get the email from the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email in token claims")
	}

	// Create a new token with a new expiration time
	tokenExpires := time.Now().Add(time.Hour * 2).Unix()
	newClaims := jwt.MapClaims{
		"email": email,
		"exp":   tokenExpires,
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	signedToken, err := newToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Function to check if a JWT token is expired
func IsTokenExpired(tokenString string) bool {
	// Parse the token string to get the claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return true
	}

	// Check if the token is expired
	if !token.Valid {
		return true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	tokenExpires, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	return int64(tokenExpires) < time.Now().Unix()
}

func RemoveToken(tokenString string) error {
	// Parse the JWT token string

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You need to provide a key to verify the signature of the token
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}

	// Invalidate the token by setting its expiry time to a past time
	token.Claims.(jwt.MapClaims)["exp"] = 0

	return nil
}
