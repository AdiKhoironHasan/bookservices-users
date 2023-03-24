package cmd

import (
	"github.com/AdiKhoironHasan/bookservices-users/config"
	"github.com/AdiKhoironHasan/bookservices-users/domain/service"
	"github.com/AdiKhoironHasan/bookservices-users/grpc/client"
)

// CommandOption is an option type
type CommandOption func(c *Command)

// WithConfig is a function of command option
func WithConfig(conf *config.Config) CommandOption {
	return func(c *Command) {
		c.conf = conf
	}
}

// WithRepo is a function of command option
func WithRepo(repo *service.Repositories) CommandOption {
	return func(c *Command) {
		c.repo = repo
	}
}

// WithGRPCClient is function of command option
func WithGRPCClient(gClient *client.GRPCClient) CommandOption {
	return func(r *Command) {
		r.grpcClient = gClient
	}
}
