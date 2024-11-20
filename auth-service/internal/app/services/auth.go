package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent/personalaccesstoken"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	"go.uber.org/zap"
	"time"
)

type AuthService struct {
	client *ent.Client
	Logger *zap.Logger
}

func NewAuthService(client *ent.Client, logger *zap.Logger) *AuthService {
	return &AuthService{
		client: client,
		Logger: logger,
	}
}

func (s *AuthService) CreateToken(ctx context.Context, request *proto.CreateTokenRequest) (string, error) {
	token := uuid.New().String()
	hash := sha256.Sum256([]byte(token))

	userUuid, err := uuid.Parse(request.UserId)
	if err != nil {
		return "", err
	}

	_, err = s.client.PersonalAccessToken.Create().
		SetTokenHash(hex.EncodeToString(hash[:])).
		SetName(request.Name).
		SetUserID(userUuid).
		SetAbilities(request.Abilities).
		SetExpiresAt(time.Now().AddDate(1, 0, 0)).
		Save(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*ent.PersonalAccessToken, error) {
	hash := sha256.Sum256([]byte(token))
	hashStr := hex.EncodeToString(hash[:])

	pat, err := s.client.PersonalAccessToken.Query().
		Where(personalaccesstoken.TokenHashEQ(hashStr)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if pat.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	pat, err = pat.Update().
		SetLastUsed(time.Now()).
		Save(ctx)

	return pat, err
}

func (s *AuthService) DeleteToken(ctx context.Context, token string) error {
	hash := sha256.Sum256([]byte(token))
	hashStr := hex.EncodeToString(hash[:])

	pat, err := s.client.PersonalAccessToken.Query().
		Where(personalaccesstoken.TokenHashEQ(hashStr)).
		Only(ctx)
	if err != nil {
		return err
	}

	if err := s.client.PersonalAccessToken.DeleteOne(pat).Exec(ctx); err != nil {
		return err
	}

	return nil
}
