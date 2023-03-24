package handler

import (
	protoUser "github.com/AdiKhoironHasan/bookservice-protobank/proto/user"
	"github.com/AdiKhoironHasan/bookservices/config"
	"github.com/AdiKhoironHasan/bookservices/domain/service"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	protoUser.UserServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	protoUser.UnimplementedUserServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
