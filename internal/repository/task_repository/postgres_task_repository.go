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
        INSERT INTO %s (title, description, is_completed, kanban_id, owner_id) VALUES ($1, $2, $3, $4, $5) 
        RETURNING id, title, description, is_completed, created_at, updated_at, kanban_id;
    `, database.TaskTable) // Define TaskTable in your database package

	err := r.db.QueryRowContext(ctx, q, task.Title, task.Description, task.IsCompleted, task.KanbanID, task.OwnerID).Scan(
		&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt, &task.KanbanID,
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

	task := new(entity.Task) // Initialize task as a pointer

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
	fmt.Println("GetTasksByKanbanID", kanbanID)
	q := fmt.Sprintf(`
        SELECT id, title, description, is_completed, created_at, updated_at, kanban_id
        FROM %s 
        WHERE kanban_id = $1;
    `, database.TaskTable)

	rows, err := r.db.QueryContext(ctx, q, kanbanID)
	if err != nil {
		return nil, fmt.Errorf("get tasks by kanban id: %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t entity.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt, &t.KanbanID); err != nil {
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
	q := `
        SELECT DISTINCT t.id, t.title, t.description, t.is_completed, t.created_at, t.updated_at
        FROM task t
        LEFT JOIN kanban k ON t.kanban_id = k.id
        LEFT JOIN projects p ON k.project_id = p.id
        -- Присоединяем project_users, чтобы учесть приглашенных пользователей
        LEFT JOIN project_users pu ON p.id = pu.project_id
        WHERE
            -- Задачи из проектов, где пользователь - владелец
            p.owner_id = $1
            -- ИЛИ задачи из проектов, куда пользователь приглашен
            OR pu.user_id = $1
            -- ИЛИ задачи, которые не привязаны к канбану/проекту (если такие задачи существуют и должны быть видны)
            OR t.kanban_id IS NULL;
    `

	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, fmt.Errorf("get tasks by user id: %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t entity.Task
		// Убедитесь, что все поля задачи, которые вы ожидаете, сканируются
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
