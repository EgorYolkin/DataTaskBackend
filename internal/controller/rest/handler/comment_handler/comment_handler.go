package comment_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/comment_usecase"
	"DataTask/internal/usecase/notification_usecase"
	"DataTask/pkg/http/response"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CommentHandler is base handler struct
type CommentHandler struct {
	useCase             comment_usecase.CommentUseCase
	notificationUseCase notification_usecase.NotificationUseCase
}

// NewTaskHandler is base function of handler creation
func NewCommentHandler(useCase comment_usecase.CommentUseCase, notificationUseCase notification_usecase.NotificationUseCase) *CommentHandler {
	return &CommentHandler{useCase: useCase, notificationUseCase: notificationUseCase}
}

type HandleCreateCommentForTaskRequestParam struct {
	Text   string `json:"text" binding:"required"`
	TaskID int    `json:"task_id" binding:"required"`
}

// HandleCreateCommentForTask
// @Summary Create Comment For Task
// @Description Create a new comment for task
// @Tags Comment
// @Accept json
// @Produce json
// @Param request body HandleCreateCommentForTaskRequestParam true "Comment data for creation"
// @Success 201 {object} response.JSONResponse{data=dto.Comment}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /comment/forTask [post]
func (h *CommentHandler) HandleCreateCommentForTask(ctx *gin.Context) {
	var param HandleCreateCommentForTaskRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	authUserID, exists := ctx.Get("user_id")
	if !exists {
		response.JSON(ctx, http.StatusUnauthorized, false, nil, "user_id not found in context")
		return
	}

	userID, ok := authUserID.(int)
	if !ok {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, "invalid user_id in context")
		return
	}

	commentAuthor := dto.User{
		ID: userID,
	}
	comment := dto.Comment{
		Text:   param.Text,
		Author: &commentAuthor,
	}

	createdComment, err := h.useCase.CreateComment(ctx, &comment, param.TaskID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	authUserEmail := ctx.GetString("user_email")
	notification := dto.Notification{
		OwnerID:     createdComment.Author.ID,
		Title:       "New comment",
		Description: fmt.Sprintf(`%s: %s`, authUserEmail, createdComment.Text),
	}
	err = h.notificationUseCase.CreateNotification(ctx, &notification)
	if err != nil {
		fmt.Print(err)
	}

	response.JSON(ctx, http.StatusCreated, true, createdComment, "")
}

// HandleGetCommentsByTaskID
// @Summary Get Comments By Task ID
// @Description Get all comments for a specific task
// @Tags Comment
// @Accept json
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 200 {object} response.JSONResponse{data=[]dto.Comment}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /comment/forTask/{task_id} [get]
func (h *CommentHandler) HandleGetCommentsByTaskID(ctx *gin.Context) {
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "invalid task_id")
		return
	}

	comments, err := h.useCase.GetCommentsByTaskID(ctx, taskID)
	if err != nil {
		response.JSON(ctx, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusOK, true, comments, "")
}
