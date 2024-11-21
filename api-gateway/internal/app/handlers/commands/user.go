package commands

import (
	"context"
	"github.com/gin-gonic/gin"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserCommandsHandler struct {
	logger  *zap.Logger
	userSvc userProto.UserServiceClient
}

func NewUserCommandsHandler(logger *zap.Logger, userClient userProto.UserServiceClient) *UserCommandsHandler {
	return &UserCommandsHandler{
		logger:  logger,
		userSvc: userClient,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"required,min=3,max=32"`
}

func (h *UserCommandsHandler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.userSvc.Register(ctx, &userProto.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserCommandsHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userSvc.Login(ctx, &userProto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}

func (h *UserCommandsHandler) Logout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.userSvc.Logout(ctx, &userProto.LogoutRequest{
		Token: c.GetString("token"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}

type UpdateRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (h *UserCommandsHandler) UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	protoRequest := &userProto.UpdateUserRequest{
		Token: c.GetString("token"),
	}
	if req.Email != "" {
		protoRequest.Email = &req.Email
	}
	if req.Username != "" {
		protoRequest.Username = &req.Username
	}
	if req.Avatar != "" {
		protoRequest.Avatar = &req.Avatar
	}

	resp, err := h.userSvc.UpdateUser(ctx, protoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}
