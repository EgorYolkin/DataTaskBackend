package task_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/task_usecase"
	"DataTask/pkg/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	useCase task_usecase.TaskUseCase
}

func NewTaskHandler(useCase task_usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{useCase: useCase}
}

// CreateTaskRequestParam for request parameters to avoid binding unwanted fields
type CreateTaskRequestParam struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
	// Add other necessary fields for creation here, e.g., ProjectID, KanbanID, etc.
	// Assuming these might come from the route or other means if not in the body
	KanbanID int `json:"kanban_id" binding:"required"` // Example: assuming kanban_id is required in body for task creation
}

type UpdateTaskRequestParam struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted *bool  `json:"is_completed"` // Use pointer to distinguish between false and not provided
}

// HandleCreateTask
// @Summary Create Task
// @Description Create a new Task
// @Tags Task
// @Accept json
// @Produce json
// @Param request body CreateTaskRequestParam true "Task data for creation"
// @Success 201 {object} response.JSONResponse{data=dto.Task}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /task [post]
func (h *TaskHandler) HandleCreateTask(ctx *gin.Context) {
	var param CreateTaskRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	// Map request param data to dto.Task for use case
	task := dto.Task{
		Title:       param.Title,
		Description: param.Description,
		IsCompleted: param.IsCompleted,
		KanbanID:    param.KanbanID,
	}

	createdTask, err := h.useCase.CreateTask(ctx, &task)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, createdTask, "")
}

// HandleGetTaskByID
// @Summary Get Task by ID
// @Description Get a Task by its ID
// @Tags Task
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} response.JSONResponse{data=dto.Task}
// @Failure 400 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /task/{id} [get]
func (h *TaskHandler) HandleGetTaskByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Task ID")
		return
	}

	task, err := h.useCase.GetTaskByID(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	if task == nil {
		response.JSON(ctx, http.StatusNotFound, false, nil, "Task not found")
		return
	}

	response.JSON(ctx, http.StatusOK, true, task, "")
}

// HandleUpdateTask
// @Summary Update Task
// @Description Update a Task's details
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body UpdateTaskRequestParam true "Updated Task data"
// @Success 200 {object} response.JSONResponse{data=dto.Task}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /task/{id} [put]
func (h *TaskHandler) HandleUpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Task ID")
		return
	}

	var param UpdateTaskRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	// Map request param data to a struct for use case
	// Only include fields that were potentially provided
	updateData := dto.Task{ID: id} // Start with the ID from the path

	// Check if fields were provided in the request body and update the struct
	// This logic can be more sophisticated if needed (e.g., distinguishing empty string vs not provided)
	if param.Title != "" { // Assuming empty string means not provided or cleared
		updateData.Title = param.Title
	}
	if param.Description != "" { // Assuming empty string means not provided or cleared
		updateData.Description = param.Description
	}
	if param.IsCompleted != nil { // Check if the pointer is not nil
		updateData.IsCompleted = *param.IsCompleted
	}

	// The use case should handle partial updates based on the provided fields
	updatedTask, err := h.useCase.UpdateTask(ctx, &updateData)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, updatedTask, "")
}

// HandleDeleteTask
// @Summary Delete Task
// @Description Delete a Task by its ID
// @Tags Task
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /task/{id} [delete]
func (h *TaskHandler) HandleDeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Task ID")
		return
	}

	err = h.useCase.DeleteTask(ctx, id)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent) // 204 No Content for successful deletion
}

// HandleGetTasksByKanbanID
// @Summary Get Tasks by Kanban ID
// @Description Get all tasks associated with a Kanban board
// @Tags Task
// @Produce json
// @Param kanban_id path int true "Kanban Board ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Task}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /kanban_tasks/{kanban_id} [get]
func (h *TaskHandler) HandleGetTasksByKanbanID(ctx *gin.Context) {
	kanbanIDStr := ctx.Param("kanban_id")
	kanbanID, err := strconv.Atoi(kanbanIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Kanban ID")
		return
	}

	tasks, err := h.useCase.GetTasksByKanbanID(ctx, kanbanID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, tasks, "")
}

// HandleGetTasksByUserID
// @Summary Get Tasks by User ID
// @Description Get all tasks assigned to a user
// @Tags Task
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Task}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/{user_id}/tasks [get]
func (h *TaskHandler) HandleGetTasksByUserID(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid User ID")
		return
	}

	tasks, err := h.useCase.GetTasksByUserID(ctx, userID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, tasks, "")
}

// HandleAssignUserToTask
// @Summary Assign User to Task
// @Description Assign a user to a task
// @Tags Task
// @Accept json
// @Produce json
// @Param task_id path int true "Task ID"
// @Param request body map[string]int true "User ID to assign"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /task/{task_id}/assign [post]
func (h *TaskHandler) HandleAssignUserToTask(ctx *gin.Context) {
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Task ID")
		return
	}

	var requestBody map[string]int
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	userID, ok := requestBody["user_id"]
	if !ok {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Missing user_id in request")
		return
	}

	err = h.useCase.AssignUserToTask(ctx, taskID, userID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, nil, "User assigned to task")
}

// HandleGetTasksByProjectID
//
//	@Summary Get Tasks by Project ID
//	@Description Get all tasks associated with a project
//	@Tags Task
//	@Produce json
//	@Param project_id path int true "Project ID"
//	@Success 200 {object} response.JSONResponse{data=[]dto.Task}
//	@Failure 400 {object} response.JSONResponse
//	@Failure 500 {object} response.JSONResponse
//	@Router /project_tasks/{project_id} [get]
func (h *TaskHandler) HandleGetTasksByProjectID(ctx *gin.Context) {
	projectIDStr := ctx.Param("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid Project ID")
		return
	}

	tasks, err := h.useCase.GetTasksByProjectID(ctx, projectID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, tasks, "")
}
