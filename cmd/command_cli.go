package cmd

import (
	"fmt"

	"github.com/AdiKhoironHasan/bookservices-users/grpc/server"
	"github.com/AdiKhoironHasan/bookservices-users/infrastructure/persistence"
	"github.com/urfave/cli/v2"
)

// newGRPCServer is a method command cli to run grpc
func (cmd *Command) newGRPCServer() *cli.Command {
	return &cli.Command{
		Name:  "grpc:start",
		Usage: "A command to run gRPC server",
		Action: func(_ *cli.Context) error {
			grpcServer := server.NewGRPCServer(
				server.WithConfig(cmd.conf),
				server.WithRepository(cmd.repo),
				server.WithGRPCClient(cmd.grpcClient),
			)
			err := grpcServer.Run(cmd.conf.GRPCPort)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

// newDBMigrate is a method command cli to run db migration
func (cmd *Command) newDBMigrate() *cli.Command {
	return &cli.Command{
		Name:  "db:migrate",
		Usage: "A command to run database migration",
		Action: func(_ *cli.Context) error {
			db, errConn := persistence.NewDBConnection(cmd.conf.DBConfig)
			if errConn != nil {
				return fmt.Errorf("unable to connect to database: %w", errConn)
			}

			err := persistence.AutoMigrate(db)
			if err != nil {
				return fmt.Errorf("cannot run auto migrate: %w", err)
			}

			return nil
		},
	}
}
