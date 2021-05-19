package delivery

import (
	"context"

	"github.com/rssh-jp/api-develop/domain"
	"github.com/rssh-jp/api-develop/internal/grpc/pb"

	"google.golang.org/grpc"
)

type userGRPCDelivery struct {
	userUsecase domain.UserUsecase
	pb.UnimplementedUsersServer
}

func HandleUserGRPCDelivery(userUsecase domain.UserUsecase) *grpc.Server {
	handler := &userGRPCDelivery{
		userUsecase: userUsecase,
	}

	s := grpc.NewServer()

	pb.RegisterUsersServer(s, handler)

	return s
}

func (ud *userGRPCDelivery) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchReply, error) {
	users, err := ud.userUsecase.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	retUsers := make([]*pb.User, 0, len(users))
	for _, user := range users {
		retUsers = append(retUsers, &pb.User{
			Id:   user.ID,
			Name: user.Name,
			Age:  int32(user.Age),
		})
	}

	return &pb.FetchReply{
		Users: retUsers,
	}, nil
}

func (ud *userGRPCDelivery) GetByID(ctx context.Context, id int64) error {
	return nil
}

func (ud *userGRPCDelivery) Update(ctx context.Context) error {
	return nil
}
