package project_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/project_usecase"
	"DataTask/pkg/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProjectHandler struct {
	useCase project_usecase.ProjectUseCase
}

func NewProjectHandler(useCase project_usecase.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{useCase: useCase}
}

type HandleCreateProjectParam struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Color           string `json:"color"`
	ParentProjectID *int   `json:"parent_project_id,omitempty"`
}

// HandleCreateProject
// @Summary Create Project
// @Description Create a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param request body HandleCreateProjectParam true "Project data"
// @Success 201 {object} response.JSONResponse{data=dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project [post]
func (h *ProjectHandler) HandleCreateProject(ctx *gin.Context) {
	var param HandleCreateProjectParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	authUserID := ctx.GetInt("user_id")

	project := dto.Project{
		Name:            param.Name,
		Description:     param.Description,
		Color:           param.Color,
		ParentProjectID: param.ParentProjectID,
		OwnerID:         authUserID,
	}

	createdProject, err := h.useCase.CreateProject(ctx, &project) // UseCase должен вернуть entity
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	// Преобразуем entity в dto, чтобы вернуть клиенту, и устанавливаем CreatedAt
	responseProject := dto.Project{
		ID:              createdProject.ID,
		OwnerID:         createdProject.OwnerID,
		Name:            createdProject.Name,
		Description:     createdProject.Description,
		Color:           createdProject.Color,
		ParentProjectID: createdProject.ParentProjectID,
		CreatedAt:       createdProject.CreatedAt,
		UpdatedAt:       createdProject.UpdatedAt,
	}

	response.JSON(ctx, http.StatusCreated, true, responseProject, "")
}

// HandleGetProjectByID
// @Summary Get Project by ID
// @Description Get a project by its ID
// @Tags Project
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} response.JSONResponse{data=dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project/{id} [get]
func (h *ProjectHandler) HandleGetProjectByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	project, err := h.useCase.GetProjectByID(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	if project == nil {
		response.JSON(ctx, http.StatusNotFound, false, nil, "Project not found")
		return
	}

	response.JSON(ctx, http.StatusOK, true, project, "")
}

// HandleUpdateProject
// @Summary Update Project
// @Description Update a project's details
// @Tags Project
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param request body dto.Project true "Updated project data"
// @Success 200 {object} response.JSONResponse{data=dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project/{id} [put]
func (h *ProjectHandler) HandleUpdateProject(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	var project dto.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}
	project.ID = id // Ensure ID from path is used

	updatedProject, err := h.useCase.UpdateProject(ctx, &project)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, updatedProject, "")
}

// HandleDeleteProject
// @Summary Delete Project
// @Description Delete a project by its ID
// @Tags Project
// @Produce json
// @Param id path int true "Project ID"
// @Success 204 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project/{id} [delete]
func (h *ProjectHandler) HandleDeleteProject(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	err = h.useCase.DeleteProject(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent) // 204 No Content for successful deletion
}

// HandleGetProjectsByOwnerID
// @Summary Get Projects by Owner ID
// @Description Get all projects owned by a user
// @Tags Project
// @Produce json
// @Param owner_id path int true "Owner User ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user_projects/{owner_id} [get]
func (h *ProjectHandler) HandleGetProjectsByOwnerID(ctx *gin.Context) {
	ownerIDStr := ctx.Param("owner_id")
	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Owner User ID")
		return
	}

	projects, err := h.useCase.GetProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, projects, "")
}

// HandleGetSharedProjectsByOwnerID
// @Summary Get Projects by Owner ID
// @Description Get all projects owned by a user
// @Tags Project
// @Produce json
// @Param owner_id path int true "Owner User ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user_shared_projects/{owner_id} [get]
func (h *ProjectHandler) HandleGetSharedProjectsByOwnerID(ctx *gin.Context) {
	ownerIDStr := ctx.Param("owner_id")
	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Owner User ID")
		return
	}

	projects, err := h.useCase.GetSharedProjectsByOwnerID(ctx, ownerID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, projects, "")
}

// HandleGetSubprojects
// @Summary Get Subprojects
// @Description Get all subprojects of a project
// @Tags Project
// @Produce json
// @Param parent_project_id path int true "Parent Project ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Project}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project_subprojects/{parent_project_id} [get]
func (h *ProjectHandler) HandleGetSubprojects(ctx *gin.Context) {
	parentProjectIDStr := ctx.Param("parent_project_id")
	parentProjectID, err := strconv.Atoi(parentProjectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Parent Project ID")
		return
	}

	subprojects, err := h.useCase.GetSubprojects(ctx, parentProjectID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, subprojects, "")
}

// HandleInviteUserToProject
// @Summary Invite User to Project
// @Description Invite a user to a project with specific permissions
// @Tags Project
// @Accept json
// @Produce json
// @Param request body dto.ProjectUserInvite true "User invitation details"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project_users/{project_id}/invite [post]
func (h *ProjectHandler) HandleInviteUserToProject(ctx *gin.Context) {
	var invite dto.ProjectUserInvite
	if err := ctx.ShouldBindJSON(&invite); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	authUserID := ctx.GetInt("user_id")

	err := h.useCase.InviteUserToProject(ctx, &invite, authUserID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, nil, "User invited to project")
}

// HandleGetUserPermissionsForProject
// @Summary Get User Permissions for Project
// @Description Get the permissions of a user in a project
// @Tags Project
// @Produce json
// @Param project_id path int true "Project ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} response.JSONResponse{data=string}
// @Failure 400 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project_users/{project_id}/permissions/{user_id} [get]
func (h *ProjectHandler) HandleGetUserPermissionsForProject(ctx *gin.Context) {
	projectIDStr := ctx.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid User ID")
		return
	}

	permission, err := h.useCase.GetUserPermissionsForProject(ctx, projectID, userID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	if permission == "" {
		response.JSON(ctx, http.StatusNotFound, false, nil, "User is not in the project")
		return
	}

	response.JSON(ctx, http.StatusOK, true, permission, "")
}

// HandleGetUsersInProject
// @Summary Get Users in Project
// @Description Get all users and their permissions in a project
// @Tags Project
// @Produce json
// @Param project_id path int true "Project ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.ProjectUser}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project_users/{project_id} [get]
func (h *ProjectHandler) HandleGetUsersInProject(ctx *gin.Context) {
	projectIDStr := ctx.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	users, err := h.useCase.GetUsersInProject(ctx, projectID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, users, "")
}

// HandleAcceptProjectInvitation
// @Summary Accept Project Invitation
// @Description Accept an invitation to join a project
// @Tags Project
// @Produce json
// @Param project_id path int true "Project ID"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /project_users/{project_id}/accept [post]
func (h *ProjectHandler) HandleAcceptProjectInvitation(ctx *gin.Context) {
	projectIDStr := ctx.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	// Assuming you have middleware to get the current user's ID
	// (e.g., from a JWT). Replace with your actual logic.
	authUserID := ctx.GetInt("user_id")

	err = h.useCase.AcceptProjectInvitation(ctx, projectID, authUserID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, nil, "Project invitation accepted")
}
