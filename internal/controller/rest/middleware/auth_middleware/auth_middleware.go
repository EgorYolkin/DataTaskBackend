package auth_middleware

import (
	"DataTask/internal/usecase/user_usecase"
	"DataTask/pkg/http/response"
	"DataTask/pkg/jwt"
	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"net/http"
)

type AuthMiddleware struct {
	jwtSecretKey string
	useCase      user_usecase.UserUseCase
}

func NewAuthMiddleware(useCase user_usecase.UserUseCase, jwtSecretKey string) *AuthMiddleware {
	return &AuthMiddleware{
		useCase:      useCase,
		jwtSecretKey: jwtSecretKey,
	}
}

func (m *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		if authorization == "" {
			ctx.Abort()
			response.JSON(ctx, http.StatusUnauthorized, false, nil, "Authorization header is empty")
			return
		}

		refreshToken, err := ctx.Cookie("refresh_token")

		if err != nil || len(refreshToken) == 0 {
			ctx.Abort()
			response.JSON(ctx, http.StatusUnauthorized, false, nil, "refresh_token cookie is empty")
			return
		}

		claims, err := jwt.VerifyJWT(authorization, m.jwtSecretKey)

		if err != nil {
			ctx.Abort()
			response.JSON(ctx, http.StatusUnauthorized, false, nil, err.Error())
			return
		}

		userEmail := claims.(jwtLib.MapClaims)["user_email"].(string)

		u, _ := m.useCase.GetUserEntityByEmail(ctx, userEmail)

		ctx.Set("user", claims)
		ctx.Set("user_email", userEmail)
		ctx.Set("user_id", u.ID)

		ctx.Next()
	}
}
