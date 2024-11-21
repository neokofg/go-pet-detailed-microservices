package queries

import (
	"context"
	"github.com/gin-gonic/gin"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserQueriesHandler struct {
	logger  *zap.Logger
	userSvc userProto.UserServiceClient
}

func NewUserQueriesHandler(logger *zap.Logger, userClient userProto.UserServiceClient) *UserQueriesHandler {
	return &UserQueriesHandler{
		logger:  logger,
		userSvc: userClient,
	}
}

func (h *UserQueriesHandler) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userSvc.GetUser(ctx, &userProto.GetUserRequest{
		Token: c.GetString("token"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
