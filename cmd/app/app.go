package app

import (
	"DataTask/internal/config"
	"DataTask/internal/di"
	"DataTask/pkg/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
	//"time"
)

func Run(cfg *config.Config) error {
	app, err := di.InitializeApp(cfg)
	if err != nil {
		return err
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.New(getCORSConfig(app.Config)))
	router.Static("/docs", "./docs")

	apiRouter := router.Group("/api/v1")
	protectedApiRouter := apiRouter.Group("/")
	protectedApiRouter.Use(app.AuthMiddleware.Middleware())

	apiRouter.GET("/metrics", app.PrometheusHandler)

	apiRouter.GET("/swagger/*any", app.SwaggerHandler)

	usersHandlerRouterGroup := apiRouter.Group("/user")
	{
		usersHandlerRouterGroup.POST("/create", app.UsersHandler.HandleCreateUser)
		usersHandlerRouterGroup.POST("/update", app.UsersHandler.HandleUpdateUser)
		usersHandlerRouterGroup.DELETE("/delete", app.UsersHandler.HandleDeleteUser)
	}

	authHandlerRouterGroup := apiRouter.Group("/auth")
	{
		authHandlerRouterGroup.POST("/login", app.UsersHandler.LoginUserHandler)
	}

	protectedUsersRouterGroup := protectedApiRouter.Group("/user")
	{
		protectedUsersRouterGroup.GET("/me", app.UsersHandler.HandleGetCurrentUser)
	}

	// Kanban Routes
	kanbanHandlerRouterGroup := protectedApiRouter.Group("/kanban")
	{
		kanbanHandlerRouterGroup.GET("/project/:project_id", app.KanbanHandler.HandleGetKanbansByProjectID)
		kanbanHandlerRouterGroup.POST("/", app.KanbanHandler.HandleCreateKanban)
		kanbanHandlerRouterGroup.GET("/:id", app.KanbanHandler.HandleGetKanbanByID)
		kanbanHandlerRouterGroup.PUT("/:id", app.KanbanHandler.HandleUpdateKanban)
		kanbanHandlerRouterGroup.DELETE("/:id", app.KanbanHandler.HandleDeleteKanban)
	}

	// Task Routes
	taskHandlerRouterGroup := protectedApiRouter.Group("/task")
	{
		taskHandlerRouterGroup.POST("/", app.TaskHandler.HandleCreateTask)
		taskHandlerRouterGroup.GET("/:id", app.TaskHandler.HandleGetTaskByID)
		taskHandlerRouterGroup.PUT("/:id", app.TaskHandler.HandleUpdateTask)
		taskHandlerRouterGroup.DELETE("/:id", app.TaskHandler.HandleDeleteTask)
		taskHandlerRouterGroup.POST("/:task_id/assign", app.TaskHandler.HandleAssignUserToTask)
	}
	apiRouter.GET("/kanban_tasks/:kanban_id", app.TaskHandler.HandleGetTasksByKanbanID)
	apiRouter.GET("/user/:user_id/tasks", app.TaskHandler.HandleGetTasksByUserID)
	apiRouter.GET("/project_tasks/:project_id", app.TaskHandler.HandleGetTasksByProjectID)

	// Comment Routes
	commentHandlerRouterGroup := protectedApiRouter.Group("/comment")
	{
		commentHandlerRouterGroup.POST("/forTask", app.CommentHandler.HandleCreateCommentForTask)
		commentHandlerRouterGroup.GET("/forTask/:task_id", app.CommentHandler.HandleGetCommentsByTaskID)
	}

	// Project Routes
	projectHandlerRouterGroup := protectedApiRouter.Group("/project")
	{
		projectHandlerRouterGroup.POST("/", app.ProjectHandler.HandleCreateProject)
		projectHandlerRouterGroup.GET("/:id", app.ProjectHandler.HandleGetProjectByID)
		projectHandlerRouterGroup.PUT("/:id", app.ProjectHandler.HandleUpdateProject)
		projectHandlerRouterGroup.DELETE("/:id", app.ProjectHandler.HandleDeleteProject)
	}

	projectUsersHandlerRouterGroup := protectedApiRouter.Group("/project_users")
	{
		projectUsersHandlerRouterGroup.POST("/:project_id/invite", app.ProjectHandler.HandleInviteUserToProject)
		projectUsersHandlerRouterGroup.GET("/:project_id/permissions/:user_id", app.ProjectHandler.HandleGetUserPermissionsForProject) //  !!!  ИЗМЕНЕНО ПОРЯДОК СЕГМЕНТОВ !!!
		projectUsersHandlerRouterGroup.GET("/:project_id", app.ProjectHandler.HandleGetUsersInProject)
		projectUsersHandlerRouterGroup.POST("/:project_id/accept", app.ProjectHandler.HandleAcceptProjectInvitation)
	}

	apiRouter.GET("/user_projects/:owner_id", app.ProjectHandler.HandleGetProjectsByOwnerID)
	apiRouter.GET("/user_shared_projects/:owner_id", app.ProjectHandler.HandleGetSharedProjectsByOwnerID)
	apiRouter.GET("/project_subprojects/:parent_project_id", app.ProjectHandler.HandleGetSubprojects)

	routerDSN := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)

	logger.Log.WithFields(log.Fields{
		"host": cfg.HTTP.Host,
		"port": cfg.HTTP.Port,
	}).Info("running app")

	return router.Run(routerDSN)
}

func getCORSConfig(cfg *config.Config) cors.Config {
	return cors.Config{
		AllowOrigins:     cfg.HTTP.AllowOrigins,
		AllowMethods:     cfg.HTTP.AllowMethods,
		AllowHeaders:     cfg.HTTP.AllowHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
