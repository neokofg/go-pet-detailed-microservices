package middleware

import (
	"github.com/gin-gonic/gin"
	authProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	"net/http"
)

func AuthMiddleware(authSvc authProto.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		resp, err := authSvc.ValidateToken(c.Request.Context(), &authProto.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", resp.UserId)
		c.Set("token", token)
		c.Next()
	}
}
