package notification_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/notification_usecase"
	"DataTask/pkg/http/response"
	"fmt"
	"net/http"
	"strconv" // Для преобразования ID из строки в int

	"github.com/gin-gonic/gin"
)

// NotificationHandler is base handler struct
type NotificationHandler struct {
	useCase             notification_usecase.NotificationUseCase
}

// NewNotificationHandler is base function of handler creation
func NewNotificationHandler(useCase notification_usecase.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{useCase: useCase}
}

type HandleCreateNotificationParam struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// @Summary Create a new notification
// @Description Creates a new notification for a user
// @Tags Notifications
// @Accept json
// @Produce json
// @Param notification body HandleCreateNotificationParam true "Notification object to be created"
// @Success 201 {object} map[string]string "Notification created successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /notification [post]
func (h *NotificationHandler) HandleCreateNotification(ctx *gin.Context) {
	var params HandleCreateNotificationParam
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid request payload")
		return
	}

	ownerID := ctx.GetInt("user_id")

	notificationDTO := dto.Notification{
		Title:       params.Title,
		Description: params.Description,
		OwnerID:     ownerID,
	}

	if err := h.useCase.CreateNotification(ctx, &notificationDTO); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Failed to create notification")
		return
	}

	response.JSON(ctx, http.StatusCreated, true, nil, "")
}

// @Summary Get user notifications
// @Description Retrieves all notifications for a specific user
// @Tags Notifications
// @Produce json
// @Success 200 {array} dto.Notification "List of notifications"
// @Failure 400 {object} map[string]string "Invalid owner ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /notification [get]
func (h *NotificationHandler) HandleGetUserNotifications(ctx *gin.Context) {
	ownerID := ctx.GetInt("user_id")

	notifications, err := h.useCase.GetUserNotificationsByID(ctx, ownerID)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Failed to retrieve notifications")
		return
	}

	response.JSON(ctx, http.StatusOK, true, notifications, "")
}

// @Summary Mark notification as read
// @Description Marks a specific notification as read by its ID
// @Tags Notifications
// @Produce json
// @Param notification_id path int true "Notification ID"
// @Success 200 {object} map[string]string "Notification marked as read successfully"
// @Failure 400 {object} map[string]string "Invalid notification ID"
// @Failure 404 {object} map[string]string "Notification not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /notification/{notification_id}/read [patch]
func (h *NotificationHandler) HandleSetNotificationIsRead(ctx *gin.Context) {
	notificationIDStr := ctx.Param("notification_id")
	notificationID, err := strconv.Atoi(notificationIDStr)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Invalid notification ID")
		return
	}

	err = h.useCase.SetNotificationIsReadByID(ctx, notificationID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no notification found with ID %d to update", notificationID) {
			response.JSON(ctx, http.StatusBadRequest, false, nil, "Notification not found")
			return
		}
		response.JSON(ctx, http.StatusBadRequest, false, nil, "Failed to mark notification as read")
		return
	}

	response.JSON(ctx, http.StatusOK, true, nil, "")
}
