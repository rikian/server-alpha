package main

import (
	"go/service1/src/config"
	"go/service1/src/listener/grpc"
	"go/service1/src/listener/postgres"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error when loading .env file")
	}

	postgres.ConnectDB()

	log.Print("Grpc service 1 running at " + config.Address)

	if err := grpc.Run(config.Address); err != nil {
		log.Fatalf("Failed running grpc server at %v", config.Address)
	}
}
