package di

import (
	"DataTask/internal/config"
	"DataTask/internal/controller/rest/handler/comment_handler"
	"DataTask/internal/controller/rest/handler/kanban_handler"
	"DataTask/internal/controller/rest/handler/notification_handler"
	project_handler "DataTask/internal/controller/rest/handler/project_handler"
	"DataTask/internal/controller/rest/handler/task_handler"
	"DataTask/internal/controller/rest/handler/users_handler"
	"DataTask/internal/repository/comment_repository"
	"DataTask/internal/repository/kanban_repository"
	"DataTask/internal/repository/notification_repository"
	"DataTask/internal/repository/project_repository"
	"DataTask/internal/repository/task_repository"
	"DataTask/internal/repository/user_repository"
	"DataTask/internal/usecase/comment_usecase"
	"DataTask/internal/usecase/kanban_usecase"
	"DataTask/internal/usecase/notification_usecase"
	"DataTask/internal/usecase/project_usecase"
	"DataTask/internal/usecase/task_usecase"
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
		"base_path": "/api/v1",
		"version":   cfg.Swagger.Version,
	}).Info("set up swagger")

	return ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName(instanceName),
		ginSwagger.URL("/docs/swagger.json"), //  !!!  ИЗМЕНЕНО !!!
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

func InitializeNotificationHandler(db *sql.DB) *notification_handler.NotificationHandler {
	repo := notification_repository.NewPostgresNotificationRepository(db)
	useCase := notification_usecase.NewNotificationUseCase(repo)
	handler := notification_handler.NewNotificationHandler(useCase)
	return handler
}

func InitializeKanbanHandler(db *sql.DB, notificationUseCase notification_usecase.NotificationUseCase) *kanban_handler.KanbanHandler {
	repo := kanban_repository.NewPostgresKanbanRepository(db)
	useCase := kanban_usecase.NewKanbanUseCase(repo)
	handler := kanban_handler.NewKanbanHandler(useCase, notificationUseCase)
	return handler
}

func InitializeCommentHandler(db *sql.DB, notificationUseCase notification_usecase.NotificationUseCase) *comment_handler.CommentHandler {
	repo := comment_repository.NewPostgresCommentRepository(db)
	useCase := comment_usecase.NewCommentUseCase(repo)
	handler := comment_handler.NewCommentHandler(useCase, notificationUseCase)
	return handler
}

func InitializeTaskHandler(db *sql.DB, notificationUseCase notification_usecase.NotificationUseCase) *task_handler.TaskHandler {
	repo := task_repository.NewPostgresTaskRepository(db)
	useCase := task_usecase.NewTaskUseCase(repo)
	handler := task_handler.NewTaskHandler(useCase, notificationUseCase)
	return handler
}

func InitializeProjectHandler(db *sql.DB, notificationUseCase notification_usecase.NotificationUseCase) *project_handler.ProjectHandler {
	repo := project_repository.NewPostgresProjectRepository(db)
	useCase := project_usecase.NewProjectUseCase(repo)
	handler := project_handler.NewProjectHandler(useCase, notificationUseCase)
	return handler
}
