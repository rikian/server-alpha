package users

import (
	"context"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
)

type UsersService interface {
	SelectUser(ctx context.Context, input *pb.DataSelectUser) (*pb.ResponseSelectUser, error)
}

type UserImpl struct {
	pb.UnimplementedUserRPCServer
	Repository repository.UserRepository
}

func (s *UserImpl) SelectUser(ctx context.Context, input *pb.DataSelectUser) (*pb.ResponseSelectUser, error) {
	res, err := s.Repository.SelectUser(input)

	if err != nil {
		return nil, err
	}

	return res, nil
}
