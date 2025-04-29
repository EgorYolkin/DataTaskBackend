package users_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/user_usecase"
	"DataTask/pkg/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandler struct {
	UseCase user_usecase.UserUseCase
}

func NewUsersHandler(useCase user_usecase.UserUseCase) *UsersHandler {
	return &UsersHandler{
		UseCase: useCase,
	}
}

// HandleCreateUser HTTP POST handler to create a user
// @Summary create user
// @Description create user
// @Tags User
// @Param request body dto.User true "User create data"
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse{data=dto.User}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/create [get]
func (h *UsersHandler) HandleCreateUser(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	err := h.UseCase.CreateUser(ctx, &user)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, nil, "")
}

// HandleUpdateUser HTTP POST handler to update a user
// @Summary update user
// @Description update user
// @Tags User
// @Param request body dto.User true "User create data"
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse{data=dto.User}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/update [get]
func (h *UsersHandler) HandleUpdateUser(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	u, err := h.UseCase.UpdateUser(ctx, &user)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, u, "")
}

// HandleDeleteUser HTTP POST handler to delete a user
// @Summary delete user
// @Description delete user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/delete [get]
func (h *UsersHandler) HandleDeleteUser(ctx *gin.Context) {
	uid := ctx.GetInt("user_id")

	err := h.UseCase.DeleteUser(ctx, uid)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, nil, "")
}
