package router

import (
	"github.com/DaniloFaraum/studere-backend/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default() //Creates router from gin library

	router.Use(config.SetupCORS())

	InitializeRoutes(router)

	router.Run(":8080") //Starts the server on api/v1/8080
}