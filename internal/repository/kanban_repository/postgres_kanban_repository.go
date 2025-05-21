package kanban_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"fmt"
)

type PostgresKanbanRepository struct {
	db *sql.DB
}

func NewPostgresKanbanRepository(db *sql.DB) *PostgresKanbanRepository {
	return &PostgresKanbanRepository{db: db}
}

func (r *PostgresKanbanRepository) CreateKanban(ctx context.Context, kanban *entity.Kanban) (*entity.Kanban, error) {
	q := fmt.Sprintf(`
        INSERT INTO %s (name, project_id) VALUES ($1, $2) RETURNING id, name, created_at, updated_at;
    `, database.KanbanTable) // You'll need to define KanbanTable in your database package

	err := r.db.QueryRowContext(ctx, q, kanban.Name, kanban.ProjectID).Scan(&kanban.ID, &kanban.Name, &kanban.CreatedAt, &kanban.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create kanban: %w", err)
	}
	return kanban, nil
}

func (r *PostgresKanbanRepository) GetKanbanByID(ctx context.Context, id int) (*entity.Kanban, error) {
	q := fmt.Sprintf(`
        SELECT id, name, created_at, updated_at FROM %s WHERE id = $1;
    `, database.KanbanTable)

	row := r.db.QueryRowContext(ctx, q, id)

	kanban := new(entity.Kanban) // Initialize kanban as a pointer

	err := row.Scan(&kanban.ID, &kanban.Name, &kanban.CreatedAt, &kanban.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get kanban by id: %w", err)
	}

	return kanban, nil
}

func (r *PostgresKanbanRepository) GetKanbansByProjectID(ctx context.Context, projectID int) ([]*entity.Kanban, error) {
	q := fmt.Sprintf(`
        SELECT id, name, created_at, updated_at FROM %s WHERE project_id = $1;
    `, database.KanbanTable)

	rows, err := r.db.QueryContext(ctx, q, projectID)
	if err != nil {
		return nil, fmt.Errorf("get all kanbans: %w", err)
	}
	defer rows.Close()

	var kanbans []*entity.Kanban
	for rows.Next() {
		var k entity.Kanban
		if err := rows.Scan(&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan kanban: %w", err)
		}
		kanbans = append(kanbans, &k)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return kanbans, nil
}

func (r *PostgresKanbanRepository) UpdateKanban(ctx context.Context, kanban *entity.Kanban) (*entity.Kanban, error) {
	q := fmt.Sprintf(`
        UPDATE %s SET name = $1, updated_at = NOW() WHERE id = $2 
        RETURNING id, name, created_at, updated_at;
    `, database.KanbanTable)

	err := r.db.QueryRowContext(ctx, q, kanban.Name, kanban.ID).Scan(&kanban.ID, &kanban.Name, &kanban.CreatedAt, &kanban.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("update kanban: %w", err)
	}
	return kanban, nil
}

func (r *PostgresKanbanRepository) DeleteKanban(ctx context.Context, id int) error {
	q := fmt.Sprintf(`
        DELETE FROM %s WHERE id = $1;
    `, database.KanbanTable)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete kanban: %w", err)
	}
	return nil
}

func (r *PostgresKanbanRepository) GetAllKanbans(ctx context.Context) ([]*entity.Kanban, error) {
	q := fmt.Sprintf(`SELECT id, name, created_at, updated_at FROM %s;`, database.KanbanTable)

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("get all kanbans: %w", err)
	}
	defer rows.Close()

	var kanbans []*entity.Kanban
	for rows.Next() {
		var k entity.Kanban
		if err := rows.Scan(&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan kanban: %w", err)
		}
		kanbans = append(kanbans, &k)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return kanbans, nil
}
