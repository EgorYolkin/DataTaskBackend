package di

import (
	"DataTask/internal/config"
	"DataTask/internal/controller/rest/handler/users_handler"
	"DataTask/internal/repository/user_repository"
	"DataTask/internal/usecase/user_usecase"
	"DataTask/pkg/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializePrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func InitializeSwaggerHandler(cfg *config.Config, instanceName string) gin.HandlerFunc {
	logger.Log.WithFields(log.Fields{
		"base_path": cfg.Swagger.BasePath,
		"version":   cfg.Swagger.Version,
	}).Info("set up swagger")

	return ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName(instanceName),
		ginSwagger.URL(cfg.Swagger.BasePath+"/doc/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	)
}

func InitializeUsersHandler(db *sql.DB, jwtSecretKey string) *users_handler.UsersHandler {
	repo := user_repository.NewPostgresUserRepository(db)
	useCase := user_usecase.NewUserUseCase(repo)
	handler := users_handler.NewUsersHandler(useCase, jwtSecretKey)
	return handler
}
