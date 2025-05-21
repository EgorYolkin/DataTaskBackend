package main

import (
	"DataTask/cmd/app"
	"DataTask/internal/config"
	"DataTask/pkg/logger"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
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

	isDockerMode := isDocker()
	envFilePath := "infra/config/.env"

	log.Info(
		fmt.Sprintf("docker mode: %v", isDockerMode),
	)

	if isDockerMode {
		envFilePath = "infra/config/.env.docker"
	}

	cfg, err := config.NewConfig(
		envFilePath,
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

func isDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
