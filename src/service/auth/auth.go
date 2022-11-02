package auth

import (
	"context"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
	"log"
)

type AuthService interface {
	LoginUser(ctx context.Context, input *pb.DataLogin) (*pb.ResponseLogin, error)
	RegisterUser(ctx context.Context, input *pb.DataRegister) (*pb.ResponseRegister, error)
	SelectSessionUserById(ctx context.Context, input *pb.DataSession) (*pb.ResponseSession, error)
}

type AuthImpl struct {
	pb.UnimplementedAuthRPCServer
	Repository       repository.AuthRepository
	responseRegister *pb.ResponseRegister
	responseLogin    *pb.ResponseLogin
	responseSession  *pb.ResponseSession
}

func (a *AuthImpl) LoginUser(ctx context.Context, input *pb.DataLogin) (*pb.ResponseLogin, error) {
	res, err := a.Repository.LoginUser(input)

	if err != nil {
		log.Print(err.Error())
		return nil, nil
	}

	// store session in to redis
	// rDB := RedisClient()
	// setSessionInToRedis := rDB.Set(u.user.UserId, u.user.UserSession, time.Duration(18000)*time.Second)

	// if err := setSessionInToRedis.Err(); err != nil {
	// 	log.Print(err.Error())
	// }

	a.responseLogin = &pb.ResponseLogin{
		Status:  200,
		Message: "ok",
		UserId:  res.UserId,
		Session: res.Session,
	}

	return a.responseLogin, nil
}

func (a *AuthImpl) RegisterUser(ctx context.Context, input *pb.DataRegister) (*pb.ResponseRegister, error) {
	res, err := a.Repository.RegisterUser(input)

	if err != nil || res.Status != "ok" {
		log.Print(err.Error())
		return nil, err
	}

	a.responseRegister = &pb.ResponseRegister{
		Status:  200,
		Message: res.Status,
	}

	return a.responseRegister, nil
}

func (a *AuthImpl) SelectSessionUserById(ctx context.Context, input *pb.DataSession) (*pb.ResponseSession, error) {
	res, err := a.Repository.SelectSessionUserById(input)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	a.responseSession = &pb.ResponseSession{
		UserSession: res.UserSession,
		RememberMe:  res.RememberMe,
	}

	return a.responseSession, nil
}
