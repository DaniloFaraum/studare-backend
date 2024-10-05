package utils

import "github.com/DaniloFaraum/studere-backend/config"

var (
	logger *config.Logger
)

func InitializeUtils() {
	logger = config.NewLogger("utils")
}
