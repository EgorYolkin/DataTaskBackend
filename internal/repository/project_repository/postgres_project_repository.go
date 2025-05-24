package project_repository

import (
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type PostgresProjectRepository struct {
	db *sql.DB
}

func NewPostgresProjectRepository(db *sql.DB) *PostgresProjectRepository {
	return &PostgresProjectRepository{db: db}
}

func (r *PostgresProjectRepository) CreateProject(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	q := fmt.Sprintf(`
        INSERT INTO %s (owner_id, name, description, color, parent_project_id) 
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, owner_id, name, description, color, parent_project_id, created_at, updated_at;
    `, database.ProjectsTable) // Define ProjectsTable

	err := r.db.QueryRowContext(ctx, q,
		project.OwnerID, project.Name, project.Description, project.Color, project.ParentProjectID).Scan(
		&project.ID, &project.OwnerID, &project.Name, &project.Description, &project.Color,
		&project.ParentProjectID, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}
	return project, nil
}

func (r *PostgresProjectRepository) GetProjectByID(ctx context.Context, id int) (*entity.Project, error) {
	q := fmt.Sprintf(`
        SELECT id, owner_id, name, description, color, parent_project_id, created_at, updated_at
        FROM %s WHERE id = $1;
    `, database.ProjectsTable)

	row := r.db.QueryRowContext(ctx, q, id)

	project := new(entity.Project) // Initialize project as a pointer

	err := row.Scan(
		&project.ID, &project.OwnerID, &project.Name, &project.Description, &project.Color,
		&project.ParentProjectID, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get project by id: %w", err)
	}
	return project, nil
}

func (r *PostgresProjectRepository) UpdateProject(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	q := fmt.Sprintf(`
        UPDATE %s SET 
            name = $1, description = $2, color = $3, parent_project_id = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING id, owner_id, name, description, color, parent_project_id, created_at, updated_at;
    `, database.ProjectsTable)

	err := r.db.QueryRowContext(ctx, q,
		project.Name, project.Description, project.Color, project.ParentProjectID, project.ID).Scan(
		&project.ID, &project.OwnerID, &project.Name, &project.Description, &project.Color,
		&project.ParentProjectID, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("update project: %w", err)
	}
	return project, nil
}

func (r *PostgresProjectRepository) DeleteProject(ctx context.Context, id int) error {
	q := fmt.Sprintf(`
        DELETE FROM %s WHERE id = $1;
    `, database.ProjectsTable)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete project: %w", err)
	}
	return nil
}

func (r *PostgresProjectRepository) GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]*entity.Project, error) {
	q := fmt.Sprintf(`
        SELECT id, owner_id, name, description, color, parent_project_id, created_at, updated_at
        FROM %s WHERE owner_id = $1;
    `, database.ProjectsTable)

	rows, err := r.db.QueryContext(ctx, q, ownerID)
	if err != nil {
		return nil, fmt.Errorf("get projects by owner id: %w", err)
	}
	defer rows.Close()

	var projects []*entity.Project
	for rows.Next() {
		var p entity.Project
		if err := rows.Scan(
			&p.ID, &p.OwnerID, &p.Name, &p.Description, &p.Color,
			&p.ParentProjectID, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}
		projects = append(projects, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return projects, nil
}

func (r *PostgresProjectRepository) GetSharedProjectsByOwnerID(ctx context.Context, ownerID int) ([]*entity.Project, error) {
	q := fmt.Sprintf(`
        SELECT p.id, p.owner_id, p.name, p.description, p.color, p.parent_project_id, p.created_at, p.updated_at
        FROM %s p
        JOIN %s pu ON p.id = pu.project_id
        WHERE pu.user_id = $1;
    `, database.ProjectsTable, database.ProjectUsersTable)

	rows, err := r.db.QueryContext(ctx, q, ownerID)
	if err != nil {
		return nil, fmt.Errorf("get shared projects by user id: %w", err)
	}
	defer rows.Close()

	var projects []*entity.Project
	for rows.Next() {
		var p entity.Project
		if err := rows.Scan(
			&p.ID, &p.OwnerID, &p.Name, &p.Description, &p.Color,
			&p.ParentProjectID, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}
		projects = append(projects, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return projects, nil
}

func (r *PostgresProjectRepository) GetSubprojects(ctx context.Context, parentProjectID int) ([]*entity.Project, error) {
	q := fmt.Sprintf(`
        SELECT id, owner_id, name, description, color, parent_project_id, created_at, updated_at
        FROM %s WHERE parent_project_id = $1;
    `, database.ProjectsTable)

	rows, err := r.db.QueryContext(ctx, q, parentProjectID)
	if err != nil {
		return nil, fmt.Errorf("get subprojects: %w", err)
	}
	defer rows.Close()

	var projects []*entity.Project
	for rows.Next() {
		var p entity.Project
		if err := rows.Scan(
			&p.ID, &p.OwnerID, &p.Name, &p.Description, &p.Color,
			&p.ParentProjectID, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}
		projects = append(projects, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return projects, nil
}

/*
func (r *PostgresProjectRepository) GetProjectsByTaskID(ctx context.Context, taskID int) ([]*entity.Project, error) {
    q := fmt.Sprintf(`
        SELECT p.id, p.owner_id, p.name, p.description, p.color, p.parent_project_id, p.created_at, p.updated_at
        FROM %s pt
        JOIN %s p ON pt.project_id = p.id
        WHERE pt.task_id = $1;
    `, database.ProjectTasksTable, database.ProjectsTable) // If needed

    rows, err := r.db.QueryContext(ctx, q, taskID)
    if err != nil {
        return nil, fmt.Errorf("get projects by task id: %w", err)
    }
    defer rows.Close()

    var projects []*entity.Project
    for rows.Next() {
        var p entity.Project
        if err := rows.Scan(
            &p.ID, &p.OwnerID, &p.Name, &p.Description, &p.Color,
            &p.ParentProjectID, &p.CreatedAt, &p.UpdatedAt,
        ); err != nil {
            return nil, fmt.Errorf("scan project: %w", err)
        }
        projects = append(projects, &p)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows error: %w", err)
    }
    return projects, nil
}
*/

func (r *PostgresProjectRepository) InviteUserToProject(ctx context.Context, projectUser *entity.ProjectUser) error {
	// 1. Get user_id by userEmail
	var userID int64
	selectUserQuery :=
		fmt.Sprintf(`
        SELECT id FROM %s WHERE email = $1;
    `, database.UsersTable)
	err := r.db.QueryRowContext(ctx, selectUserQuery, projectUser.UserEmail).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invite user to project: user with email %s not found", projectUser.UserEmail)
		}
		return fmt.Errorf("invite user to project: failed to get user ID by email: %w", err)
	}

	// 2. Use getting userID for insert in project_users
	insertQuery := fmt.Sprintf(`
        INSERT INTO %s (project_id, user_id, permission, invited_by_user_id)
        VALUES ($1, $2, $3, $4);
    `, database.ProjectUsersTable)

	_, err = r.db.ExecContext(ctx, insertQuery,
		projectUser.ProjectID, userID, projectUser.Permission, projectUser.InvitedByUserID) // Используем userID здесь
	if err != nil {
		return fmt.Errorf("invite user to project: failed to insert into project users: %w", err)
	}

	return nil
}

func (r *PostgresProjectRepository) GetUserPermissionsForProject(ctx context.Context, projectID int, userID int) (string, error) {
	q := fmt.Sprintf(`
        SELECT permission FROM %s
        WHERE project_id = $1 AND user_id = $2;
    `, database.ProjectUsersTable)

	row := r.db.QueryRowContext(ctx, q, projectID, userID)
	var permission string
	err := row.Scan(&permission)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user is not in project")
		}
		return "", fmt.Errorf("get user permissions: %w", err)
	}
	return permission, nil
}

func (r *PostgresProjectRepository) GetUsersInProject(ctx context.Context, projectID int) ([]*entity.User, error) {
	q := fmt.Sprintf(`
		SELECT u.id, u.name, u.surname, u.email, u.avatar_url, u.created_at, u.updated_at
		FROM %s pu
		JOIN users u ON pu.user_id = u.id
		WHERE pu.project_id = $1;
	`, database.ProjectUsersTable)

	rows, err := r.db.QueryContext(ctx, q, projectID)
	if err != nil {
		return nil, fmt.Errorf("get users in project: %w", err)
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(
			&u.ID, &u.Name, &u.Surname, &u.Email, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}

func (r *PostgresProjectRepository) AcceptProjectInvitation(ctx context.Context, projectID int, userID int) error {
	q := fmt.Sprintf(`
        UPDATE %s SET joined_at = NOW()
        WHERE project_id = $1 AND user_id = $2;
    `, database.ProjectUsersTable)

	_, err := r.db.ExecContext(ctx, q, projectID, userID)
	if err != nil {
		return fmt.Errorf("accept project invitation: %w", err)
	}
	return nil
}
