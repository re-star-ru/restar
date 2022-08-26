package user

import (
	"context"
	"google.golang.org/grpc"
	"restar/pkg/domain"
	"restar/pkg/user/pb"
)

type IUser interface {
	Create(user *domain.User) error
}

type GRPCHandler struct {
	pb.UnimplementedUserServiceServer
	userUsecase IUser
}

func RegisterService(srv grpc.ServiceRegistrar, userUsecase IUser) {
	srv.RegisterService(
		&pb.UserService_ServiceDesc,
		&GRPCHandler{
			userUsecase: userUsecase,
		},
	)
}

func (g *GRPCHandler) UserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Name: "name is :" + req.Id}, nil
}
