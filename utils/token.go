package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(id uint, exp int, tokenType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   id,
			"type": tokenType,
			"exp":  time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	if claims["type"] != "access" {
		return nil, fmt.Errorf("invalid token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	expTime := time.Unix(int64(exp), 0)
	if expTime.Before(time.Now()) {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	if claims["type"] != "refresh" {
		return nil, fmt.Errorf("invalid token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	expTime := time.Unix(int64(exp), 0)
	if expTime.Before(time.Now()) {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
