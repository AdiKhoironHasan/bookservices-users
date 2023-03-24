package handler

import (
	"github.com/AdiKhoironHasan/bookservices-users/config"
	"github.com/AdiKhoironHasan/bookservices-users/domain/service"
	"github.com/AdiKhoironHasan/bookservices-users/grpc/client"
	"github.com/AdiKhoironHasan/bookservices-users/proto/user"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	user.UserServiceServer
}

// Handler is struct
type Handler struct {
	config     *config.Config
	repo       *service.Repositories
	grpcClient *client.GRPCClient

	user.UnimplementedUserServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories, grpcClient *client.GRPCClient) *Handler {
	return &Handler{
		config:     conf,
		repo:       repo,
		grpcClient: grpcClient,
	}
}

var _ Interface = &Handler{}
