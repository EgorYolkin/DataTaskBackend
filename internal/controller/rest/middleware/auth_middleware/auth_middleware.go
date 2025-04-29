package auth_middleware

import (
	"DataTask/pkg/http/response"
	"DataTask/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	jwtSecretKey string
}

func NewAuthMiddleware(jwtSecretKey string) *AuthMiddleware {
	return &AuthMiddleware{
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

		ctx.Set("user", claims)

		ctx.Next()
	}
}
