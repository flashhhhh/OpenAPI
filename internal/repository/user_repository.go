package repository

import (
	"errors"
	"go-swag/pkg/gorm"
)

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique, not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

func CreateUser(username, password, name string) (User, error) {
	db, err := gorm.NewGormDB()
	if err != nil {
		return User{}, errors.New("Failed to connect to database")
	}

	user := User{
		Username: username,
		Password: password,
		Name:     name,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return User{}, errors.New("Failed to create user")
	}

	return user, nil
}

func GetUserByID(id int) (User, error) {
	db, err := gorm.NewGormDB()
	if err != nil {
		return User{}, errors.New("Failed to connect to database")
	}

	var user User
	db.First(&user, id)
	
	if user.ID == 0 {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	db, err := gorm.NewGormDB()
	if err != nil {
		return User{}, errors.New("Failed to connect to database")
	}

	var user User
	db.Where("username = ?", username).First(&user)
	
	if user.ID == 0 {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func GetAllUsers() ([]User, error) {
	db, err := gorm.NewGormDB()
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}

	var users []User
	db.Find(&users)

	return users, nil
}