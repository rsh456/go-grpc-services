package grpc

import (
	"context"
	"log"
	"net"

	rkt "github.com/rsh456/go-grpc-protos"
	"github.com/rsh456/go-grpc-services/internal/rocket"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

//Handler - will handle incoming gRPC request
type Handler struct {
	RocketService RocketService
}

//New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen in port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failes to serve: %s\n", err)
		return err
	}
	return nil
}
