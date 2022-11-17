package main

import (
	"github.com/shsma/grpc-microservice/internal/db"
	"github.com/shsma/grpc-microservice/internal/rocket"
	"github.com/shsma/grpc-microservice/internal/transport/grpc"
	"log"
)

func Run() error {
	// Responsible for init and
	// starting our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	err = rktHandler.Serve()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
