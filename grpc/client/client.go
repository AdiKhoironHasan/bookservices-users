package client

import (
	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	Book protoBook.BookServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(
	connBook grpc.ClientConnInterface,
) *GRPCClient {
	return &GRPCClient{
		Book: protoBook.NewBookServiceClient(connBook),
	}
}
