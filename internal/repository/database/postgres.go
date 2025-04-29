package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	UsersTable        = "users"
	KanbanTable       = "kanban"
	ProjectTasksTable = "project_tasks"
	TaskTable         = "task"
	ProjectsTable     = "projects"
	TaskUsersTable    = "task_users"
	ProjectUsersTable = "project_users"
	KanbanTasksTable  = "kanban_tasks"
)

func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
