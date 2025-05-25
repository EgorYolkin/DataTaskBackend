package users_handler

import (
	"DataTask/internal/domain/dto"
	"DataTask/internal/usecase/user_usecase"
	"DataTask/pkg/hashing"
	"DataTask/pkg/http/response"
	"DataTask/pkg/jwt"
	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// ProjectHandler is base handler struct
type UsersHandler struct {
	useCase      user_usecase.UserUseCase
	jwtSecretKey string
}

// Base JWT TTL`s
var (
	accessTTL  = time.Hour * 24
	refreshTTL = time.Hour * 24 * 7
)

// NewUsersHandler is base function of handler creation
func NewUsersHandler(useCase user_usecase.UserUseCase, jwtSecretKey string) *UsersHandler {
	return &UsersHandler{
		useCase:      useCase,
		jwtSecretKey: jwtSecretKey,
	}
}

type CreateUserRequestParam struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

// HandleCreateUser HTTP POST handler to create a user
// @Summary create user
// @Description create user
// @Tags User
// @Param request body CreateUserRequestParam true "User create data"
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse{data=dto.User}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/create [post]
func (h *UsersHandler) HandleCreateUser(ctx *gin.Context) {
	var param CreateUserRequestParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	user := dto.User{
		Name:     param.Name,
		Surname:  param.Surname,
		Email:    param.Email,
		Password: param.Password,
	}

	err := h.useCase.CreateUser(ctx, &user)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	tokenPair, err := jwt.CreateTokenPair(int(user.ID), user.Email, h.jwtSecretKey, accessTTL, refreshTTL)

	ctx.SetCookie(
		"refresh_token",
		tokenPair.RefreshToken,
		60*60*24,
		"/",
		"",
		false,
		true,
	)

	responseData := map[string]string{
		"access_token": tokenPair.AccessToken,
	}

	response.JSON(ctx, http.StatusOK, true, responseData, "")
	return
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
// @Router /user/update [post]
func (h *UsersHandler) HandleUpdateUser(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	u, err := h.useCase.UpdateUser(ctx, &user)
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
// @Router /user/delete [delete]
func (h *UsersHandler) HandleDeleteUser(ctx *gin.Context) {
	uid := ctx.GetInt("user_id")

	err := h.useCase.DeleteUser(ctx, uid)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	response.JSON(ctx, http.StatusCreated, true, nil, "")
}

type LoginUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginUserHandler handles the HTTP POST request to login user
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginUserParams true "Login user params"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /auth/login [post]
func (h *UsersHandler) LoginUserHandler(ctx *gin.Context) {
	params := &LoginUserParams{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	user, err := h.useCase.GetUserEntityByEmail(ctx, params.Email)

	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	if user == nil {
		response.JSON(ctx, http.StatusNotFound, false, nil, ErrUserNotFound.Error())
		return
	}

	ok, err := hashing.VerifyPassword(params.Password, user.HashedPassword)

	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	if !ok {
		response.JSON(ctx, http.StatusBadRequest, false, nil, ErrIncorrectAuthData.Error())
		return
	}

	tokenPair, err := jwt.CreateTokenPair(int(user.ID), user.Email, h.jwtSecretKey, accessTTL, refreshTTL)

	ctx.SetCookie(
		"refresh_token",
		tokenPair.RefreshToken,
		60*60*24,
		"/",
		"",
		false,
		true,
	)

	responseData := map[string]string{
		"access_token": tokenPair.AccessToken,
	}

	response.JSON(ctx, http.StatusOK, true, responseData, "")
	return
}

// HandleGetCurrentUser handles the HTTP POST request to get user
// @Summary get user
// @Description get user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /user/me [post]
func (h *UsersHandler) HandleGetCurrentUser(ctx *gin.Context) {
	authUser, _ := ctx.Get("user")
	authUserEmail := authUser.(jwtLib.MapClaims)["user_email"].(string)

	u, err := h.useCase.GetUserEntityByEmail(ctx, authUserEmail)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, false, nil, err.Error())
		return
	}
	response.JSON(ctx, http.StatusOK, true, u, "")
}
