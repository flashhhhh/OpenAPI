package gorm

import (
	"fmt"
	cfg "go-swag/configs"
	"go-swag/pkg/logger"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Cache gorm.DB instance
var db *gorm.DB

func NewGormDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	config_data, err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}

	databaseConfig := config_data.GetDatabaseConfig()

	host := databaseConfig.Host
	port := databaseConfig.Port
	user := databaseConfig.User
	password := databaseConfig.Password
	dbname := databaseConfig.DBname
	dsn := "host=" + host + "  user=" + user + " port=" + strconv.Itoa(port)  + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	logger.Info("Connecting to database...")
	fmt.Println("DSN:", dsn)

	err = nil
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}