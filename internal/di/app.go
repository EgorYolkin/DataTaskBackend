package di

import (
	"DataTask/internal/config"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config   *config.Config
	Database *sql.DB

	PrometheusHandler gin.HandlerFunc
	SwaggerHandler    gin.HandlerFunc
}

func InitializeApp(cfg *config.Config) (*App, error) {
	prometheusHandler := InitializePrometheusHandler()
	swaggerHandler := InitializeSwaggerHandler(cfg, "DataTask")

	return &App{
		Config: cfg,

		PrometheusHandler: prometheusHandler,
		SwaggerHandler:    swaggerHandler,
	}, nil
}
