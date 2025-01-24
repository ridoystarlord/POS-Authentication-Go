package utils

import (
	"authentication/config"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id uint, exp int, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
		})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GenerateTokenPair(id uint) (string, string, error) {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	accessTokenDuration, err := strconv.Atoi(config.JWTAccessExp)
	if err != nil {
		log.Fatalf("Invalid env format: %v", err)
	}
	accessToken, err := CreateToken(id, accessTokenDuration, config.JWTAccessSecret)
	if err != nil {
		return "", "", err
	}
	refreshTokenDuration, err := strconv.Atoi(config.JWTRefreshExp)
	if err != nil {
		log.Fatalf("Invalid env format: %v", err)
	}
	refreshToken, err := CreateToken(id, refreshTokenDuration, config.JWTRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
