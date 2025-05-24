package project_usecase

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/domain/entity"
	"DataTask/internal/repository/project_repository"
	"context"
	"fmt"
)

type ProjectUseCaseImpl struct {
	repo project_repository.ProjectRepository
}

func NewProjectUseCase(repo project_repository.ProjectRepository) *ProjectUseCaseImpl {
	return &ProjectUseCaseImpl{repo: repo}
}

func (uc *ProjectUseCaseImpl) CreateProject(ctx context.Context, project *dto.Project) (*dto.Project, error) {
	entityProject := &entity.Project{
		OwnerID:         project.OwnerID,
		Name:            project.Name,
		Description:     project.Description,
		Color:           project.Color,
		ParentProjectID: project.ParentProjectID,
	}

	createdProject, err := uc.repo.CreateProject(ctx, entityProject)
	if err != nil {
		return nil, err
	}

	dtoProject := &dto.Project{
		ID:              createdProject.ID,
		OwnerID:         createdProject.OwnerID,
		Name:            createdProject.Name,
		Description:     createdProject.Description,
		Color:           createdProject.Color,
		ParentProjectID: createdProject.ParentProjectID,
		CreatedAt:       createdProject.CreatedAt,
		UpdatedAt:       createdProject.UpdatedAt,
	}

	return dtoProject, nil
}

func (uc *ProjectUseCaseImpl) GetProjectByID(ctx context.Context, id int) (*dto.Project, error) {
	project, err := uc.repo.GetProjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dtoProject := &dto.Project{
		ID:              project.ID,
		OwnerID:         project.OwnerID,
		Name:            project.Name,
		Description:     project.Description,
		Color:           project.Color,
		ParentProjectID: project.ParentProjectID,
		CreatedAt:       project.CreatedAt,
		UpdatedAt:       project.UpdatedAt,
	}

	return dtoProject, nil
}

func (uc *ProjectUseCaseImpl) UpdateProject(ctx context.Context, project *dto.Project) (*dto.Project, error) {
	entityProject := &entity.Project{
		ID:              project.ID,
		Name:            project.Name,
		Description:     project.Description,
		Color:           project.Color,
		ParentProjectID: project.ParentProjectID,
	}

	updatedProject, err := uc.repo.UpdateProject(ctx, entityProject)
	if err != nil {
		return nil, err
	}

	dtoProject := &dto.Project{
		ID:              updatedProject.ID,
		OwnerID:         updatedProject.OwnerID,
		Name:            updatedProject.Name,
		Description:     updatedProject.Description,
		Color:           updatedProject.Color,
		ParentProjectID: updatedProject.ParentProjectID,
		CreatedAt:       updatedProject.CreatedAt,
		UpdatedAt:       updatedProject.UpdatedAt,
	}

	return dtoProject, nil
}

func (uc *ProjectUseCaseImpl) DeleteProject(ctx context.Context, id int) error {
	project, err := uc.repo.GetProjectByID(ctx, id)
	if err != nil {
		return err
	}

	if project.OwnerID != ctx.Value("user_id").(int) {
		return fmt.Errorf("delete project: user is not the creator of project %d", id)
	}

	err = uc.repo.DeleteProject(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ProjectUseCaseImpl) GetProjectsByOwnerID(ctx context.Context, ownerID int) ([]*dto.Project, error) {
	projects, err := uc.repo.GetProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	var dtoProjects []*dto.Project
	for _, p := range projects {
		dtoProjects = append(dtoProjects, &dto.Project{
			ID:              p.ID,
			OwnerID:         p.OwnerID,
			Name:            p.Name,
			Description:     p.Description,
			Color:           p.Color,
			ParentProjectID: p.ParentProjectID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}

	return dtoProjects, nil
}

func (uc *ProjectUseCaseImpl) GetSharedProjectsByOwnerID(ctx context.Context, ownerID int) ([]*dto.Project, error) {
	projects, err := uc.repo.GetSharedProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	var dtoProjects []*dto.Project
	for _, p := range projects {
		dtoProjects = append(dtoProjects, &dto.Project{
			ID:              p.ID,
			OwnerID:         p.OwnerID,
			Name:            p.Name,
			Description:     p.Description,
			Color:           p.Color,
			ParentProjectID: p.ParentProjectID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}

	return dtoProjects, nil
}

func (uc *ProjectUseCaseImpl) GetSubprojects(ctx context.Context, parentProjectID int) ([]*dto.Project, error) {
	projects, err := uc.repo.GetSubprojects(ctx, parentProjectID)
	if err != nil {
		return nil, err
	}

	var dtoProjects []*dto.Project
	for _, p := range projects {
		dtoProjects = append(dtoProjects, &dto.Project{
			ID:              p.ID,
			OwnerID:         p.OwnerID,
			Name:            p.Name,
			Description:     p.Description,
			Color:           p.Color,
			ParentProjectID: p.ParentProjectID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}

	return dtoProjects, nil
}

/*
func (uc *ProjectUseCaseImpl) GetProjectsByTaskID(ctx context.Context, taskID int) ([]*dto.Project, error) {
	projects, err := uc.repo.GetProjectsByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	var dtoProjects []*dto.Project
	for _, p := range projects {
		dtoProjects = append(dtoProjects, &dto.Project{
			ID:              p.ID,
			OwnerID:         p.OwnerID,
			Name:            p.Name,
			Description:     p.Description,
			Color:           p.Color,
			ParentProjectID: p.ParentProjectID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}

	return dtoProjects, nil
}
*/

func (uc *ProjectUseCaseImpl) InviteUserToProject(ctx context.Context, invite *dto.ProjectUserInvite, invitedByUserID int) error {
	entityProjectUser := &entity.ProjectUser{
		ProjectID:       invite.ProjectID,
		UserEmail:       invite.UserEmail,
		Permission:      invite.Permission,
		InvitedByUserID: &invitedByUserID,
	}

	err := uc.repo.InviteUserToProject(ctx, entityProjectUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ProjectUseCaseImpl) GetUserPermissionsForProject(ctx context.Context, projectID int, userID int) (string, error) {
	permission, err := uc.repo.GetUserPermissionsForProject(ctx, projectID, userID)
	if err != nil {
		return "", err
	}
	return permission, nil
}

func (uc *ProjectUseCaseImpl) GetUsersInProject(ctx context.Context, projectID int) ([]*dto.User, error) {
	projectUsers, err := uc.repo.GetUsersInProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	var dtoProjectUsers []*dto.User
	for _, u := range projectUsers {
		dtoProjectUsers = append(dtoProjectUsers, &dto.User{
			ID: u.ID,

			Name:      u.Name,
			Surname:   u.Surname,
			Email:     u.Email,
			AvatarURL: u.AvatarURL,
		})
	}
	return dtoProjectUsers, nil
}

func (uc *ProjectUseCaseImpl) AcceptProjectInvitation(ctx context.Context, projectID int, userID int) error {
	err := uc.repo.AcceptProjectInvitation(ctx, projectID, userID)
	if err != nil {
		return err
	}
	return nil
}
