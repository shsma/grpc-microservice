// +acceptance

package test

import (
	"log"

	rocket "github.com/shsma/grpc-microservice/proto/rocket/v1"
	"google.golang.org/grpc"
)

func GetClient() rocket.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	rocketClient := rocket.NewRocketServiceClient(conn)
	return rocketClient
}
