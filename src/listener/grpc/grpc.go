package grpc

import (
	pb "go/service1/src/protos"
	"go/service1/src/repository"
	"go/service1/src/service/auth"
	"go/service1/src/service/images"
	"go/service1/src/service/products"
	"go/service1/src/service/users"
	"log"
	"net"

	"google.golang.org/grpc"
)

func registerService(s *grpc.Server) {
	pb.RegisterAuthRPCServer(s, &auth.AuthImpl{
		Repository: repository.NewAuthRepository(),
	})
	pb.RegisterUserRPCServer(s, &users.UserImpl{
		Repository: repository.NewUserRepository(),
	})
	pb.RegisterProductRPCServer(s, &products.ProductsImpl{
		Repository: repository.NewProductRepo(),
	})
	pb.RegisterImageRPCServer(s, &images.ImageImpl{
		Repository: repository.NewImageRepository(),
	})
}

func Run(address string) error {
	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("failed listen grpc server")
	}

	server := grpc.NewServer()
	registerService(server)
	return server.Serve(listen)
}
