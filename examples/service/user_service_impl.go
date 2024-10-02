// Generated by Protostub
// see more examples: https://github.com/lovelyoyrmia/protostub
// Implement your logic below

package pb

import (
	"context"
	"log"
	pb "github.com/lovelyoyrmia/protostub/examples/pb"
)

// UserServiceImpl implements the gRPC service for UserServiceImpl.
// This is the server struct that will be used to implement the service methods.
type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

// GetUser handles requests for the GetUser method.
// It takes a GetUserRequest request and returns a GetUserResponse response.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// TODO: Implement your logic here
	log.Println("Received request:", req)
	return &pb.GetUserResponse{}, nil
}

// NewUserServiceServer returns a new instance of UserServiceImpl.
// This is used to initialize the gRPC server with the service implementation.
func NewUserServiceServer() *UserServiceImpl {
	return &UserServiceImpl{}
}
