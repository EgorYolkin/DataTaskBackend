package project_repository

import (
	"DataTask/internal/domain/entity"
	"context"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *entity.Project) (*entity.Project, error)
	GetProjectByID(ctx context.Context, id int) (*entity.Project, error)
	UpdateProject(ctx context.Context, project *entity.Project) (*entity.Project, error)
	DeleteProject(ctx context.Context, id int) error
	GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]*entity.Project, error)
	GetSharedProjectsByOwnerID(ctx context.Context, ownerID int) ([]*entity.Project, error)
	GetSubprojects(ctx context.Context, parentProjectID int) ([]*entity.Project, error)
	//GetProjectsByTaskID(ctx context.Context, taskID int) ([]*entity.Project, error) // If needed

	InviteUserToProject(ctx context.Context, projectUser *entity.ProjectUser) error
	GetUserPermissionsForProject(ctx context.Context, projectID int, userID int) (string, error)
	GetUsersInProject(ctx context.Context, projectID int) ([]*entity.ProjectUser, error)
	AcceptProjectInvitation(ctx context.Context, projectID int, userID int) error
}
