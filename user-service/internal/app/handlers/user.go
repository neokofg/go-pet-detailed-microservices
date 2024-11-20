package handlers

import (
	"context"
	authProto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/auth/v1"
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/app/factories"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/internal/app/services"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	proto.UnimplementedUserServiceServer
	authSvc     authProto.AuthServiceClient
	userService *services.UserService
}

func NewUserHandler(client *ent.Client, logger *zap.Logger, authSvc authProto.AuthServiceClient) *UserHandler {
	userService := services.NewUserService(client, logger)
	return &UserHandler{
		authSvc:     authSvc,
		userService: userService,
	}
}

func (h *UserHandler) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	u, err := h.userService.Register(ctx, request)
	if err != nil {
		h.userService.Logger.Error("Failed to create user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	defaultAbilities := []string{"auth"}

	resp, err := h.authSvc.CreateToken(ctx, &authProto.CreateTokenRequest{
		UserId:    u.ID.String(),
		Name:      u.Username,
		Abilities: defaultAbilities,
	})
	if err != nil {
		h.userService.Logger.Error("Failed to create token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create token")
	}

	return &proto.RegisterResponse{
		UserId: u.ID.String(),
		Token:  resp.Token,
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	u, err := h.userService.Login(ctx, request)
	if err != nil {
		h.userService.Logger.Error("Failed to login user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to login user")
	}

	defaultAbilities := []string{"auth"}

	resp, err := h.authSvc.CreateToken(ctx, &authProto.CreateTokenRequest{
		UserId:    u.ID.String(),
		Name:      u.Username,
		Abilities: defaultAbilities,
	})
	if err != nil {
		h.userService.Logger.Error("Failed to create token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create token")
	}

	return &proto.LoginResponse{
		Token: resp.Token,
		User:  factories.CreateGrpcUser(u),
	}, nil
}

func (h *UserHandler) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	_, err := h.authSvc.RevokeToken(ctx, &authProto.RevokeTokenRequest{
		Token: request.Token,
	})
	if err != nil {
		h.userService.Logger.Error("Failed to revoke token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to revoke token")
	}

	return &proto.LogoutResponse{}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	resp, err := h.authSvc.ValidateToken(ctx, &authProto.ValidateTokenRequest{
		Token: request.Token,
	})
	if err != nil {
		h.userService.Logger.Error("Failed to validate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to validate token")
	}

	u, err := h.userService.GetUser(ctx, resp.UserId)
	if err != nil {
		h.userService.Logger.Error("Failed to get user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get user")
	}

	return &proto.GetUserResponse{
		User: factories.CreateGrpcUser(u),
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	resp, err := h.authSvc.ValidateToken(ctx, &authProto.ValidateTokenRequest{
		Token: request.Token,
	})
	if err != nil {
		h.userService.Logger.Error("Failed to validate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to validate token")
	}

	u, err := h.userService.UpdateUser(ctx, resp.UserId, request)
	if err != nil {
		h.userService.Logger.Error("Failed to update user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to update user")
	}

	return &proto.UpdateUserResponse{
		User: factories.CreateGrpcUser(u),
	}, nil
}
