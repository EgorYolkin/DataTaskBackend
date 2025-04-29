package app

import (
	"DataTask/internal/config"
	"DataTask/internal/di"
	"DataTask/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) error {
	app, err := di.InitializeApp(cfg)
	if err != nil {
		return err
	}

	router := gin.New()
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api/v1")
	protectedApiRouter := apiRouter.Group("/")
	protectedApiRouter.Use(app.AuthMiddleware.Middleware())

	apiRouter.GET("/metrics", app.PrometheusHandler)

	usersHandlerRouterGroup := apiRouter.Group("/users")
	{
		usersHandlerRouterGroup.POST("/create", app.UsersHandler.HandleCreateUser)
		usersHandlerRouterGroup.POST("/update", app.UsersHandler.HandleUpdateUser)
		usersHandlerRouterGroup.POST("/delete", app.UsersHandler.HandleDeleteUser)
	}

	authHandlerRouterGroup := apiRouter.Group("/auth")
	{
		authHandlerRouterGroup.POST("/login", app.UsersHandler.LoginUserHandler)
	}

	protectedUsersRouterGroup := protectedApiRouter.Group("/users")
	{
		protectedUsersRouterGroup.POST("/me", app.UsersHandler.HandleGetCurrentUser)
	}

	routerDSN := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)

	logger.Log.WithFields(log.Fields{
		"host": cfg.HTTP.Host,
		"port": cfg.HTTP.Port,
	}).Info("running app")

	return router.Run(routerDSN)
}
