package grpc

import (
	"context"
	"log"
	"net"

	rkt "example.com/microservice/rocket-proto/rocket/v1"
	"example.com/microservice/rocket-service/internal/rocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RocketService - define the interface that the concrete implementation
// has to adhere to
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler will handle incoming grpc request
type Handler struct {
	RocketService RocketService
	rkt.UnimplementedRocketServiceServer
}

// New - return a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Could not listen on port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve : %s\n", err)
		return err
	}

	return nil
}

//GetRocket - retrive a rocket by id and return response.
func (h Handler) GetRocket(ctx context.Context, in *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get rocket gRPC endpoint hit.")

	rocket, err := h.RocketService.GetRocketByID(ctx, in.Id)
	if err != nil {
		log.Print("Failed top retrive rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

//AddRocket - adds a rocket to database and return response
func (h Handler) AddRocket(ctx context.Context, in *rkt.AddRockerRequest) (*rkt.AddRockerResponse, error) {
	log.Print("Add rrocket gRPC endpoint hit")

	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   in.Rocket.Id,
		Name: in.Rocket.Name,
		Type: in.Rocket.Type,
	})
	if err != nil {
		log.Print("Failed to insert rocket into database.")
		return &rkt.AddRockerResponse{}, err
	}

	return &rkt.AddRockerResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

//DeleteRocket - deletes a rocket from database and return status.
func (h Handler) DeleteRocket(ctx context.Context, in *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("Delete rocket gRPC endpoint hit")
	err := h.RocketService.DeleteRocket(ctx, in.Id)
	if err != nil {
		log.Print(err.Error())
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "Sucessfully delete rocket.",
	}, nil
}
