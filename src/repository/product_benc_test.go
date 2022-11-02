package repository

import (
	"go/service1/src/config"
	pb "go/service1/src/protos"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkGetProducts(b *testing.B) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Print(err.Error())
		log.Fatalf("Error when loading .env file")
	}

	config.ConnectDB()

	for i := 0; i < b.N; i++ {
		initProduct.Products(&pb.User{})
	}
}
