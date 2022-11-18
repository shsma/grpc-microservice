package grpc

import (
	"context"
	"github.com/shsma/grpc-microservice/internal/rocket"
	rkt "github.com/shsma/grpc-microservice/proto/rocket/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
)

// RocketService - define the interface that the
// concrete implementation has to handle
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
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
		log.Print("Could not listen on port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}
	return nil
}

// GetRocket - retrieve a rocket by ID and return the response
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("GetRocket gRPC Endpoint Hit")
	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		log.Print("Failed to retrieve rocket by ID")
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

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC endpoint hit")
	if _, err := strconv.Atoi(req.Rocket.Id); err != nil {
		errorStatus := status.Error(codes.InvalidArgument, "given id is not valid")
		log.Print("given id is not valid")
		return &rkt.AddRocketResponse{}, errorStatus
	}
	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		log.Print("failed to insert rocket into database")
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

// DeleteRocket - handler for deleting a rocket
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("delete rocket gRPC endpoint hit")
	err := h.RocketService.DeleteRocket(ctx, req.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "successfully delete rocket",
	}, nil
}
