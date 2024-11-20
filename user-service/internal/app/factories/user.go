package factories

import (
	proto "github.com/neokofg/go-pet-detailed-microservices/proto/pb/user/v1"
	"github.com/neokofg/go-pet-detailed-microservices/user-service/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateGrpcUser(user *ent.User) *proto.User {
	protoUser := &proto.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	if user.Avatar != nil {
		protoUser.Avatar = user.Avatar
	}

	return protoUser
}
