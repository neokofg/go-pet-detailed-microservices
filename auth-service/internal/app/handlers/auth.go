package handlers

import (
	"context"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/internal/app/services"
	"github.com/neokofg/go-pet-detailed-microservices/auth-service/pkg/ent"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	proto.UnimplementedAuthServiceServer
	authService *services.AuthService
}

func NewAuthHandler(client *ent.Client, logger *zap.Logger) *AuthHandler {
	authService := services.NewAuthService(client, logger)
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) CreateToken(ctx context.Context, req *proto.CreateTokenRequest) (*proto.CreateTokenResponse, error) {
	token, err := h.authService.CreateToken(
		ctx,
		req,
	)
	if err != nil {
		h.authService.Logger.Error("failed to create token", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to create token")
	}

	return &proto.CreateTokenResponse{
		Token: token,
	}, nil
}

func (h *AuthHandler) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	pat, err := h.authService.ValidateToken(
		ctx,
		req.Token,
	)
	if err != nil {
		h.authService.Logger.Error("Failed to validate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to validate token")
	}

	return &proto.ValidateTokenResponse{
		UserId:    pat.UserID.String(),
		Abilities: pat.Abilities,
		IsValid:   true,
	}, nil
}

func (h *AuthHandler) RevokeToken(ctx context.Context, req *proto.RevokeTokenRequest) (*proto.RevokeTokenResponse, error) {
	if err := h.authService.DeleteToken(ctx, req.Token); err != nil {
		h.authService.Logger.Error("Failed to revoke token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to revoke token")
	}

	return &proto.RevokeTokenResponse{}, nil
}
