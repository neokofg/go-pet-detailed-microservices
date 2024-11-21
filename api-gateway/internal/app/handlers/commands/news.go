package commands

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	newsProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type NewsCommandsHandler struct {
	logger  *zap.Logger
	newsSvc newsProto.NewsServiceClient
}

func NewNewsCommandsHandler(logger *zap.Logger, newsClient newsProto.NewsServiceClient) *NewsCommandsHandler {
	return &NewsCommandsHandler{
		logger:  logger,
		newsSvc: newsClient,
	}
}

type CreateNewsRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}

func (h *NewsCommandsHandler) CreateNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req CreateNewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.newsSvc.CreateNews(ctx, &newsProto.CreateNewsRequest{
		Token:    c.GetString("token"),
		Title:    req.Title,
		Content:  req.Content,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"news": resp})
}

type UpdateNewsRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"imageUrl"`
}

func (h *NewsCommandsHandler) UpdateNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newsUuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req UpdateNewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateNewsProtoRequest := &newsProto.UpdateNewsRequest{
		Token: c.GetString("token"),
		Id:    newsUuid.String(),
	}
	if req.Title != "" {
		updateNewsProtoRequest.Title = &req.Title
	}
	if req.Content != "" {
		updateNewsProtoRequest.Content = &req.Content
	}
	if req.ImageUrl != "" {
		updateNewsProtoRequest.ImageUrl = &req.ImageUrl
	}

	resp, err := h.newsSvc.UpdateNews(ctx, updateNewsProtoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"news": resp})
}

func (h *NewsCommandsHandler) DeleteNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newsUuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.newsSvc.DeleteNews(ctx, &newsProto.DeleteNewsRequest{
		Token: c.GetString("token"),
		Id:    newsUuid.String(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
