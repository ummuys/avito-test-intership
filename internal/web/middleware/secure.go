package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/secure"
)

func Auth(tm secure.TokenManager, access []string) gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			g.Set("msg", "empty token")
			err := models.Error{
				Message: errs.ErrMsgBadToken,
				Code:    errs.ErrCodeBadToken,
			}
			g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			g.Set("msg", "invalid token format")
			err := models.Error{
				Message: errs.ErrMsgBadToken,
				Code:    errs.ErrCodeBadToken,
			}
			g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
			return
		}

		tokenStr := parts[1]
		claims, err := tm.ValidateAccessToken(tokenStr)
		if err != nil {
			g.Set("msg", err.Error())
			err := models.Error{
				Message: errs.ErrMsgBadToken,
				Code:    errs.ErrCodeBadToken,
			}
			g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
			return
		}

		user_id := claims.UserID
		role := claims.Role
		forbidden := true
		for _, acc := range access {
			if acc == role {
				forbidden = false
				break
			}
		}
		if forbidden {
			err := models.Error{
				Message: errs.ErrMsgNotFound,
				Code:    errs.ErrCodeNotFound,
			}
			g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
			return
		}
		g.Set("user_id", user_id)
		g.Next()
	}
}
