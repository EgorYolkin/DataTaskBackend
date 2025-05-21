package di

import (
	"DataTask/internal/config"
	"DataTask/internal/controller/rest/handler/comment_handler"
	"DataTask/internal/controller/rest/handler/kanban_handler"
	"DataTask/internal/controller/rest/handler/project_handler"
	"DataTask/internal/controller/rest/handler/task_handler"
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

	UsersHandler   *users_handler.UsersHandler
	KanbanHandler  *kanban_handler.KanbanHandler
	TaskHandler    *task_handler.TaskHandler
	ProjectHandler *project_handler.ProjectHandler
	CommentHandler *comment_handler.CommentHandler

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
	kanbanHandler := InitializeKanbanHandler(db)
	taskHandler := InitializeTaskHandler(db)
	projectHandler := InitializeProjectHandler(db)
	commentHandler := InitializeCommentHandler(db)

	authMiddleware := InitializeAuthMiddleware(db, cfg.JWT.Secret)

	return &App{
		Config: cfg,

		PrometheusHandler: prometheusHandler,
		SwaggerHandler:    swaggerHandler,
		UsersHandler:      usersHandler,
		KanbanHandler:     kanbanHandler,
		TaskHandler:       taskHandler,
		ProjectHandler:    projectHandler,
		CommentHandler:    commentHandler,

		AuthMiddleware: authMiddleware,
	}, nil
}
