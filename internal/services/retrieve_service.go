package services

import (
	"errors"
	"go-swag/internal/repository"
)

func GetAllUsers() ([]repository.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, errors.New("Failed to get users")
	}

	return users, nil
}

func GetUserByID(id int) (repository.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return repository.User{}, errors.New("Failed to get user")
	}

	return user, nil
}