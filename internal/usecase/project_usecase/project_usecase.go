package project_usecase

import (
	"DataTask/internal/domain/dto"
	"context"
)

type ProjectUseCase interface {
	CreateProject(ctx context.Context, project *dto.Project) (*dto.Project, error)
	GetProjectByID(ctx context.Context, id int) (*dto.Project, error)
	UpdateProject(ctx context.Context, project *dto.Project) (*dto.Project, error)
	DeleteProject(ctx context.Context, id int) error
	GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]*dto.Project, error)
	GetSharedProjectsByOwnerID(ctx context.Context, ownerID int) ([]*dto.Project, error)
	GetSubprojects(ctx context.Context, parentProjectID int) ([]*dto.Project, error)
	//GetProjectsByTaskID(ctx context.Context, taskID int) ([]*dto.Project, error) // If needed

	InviteUserToProject(ctx context.Context, invite *dto.ProjectUserInvite, invitedByUserID int) error
	GetUserPermissionsForProject(ctx context.Context, projectID int, userID int) (string, error)
	GetUsersInProject(ctx context.Context, projectID int) ([]*dto.User, error)
	AcceptProjectInvitation(ctx context.Context, projectID int, userID int) error
}
