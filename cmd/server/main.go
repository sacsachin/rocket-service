package main

import (
	"log"

	"example.com/microservice/rocket-service/internal/db"
	"example.com/microservice/rocket-service/internal/rocket"
	"example.com/microservice/rocket-service/internal/transport/grpc"
)

func Run() error {
	// responsible for initializing and starting
	// gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()

	if err != nil {
		log.Println("Failed to run migration.")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
