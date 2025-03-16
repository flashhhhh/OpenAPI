package jwt

import (
	"errors"
	"fmt"
	cfg "go-swag/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Info []map[string]string `json:"info"`
	jwt.RegisteredClaims
}

func GenerateToken(information []map[string]string) (string, error) {
	c, err := cfg.LoadConfig()
	if err != nil {
		return "", errors.New("Failed to load config")
	}

	jwtConfig := c.GetJWTConfig()
	secretKey := jwtConfig.SecretKey
	// tokenExpiry := jwtConfig.TokenExpiry

	expiredTime := time.Now().Add(time.Hour * 24)

	fmt.Println("Information:", information)

	claims := Claims{
		Info: information,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) ([]map[string]string, error) {
	c, err := cfg.LoadConfig()
	if err != nil {
		return nil, errors.New("Failed to load config")
	}

	jwtConfig := c.GetJWTConfig()
	secretKey := jwtConfig.SecretKey
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Failed to parse token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("Claims:", claims)

	if !ok {
		return nil, errors.New("Failed to parse claims")
	}

	// Check if token is expired
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("Token is expired")
		}
	}

	infoList, ok := claims["info"].([]interface{})
	if !ok {
		return nil, errors.New("Failed to parse info")
	}

	var info []map[string]string
	for _, item := range infoList {
		infoMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("Failed to parse info item")
		}

		stringMap := make(map[string]string)
		for key, value := range infoMap {
			stringMap[key] = value.(string)
		}
		info = append(info, stringMap)
	}

	fmt.Println("Info:", info)

	return info, nil
}