package comment_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"fmt"
)

type PostgresCommentRepository struct {
	db *sql.DB
}

func NewPostgresCommentRepository(db *sql.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{db: db}
}

func (r *PostgresCommentRepository) CreateCommentForTask(ctx context.Context, comment *entity.Comment, taskID int) (*entity.Comment, error) {
	q := fmt.Sprintf(`
       INSERT INTO %s (author, text)
       VALUES ($1, $2)
       RETURNING id, author, text, created_at, updated_at;
    `, database.CommentTable)

	var authorID int
	err := r.db.QueryRowContext(ctx, q,
		comment.Author.ID,
		comment.Text,
	).Scan(
		&comment.ID,
		&authorID,
		&comment.Text,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create comment: %w", err)
	}

	comment.Author = &entity.User{ID: authorID}

	_, err = r.db.ExecContext(ctx, fmt.Sprintf(`
       INSERT INTO %s (task_id, comment_id)
       VALUES ($1, $2);
    `, database.CommentTaskTable), taskID, comment.ID)
	if err != nil {
		return nil, fmt.Errorf("create comment task link: %w", err)
	}

	return comment, nil
}

func (r *PostgresCommentRepository) GetCommentsByTaskID(ctx context.Context, taskID int) ([]*entity.Comment, error) {
	q := fmt.Sprintf(`
		SELECT
			c.id,
			c.text,
			c.created_at,
			c.updated_at,
			u.id AS author_id,
			u.name AS author_name,
			u.surname AS author_surname,
			u.email AS author_email,
			u.avatar_url AS author_avatar_url,
			u.created_at AS author_created_at,
			u.updated_at AS author_updated_at
		FROM
			%s ct
		JOIN
			%s c ON ct.comment_id = c.id
		JOIN
			%s u ON c.author = u.id
		WHERE
			ct.task_id = $1
		ORDER BY
			c.created_at ASC;
	`, database.CommentTaskTable, database.CommentTable, database.UsersTable)

	rows, err := r.db.QueryContext(ctx, q, taskID)
	if err != nil {
		return nil, fmt.Errorf("query comments by task id: %w", err)
	}
	defer rows.Close()

	comments := make([]*entity.Comment, 0)
	for rows.Next() {
		comment := &entity.Comment{
			Author: &entity.User{},
		}
		err := rows.Scan(
			&comment.ID,
			&comment.Text,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.Author.ID,
			&comment.Author.Name,
			&comment.Author.Surname,
			&comment.Author.Email,
			&comment.Author.AvatarURL,
			&comment.Author.CreatedAt,
			&comment.Author.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan comment: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return comments, nil
}
