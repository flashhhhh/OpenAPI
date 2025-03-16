package services

import (
	"errors"
	"fmt"
	"go-swag/internal/repository"
	"go-swag/pkg/hash"
	"go-swag/pkg/jwt"
)

func Signup(username, password, name string) (error) {
	hashPassword, err := hash.HashPassword(password)
	if err != nil {
		return errors.New("Failed to hash password")
	}

	_, err = repository.CreateUser(username, hashPassword, name)
	if err != nil {
		return errors.New("Failed to create user")
	}

	return nil
}

func Login(username, password string) error {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return errors.New("User not found")
	}

	if !hash.CheckPassword(password, user.Password) {
		return errors.New("Invalid password")
	}

	return nil
}

func GenerateToken(info []map[string]string) (string, error) {
	token, err := jwt.GenerateToken(info)
	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return token, nil
}

func ValidateToken(token string) error {
	info, err := jwt.ValidateToken(token)
	if err != nil {
		return errors.New("Failed to validate token")
	}

	fmt.Println("Hello ", info[0]["username"])

	return nil
}