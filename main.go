package main

import (
	"DataTask/cmd/app"
	"DataTask/internal/config"
	"DataTask/pkg/logger"
)

// @title           DataTask
// @version         1.0
// @description     Task manager
// @BasePath        /api/v1

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

// @security Authorization
func main() {
	logger.InitLogger()

	logger.Log.Info("Starting DataTask")

	cfg, err := config.NewConfig(
		"infra/config/.env",
		"infra/config",
		"config",
	)
	if err != nil {
		logger.Log.Error(err)
	}

	err = app.Run(cfg)
	if err != nil {
		logger.Log.Error(err)
	}
}
