package services

import (
	"context"
	"github.com/google/uuid"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent/user"
	"go.uber.org/zap"
)

type UserService struct {
	client          *ent.Client
	Logger          *zap.Logger
	passwordService *PasswordService
}

func NewUserService(client *ent.Client, logger *zap.Logger) *UserService {
	passwordService := NewPasswordService(logger)
	return &UserService{
		client:          client,
		Logger:          logger,
		passwordService: passwordService,
	}
}

func (s *UserService) Register(ctx context.Context, request *proto.RegisterRequest) (*ent.User, error) {
	hashedPassword, err := s.passwordService.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	u, err := s.client.User.Create().
		SetEmail(request.Email).
		SetUsername(request.Username).
		SetPasswordHash(hashedPassword).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) Login(ctx context.Context, request *proto.LoginRequest) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.EmailEQ(request.Email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.passwordService.Compare(request.Password, u.PasswordHash); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) GetUser(ctx context.Context, userId string) (*ent.User, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	u, err := s.client.User.Query().
		Where(user.IDEQ(userUuid)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userId string, request *proto.UpdateUserRequest) (*ent.User, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	builder := s.client.User.UpdateOneID(userUuid)

	if request.Email != nil {
		builder = builder.SetEmail(*request.Email)
	}
	if request.Username != nil {
		builder = builder.SetUsername(*request.Username)
	}
	if request.Avatar != nil {
		builder = builder.SetAvatar(*request.Avatar)
	}

	u, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
