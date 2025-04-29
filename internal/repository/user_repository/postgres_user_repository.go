package user_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := fmt.Sprintf(`
		INSERT INTO %s (
		   name, 
		   surname, 
		   email, 
		   hashed_password,
		   salt
		) VALUES (
		   $1, $2, $3, $4, $5
		)
		RETURNING id;
	`, database.UsersTable)

	fmt.Println(q)

	err := r.db.QueryRowContext(ctx, q, user.Name, user.Surname, user.Email, user.HashedPassword, user.Salt).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := fmt.Sprintf(`
		UPDATE %s SET
			name = $1,
			surname = $2,
			email = $3,
			avatar_url = $4,
			hashed_password = $5,
			updated_at = NOW()
		WHERE id = $6
		RETURNING id, name, surname, email, avatar_url, hashed_password, created_at, updated_at;
	`, database.UsersTable)

	row := r.db.QueryRowContext(ctx, q,
		user.Name,
		user.Surname,
		user.Email,
		user.AvatarURL,
		user.HashedPassword,
		user.ID,
	)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.AvatarURL,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id int) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE id = $1;`, database.UsersTable)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	q := fmt.Sprintf(`
		SELECT 
			id, 
			email, 
			hashed_password, 
			salt,
			name, 
			surname, 
			avatar_url, 
			created_at, 
			updated_at
		FROM %s WHERE email = $1;
	`, database.UsersTable)

	row := r.db.QueryRowContext(ctx, q, email)

	var user entity.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.HashedPassword,
		&user.Salt,
		&user.Name,
		&user.Surname,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}

	return &user, nil
}
