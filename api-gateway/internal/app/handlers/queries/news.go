package queries

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	cache "github.com/neokofg/go-pet-detailed-microservices/api-gateway/pkg/redis"
	newsProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type NewsQueriesHandler struct {
	logger  *zap.Logger
	cache   *cache.Cache
	newsSvc newsProto.NewsServiceClient
}

func NewNewsQueriesHandler(logger *zap.Logger, cache *cache.Cache, newsClient newsProto.NewsServiceClient) *NewsQueriesHandler {
	return &NewsQueriesHandler{
		logger:  logger,
		cache:   cache,
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

	var cachedResp *newsProto.GetNewsFeedResponse
	cacheName := fmt.Sprintf("news:%d_%d", req.Page, req.PageSize)
	err := h.cache.Get(ctx, cacheName, &cachedResp)
	if err == nil && cachedResp != nil {
		c.JSON(http.StatusOK, cachedResp)
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

	if err := h.cache.Set(ctx, cacheName, resp, 5*time.Minute); err != nil {
		log.Printf("Failed to cache response: %v", err)
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
