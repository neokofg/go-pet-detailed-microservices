package factories

import (
	"github.com/neokofg/go-pet-detailed-microservices/news-service/pkg/ent"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/news/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateGrpcNews(news *ent.News) *proto.News {
	protoNews := &proto.News{
		Id:        news.ID.String(),
		Title:     news.Title,
		Content:   news.Content,
		UserId:    news.UserID.String(),
		CreatedAt: timestamppb.New(news.CreatedAt),
		UpdatedAt: timestamppb.New(news.UpdatedAt),
	}

	if news.ImageURL != "" {
		protoNews.ImageUrl = news.ImageURL
	}

	return protoNews
}
