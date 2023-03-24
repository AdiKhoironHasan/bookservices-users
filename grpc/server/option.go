package server

import (
	"github.com/AdiKhoironHasan/bookservices-users/config"
	"github.com/AdiKhoironHasan/bookservices-users/domain/service"
	"github.com/AdiKhoironHasan/bookservices-users/grpc/client"
)

// WithConfig is function
func WithConfig(config *config.Config) ServerGrpcOption {
	return func(r *Server) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) ServerGrpcOption {
	return func(r *Server) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(gClient *client.GRPCClient) ServerGrpcOption {
	return func(s *Server) {
		s.grpcClient = gClient
	}
}
