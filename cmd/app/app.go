package app

import (
	"DataTask/internal/config"
	"DataTask/internal/di"
	"DataTask/pkg/logger"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {
	app, err := di.InitializeApp(cfg)
	if err != nil {
		return err
	}

	router := gin.New()
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api/v1")

	apiRouter.GET("/metrics", app.PrometheusHandler)

	routerDSN := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)

	logger.Log.WithFields(log.Fields{
		"host": cfg.HTTP.Host,
		"port": cfg.HTTP.Port,
	}).Info("running app")

	return router.Run(routerDSN)
}
