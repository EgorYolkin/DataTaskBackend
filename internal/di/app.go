package di

import (
	"DataTask/internal/config"
	"DataTask/internal/controller/rest/handler/users_handler"
	"DataTask/internal/controller/rest/middleware/auth_middleware"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config   *config.Config
	Database *sql.DB

	PrometheusHandler gin.HandlerFunc
	SwaggerHandler    gin.HandlerFunc

	UsersHandler *users_handler.UsersHandler

	AuthMiddleware *auth_middleware.AuthMiddleware
}

func InitializeApp(cfg *config.Config) (*App, error) {
	db, err := InitializeDatabase(cfg)

	if err != nil {
		return nil, err
	}

	prometheusHandler := InitializePrometheusHandler()
	swaggerHandler := InitializeSwaggerHandler(cfg, "DataTask")
	usersHandler := InitializeUsersHandler(db, cfg.JWT.Secret)

	authMiddleware := auth_middleware.NewAuthMiddleware(cfg.JWT.Secret)

	return &App{
		Config: cfg,

		PrometheusHandler: prometheusHandler,
		SwaggerHandler:    swaggerHandler,
		UsersHandler:      usersHandler,

		AuthMiddleware: authMiddleware,
	}, nil
}
