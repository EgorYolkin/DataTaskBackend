package kanban_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/kanban_usecase"
	"DataTask/internal/usecase/notification_usecase"
	"DataTask/pkg/http/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// KanbanHandler is base handler struct
type KanbanHandler struct {
	useCase             kanban_usecase.KanbanUseCase
	notificationUseCase notification_usecase.NotificationUseCase
}

// NewKanbanHandler is base function of handler creation
func NewKanbanHandler(useCase kanban_usecase.KanbanUseCase, notificationUseCase notification_usecase.NotificationUseCase) *KanbanHandler {
	return &KanbanHandler{useCase: useCase, notificationUseCase: notificationUseCase}
}

type CreateKanbanRequestParam struct {
	Name      string `json:"name" binding:"required"`
	ProjectID int    `json:"project_id" binding:"required"`
}

type UpdateKanbanRequestParam struct {
	Name string `json:"name"`
}

// HandleCreateKanban
// @Summary Create Kanban board
// @Description Create a new Kanban board
// @Tags Kanban
// @Accept json
// @Produce json
// @Param request body CreateKanbanRequestParam true "Kanban board data for creation"
// @Success 201 {object} response.JSONResponse{data=dto.Kanban}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban [post]
func (h *KanbanHandler) HandleCreateKanban(ctx *gin.Context) {
	var param CreateKanbanRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	// Map request param data to dto.Kanban for use case
	kanban := dto.Kanban{
		Name:      param.Name,
		ProjectID: param.ProjectID,
		// ID, CreatedAt, UpdatedAt will be set by the use case/repository
	}

	createdKanban, err := h.useCase.CreateKanban(ctx, &kanban)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, createdKanban, "")
}

// HandleGetKanbanByID
// @Summary Get Kanban board by ID
// @Description Get a Kanban board by its ID
// @Tags Kanban
// @Produce json
// @Param id path int true "Kanban board ID"
// @Success 200 {object} response.JSONResponse{data=dto.Kanban}
// @Failure 400 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban/{id} [get]
func (h *KanbanHandler) HandleGetKanbanByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Kanban ID")
		return
	}

	kanban, err := h.useCase.GetKanbanByID(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	if kanban == nil {
		response.JSON(ctx, http.StatusNotFound, false, nil, "Kanban board not found")
		return
	}

	response.JSON(ctx, http.StatusOK, true, kanban, "")
}

// HandleGetKanbansByProjectID
// @Summary Get Kanban board by project ID
// @Description Get a Kanban board by project ID
// @Tags Kanban
// @Produce json
// @Param project_id path int true "Kanban board project ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Kanban}
// @Failure 400 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban/project/{id} [get]
func (h *KanbanHandler) HandleGetKanbansByProjectID(ctx *gin.Context) {
	projectIDStr := ctx.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Kanban ID")
		return
	}

	kanban, err := h.useCase.GetKanbansByProjectID(ctx, projectID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	if kanban == nil {
		response.JSON(ctx, http.StatusNotFound, false, nil, "Kanban board not found")
		return
	}

	response.JSON(ctx, http.StatusOK, true, kanban, "")
}

// HandleUpdateKanban
// @Summary Update Kanban board
// @Description Update a Kanban board's details
// @Tags Kanban
// @Accept json
// @Produce json
// @Param id path int true "Kanban board ID"
// @Param request body UpdateKanbanRequestParam true "Updated Kanban board data"
// @Success 200 {object} response.JSONResponse{data=dto.Kanban}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban/{id} [put]
func (h *KanbanHandler) HandleUpdateKanban(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Kanban ID")
		return
	}

	var param UpdateKanbanRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	// Map request param data to a struct for use case
	// Only include fields that were potentially provided
	updateData := dto.Kanban{ID: id} // Start with the ID from the path

	// Check if fields were provided in the request body and update the struct
	// Using zero values to indicate "not provided" is a simple approach.
	if param.Name != "" {
		updateData.Name = param.Name
	}

	// The use case should handle partial updates based on the provided fields
	updatedKanban, err := h.useCase.UpdateKanban(ctx, &updateData)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, updatedKanban, "")
}

// HandleDeleteKanban
// @Summary Delete Kanban board
// @Description Delete a Kanban board by its ID
// @Tags Kanban
// @Produce json
// @Param id path int true "Kanban board ID"
// @Success 204 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban/{id} [delete]
func (h *KanbanHandler) HandleDeleteKanban(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Kanban ID")
		return
	}

	err = h.useCase.DeleteKanban(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent) // 204 No Content for successful deletion
}
