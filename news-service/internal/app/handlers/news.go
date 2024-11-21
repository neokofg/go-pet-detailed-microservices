package handlers

import (
	"context"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/internal/app/factories"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/internal/app/services"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	userProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"go.uber.org/zap"
	"math"
)

type NewsHandler struct {
	proto.UnimplementedNewsServiceServer
	userSvc     userProto.UserServiceClient
	newsService services.NewsService
}

func NewNewsHandler(userClient userProto.UserServiceClient, client *ent.Client, logger *zap.Logger) *NewsHandler {
	return &NewsHandler{
		userSvc:     userClient,
		newsService: *services.NewNewsService(client, logger),
	}
}

func (h *NewsHandler) GetNewsFeed(ctx context.Context, request *proto.GetNewsFeedRequest) (*proto.GetNewsFeedResponse, error) {
	n, current, total, err := h.newsService.GetNewsFeed(ctx, request)
	if err != nil {
		return nil, err
	}
	pages := int32(math.Ceil(float64(total) / float64(request.PageSize)))
	return &proto.GetNewsFeedResponse{
		News:        n,
		TotalPages:  pages,
		TotalItems:  int32(current),
		CurrentPage: request.Page,
	}, nil
}

func (h *NewsHandler) GetNewsById(ctx context.Context, request *proto.GetNewsByIdRequest) (*proto.News, error) {
	n, err := h.newsService.GetNewsById(ctx, request)
	if err != nil {
		return nil, err
	}

	return factories.CreateGrpcNews(n), nil
}

func (h *NewsHandler) CreateNews(ctx context.Context, request *proto.CreateNewsRequest) (*proto.News, error) {
	resp, err := h.userSvc.GetUser(ctx, &userProto.GetUserRequest{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}

	n, err := h.newsService.CreateNews(ctx, resp.User.Id, request)
	if err != nil {
		return nil, err
	}

	return factories.CreateGrpcNews(n), nil
}

func (h *NewsHandler) DeleteNews(ctx context.Context, request *proto.DeleteNewsRequest) (*proto.OperationResponse, error) {
	resp, err := h.userSvc.GetUser(ctx, &userProto.GetUserRequest{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}

	if err := h.newsService.DeleteNews(ctx, resp.User.Id, request); err != nil {
		return nil, err
	}

	return &proto.OperationResponse{
		Success: true,
		Message: "Success",
	}, nil
}

func (h *NewsHandler) UpdateNews(ctx context.Context, request *proto.UpdateNewsRequest) (*proto.News, error) {
	resp, err := h.userSvc.GetUser(ctx, &userProto.GetUserRequest{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}

	n, err := h.newsService.UpdateNews(ctx, resp.User.Id, request)
	if err != nil {
		return nil, err
	}

	return factories.CreateGrpcNews(n), nil
}
