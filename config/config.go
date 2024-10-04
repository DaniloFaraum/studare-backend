package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error { //Starts the Database
	var err error

	db, err = InitMySQL()
	if err != nil {
		return fmt.Errorf("could not initialize db: %v", err)
	}
	return nil
}

func GetMySQL() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
