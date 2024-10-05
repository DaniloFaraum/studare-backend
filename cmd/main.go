package main

import (
	"github.com/DaniloFaraum/studere-backend/config"
	"github.com/DaniloFaraum/studere-backend/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initialization failed: %v", err)
		return
	}

	router.Initialize()
}