package controllers

import (
	"github.com/DaniloFaraum/studere-backend/config"
	"gorm.io/gorm"
)

var (
	logger config.Logger
	db *gorm.DB
)

func InitializeController(){
	logger = *config.GetLogger("handler")
	db = config.GetMySQL()
}