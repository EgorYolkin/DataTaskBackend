package kanban_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/kanban_usecase"
	"DataTask/pkg/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type KanbanHandler struct {
	useCase kanban_usecase.KanbanUseCase
}

func NewKanbanHandler(useCase kanban_usecase.KanbanUseCase) *KanbanHandler {
	return &KanbanHandler{useCase: useCase}
}

// HandleCreateKanban
// @Summary Create Kanban board
// @Description Create a new Kanban board
// @Tags Kanban
// @Accept json
// @Produce json
// @Param request body dto.Kanban true "Kanban board data"
// @Success 201 {object} response.JSONResponse{data=dto.Kanban}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban [post]
func (h *KanbanHandler) HandleCreateKanban(ctx *gin.Context) {
	var kanban dto.Kanban
	if err := ctx.ShouldBindJSON(&kanban); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
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

// HandleUpdateKanban
// @Summary Update Kanban board
// @Description Update a Kanban board's details
// @Tags Kanban
// @Accept json
// @Produce json
// @Param id path int true "Kanban board ID"
// @Param request body dto.Kanban true "Updated Kanban board data"
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

	var kanban dto.Kanban
	if err := ctx.ShouldBindJSON(&kanban); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}
	kanban.ID = id // Ensure ID from path is used

	updatedKanban, err := h.useCase.UpdateKanban(ctx, &kanban)
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
