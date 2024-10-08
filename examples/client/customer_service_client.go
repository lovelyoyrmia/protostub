// Generated by Protostub
// see more examples: https://github.com/lovelyoyrmia/protostub
// Implement your logic below

package client

import (
    pb "github.com/lovelyoyrmia/protostub/examples/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CustomerServiceClient implements the gRPC client for CustomerService.
// It holds the gRPC client that will be used to make remote procedure calls to the CustomerService.
type CustomerServiceClient struct {
	Client pb.CustomerServiceClient
}

// InitCustomerServiceClient initializes a new gRPC client for CustomerService.
// It takes a server URL as input, sets up a gRPC connection, and returns the initialized client.
//
// url: The address of the gRPC server to connect to.
//
// Returns:
// - A pointer to the initialized CustomerServiceClient.
// - An error if the connection setup fails.
func InitCustomerServiceClient(url string) (*CustomerServiceClient, error) {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	c := CustomerServiceClient{
		Client: pb.NewCustomerServiceClient(cc),
	}

	return &c, nil
}