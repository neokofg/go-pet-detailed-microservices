package queries

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	newsProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type NewsQueriesHandler struct {
	logger  *zap.Logger
	newsSvc newsProto.NewsServiceClient
}

func NewNewsQueriesHandler(logger *zap.Logger, newsClient newsProto.NewsServiceClient) *NewsQueriesHandler {
	return &NewsQueriesHandler{
		logger:  logger,
		newsSvc: newsClient,
	}
}

type GetNewsFeedRequest struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

func (h *NewsQueriesHandler) GetNewsFeed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req GetNewsFeedRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.newsSvc.GetNewsFeed(ctx, &newsProto.GetNewsFeedRequest{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *NewsQueriesHandler) GetNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newsUuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.newsSvc.GetNewsById(ctx, &newsProto.GetNewsByIdRequest{
		Id: newsUuid.String(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
