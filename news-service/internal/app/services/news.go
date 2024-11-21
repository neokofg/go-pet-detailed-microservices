package services

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/internal/app/factories"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent"
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent/news"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"go.uber.org/zap"
	"time"
)

type NewsService struct {
	client *ent.Client
	Logger *zap.Logger
}

func NewNewsService(client *ent.Client, logger *zap.Logger) *NewsService {
	return &NewsService{client, logger}
}

func (s *NewsService) GetNewsFeed(ctx context.Context, req *proto.GetNewsFeedRequest) ([]*proto.News, int, int, error) {
	var results []struct {
		ID         uuid.UUID `sql:"id"`
		Title      string    `sql:"title"`
		Content    string    `sql:"content"`
		ImageURL   string    `sql:"image_url"`
		UserID     uuid.UUID `sql:"user_id"`
		CreatedAt  time.Time `sql:"created_at"`
		UpdatedAt  time.Time `sql:"updated_at"`
		TotalCount int       `sql:"total_count"`
	}

	err := s.client.News.Query().
		Limit(int(req.PageSize)).
		Offset(int((req.Page-1)*req.PageSize)).
		GroupBy(
			news.FieldID,
			news.FieldTitle,
			news.FieldContent,
			news.FieldUserID,
			news.FieldImageURL,
			news.FieldCreatedAt,
			news.FieldUpdatedAt,
		).
		Aggregate(func(s *sql.Selector) string {
			return sql.As(sql.Count("*"), "total_count")
		}).
		Scan(ctx, &results)

	if err != nil {
		return nil, 0, 0, err
	}

	protoNews := make([]*proto.News, len(results))
	var totalCount int

	if len(results) > 0 {
		totalCount = results[0].TotalCount
		for i, r := range results {
			n := &ent.News{
				ID:        r.ID,
				Title:     r.Title,
				Content:   r.Content,
				UserID:    r.UserID,
				ImageURL:  r.ImageURL,
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
			}
			protoNews[i] = factories.CreateGrpcNews(n)
		}
	}

	currentCount := len(protoNews)

	return protoNews, currentCount, totalCount, nil
}

func (s *NewsService) GetNewsById(ctx context.Context, req *proto.GetNewsByIdRequest) (*ent.News, error) {
	newsUuid, err := uuid.Parse(req.Id)
	if err != nil {
		s.Logger.Error("failed to parse news uuid", zap.Error(err))
		return nil, err
	}

	n, err := s.FindNews(ctx, newsUuid)
	if err != nil {
		s.Logger.Error("failed to find news", zap.Error(err))
		return nil, err
	}

	return n, nil
}

func (s *NewsService) CreateNews(ctx context.Context, userId string, req *proto.CreateNewsRequest) (*ent.News, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		s.Logger.Error("failed to parse user uuid", zap.Error(err))
		return nil, err
	}

	n, err := s.client.News.Create().
		SetUserID(userUuid).
		SetTitle(req.Title).
		SetContent(req.Content).
		SetImageURL(req.ImageUrl).
		Save(ctx)
	if err != nil {
		s.Logger.Error("failed to create news", zap.Error(err))
		return nil, err
	}
	return n, nil
}

func (s *NewsService) DeleteNews(ctx context.Context, userId string, req *proto.DeleteNewsRequest) error {
	newsUuid, err := uuid.Parse(req.Id)
	if err != nil {
		s.Logger.Error("failed to parse news uuid", zap.Error(err))
		return err
	}

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		s.Logger.Error("failed to parse user uuid", zap.Error(err))
		return err
	}

	n, err := s.FindNews(ctx, newsUuid)
	if err != nil {
		s.Logger.Error("failed to fetch news by id", zap.Error(err))
		return err
	}

	if n.UserID != userUuid {
		return errors.New("not owner")
	}

	if err := s.client.News.DeleteOne(n).Exec(ctx); err != nil {
		s.Logger.Error("failed to delete news", zap.Error(err))
		return err
	}

	return nil
}

func (s *NewsService) UpdateNews(ctx context.Context, userId string, req *proto.UpdateNewsRequest) (*ent.News, error) {
	newsUuid, err := uuid.Parse(req.Id)
	if err != nil {
		s.Logger.Error("failed to parse news uuid", zap.Error(err))
		return nil, err
	}

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		s.Logger.Error("failed to parse user uuid", zap.Error(err))
		return nil, err
	}

	n, err := s.FindNews(ctx, newsUuid)
	if err != nil {
		s.Logger.Error("failed to fetch news by id", zap.Error(err))
		return nil, err
	}

	if n.UserID != userUuid {
		return nil, errors.New("not owner")
	}

	builder := s.client.News.UpdateOne(n)
	if req.Title != nil {
		builder = builder.SetTitle(*req.Title)
	}
	if req.Content != nil {
		builder = builder.SetContent(*req.Content)
	}
	if req.ImageUrl != nil {
		builder = builder.SetImageURL(*req.ImageUrl)
	}
	n, err = builder.Save(ctx)
	if err != nil {
		s.Logger.Error("failed to update news", zap.Error(err))
		return nil, err
	}

	return n, nil
}

func (s *NewsService) FindNews(ctx context.Context, newsUuid uuid.UUID) (*ent.News, error) {
	n, err := s.client.News.Query().
		Where(news.IDEQ(newsUuid)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return n, nil
}
