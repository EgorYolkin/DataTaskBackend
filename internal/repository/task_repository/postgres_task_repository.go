package task_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"fmt"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	q := fmt.Sprintf(`
        INSERT INTO %s (title, description, is_completed) VALUES ($1, $2, $3) 
        RETURNING id, title, description, is_completed, created_at, updated_at;
    `, database.TaskTable) // Define TaskTable in your database package

	err := r.db.QueryRowContext(ctx, q, task.Title, task.Description, task.IsCompleted).Scan(
		&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	return task, nil
}

func (r *PostgresTaskRepository) GetTaskByID(ctx context.Context, id int) (*entity.Task, error) {
	q := fmt.Sprintf(`
        SELECT id, title, description, is_completed, created_at, updated_at FROM %s WHERE id = $1;
    `, database.TaskTable)

	row := r.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get task by id: %w", err)
	}
	return task, nil
}

func (r *PostgresTaskRepository) UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	q := fmt.Sprintf(`
        UPDATE %s SET title = $1, description = $2, is_completed = $3, updated_at = NOW() WHERE id = $4
        RETURNING id, title, description, is_completed, created_at, updated_at;
    `, database.TaskTable)

	err := r.db.QueryRowContext(ctx, q, task.Title, task.Description, task.IsCompleted, task.ID).Scan(
		&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}
	return task, nil
}

func (r *PostgresTaskRepository) DeleteTask(ctx context.Context, id int) error {
	q := fmt.Sprintf(`
        DELETE FROM %s WHERE id = $1;
    `, database.TaskTable)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	return nil
}

func (r *PostgresTaskRepository) GetTasksByKanbanID(ctx context.Context, kanbanID int) ([]*entity.Task, error) {
	q := fmt.Sprintf(`
        SELECT t.id, t.title, t.description, t.is_completed, t.created_at, t.updated_at
        FROM %s kt
        JOIN %s t ON kt.task_id = t.id
        WHERE kt.kanban_id = $1;
    `, database.KanbanTasksTable, database.TaskTable) // Define KanbanTasksTable and TaskTable

	rows, err := r.db.QueryContext(ctx, q, kanbanID)
	if err != nil {
		return nil, fmt.Errorf("get tasks by kanban id: %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t entity.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) GetTasksByUserID(ctx context.Context, userID int) ([]*entity.Task, error) {
	q := fmt.Sprintf(`
        SELECT t.id, t.title, t.description, t.is_completed, t.created_at, t.updated_at
        FROM %s tu
        JOIN %s t ON tu.task_id = t.id
        WHERE tu.user_id = $1;
    `, database.TaskUsersTable, database.TaskTable) // Define TaskUsersTable and TaskTable

	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, fmt.Errorf("get tasks by user id: %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t entity.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) AssignUserToTask(ctx context.Context, taskID int, userID int) error {
	q := fmt.Sprintf(`
        INSERT INTO %s (task_id, user_id) VALUES ($1, $2);
    `, database.TaskUsersTable) // Define TaskUsersTable

	_, err := r.db.ExecContext(ctx, q, taskID, userID)
	if err != nil {
		return fmt.Errorf("assign user to task: %w", err)
	}
	return nil
}

func (r *PostgresTaskRepository) GetTasksByProjectID(ctx context.Context, projectID int) ([]*entity.Task, error) {
	q := fmt.Sprintf(`
        SELECT t.id, t.title, t.description, t.is_completed, t.created_at, t.updated_at
        FROM %s pt
        JOIN %s t ON pt.task_id = t.id
        WHERE pt.project_id = $1;
    `, database.ProjectTasksTable, database.TaskTable) // Define ProjectTasksTable

	rows, err := r.db.QueryContext(ctx, q, projectID)
	if err != nil {
		return nil, fmt.Errorf("get tasks by project id: %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t entity.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return tasks, nil
}
