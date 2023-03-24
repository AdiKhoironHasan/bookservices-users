package client

import (
	"flag"
	"fmt"

	"github.com/AdiKhoironHasan/bookservices/config"
	"github.com/AdiKhoironHasan/bookservices/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// cfg        = config.New()
	serverHost = "localhost"
	serverPort = 9092
	DSN        = fmt.Sprintf("%s:%d", serverHost, serverPort)
)

var (
	addr = flag.String("addr", DSN, "The address to connect")
)

// NewGRPCConn is a constructor
func NewGRPCConn(_ *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
