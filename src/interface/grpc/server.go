package grpc

import (
	"auth-svc/src/interface/container"
	"auth-svc/src/shared/config"
	rpcUser "auth-svc/src/shared/grpc/user"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, container *container.Container, cfg *config.AppsConfig) error {
	listen, err := net.Listen("tcp", cfg.GrpcPort())
	if err != nil {
		return err
	}

	handler := SetupHandlers(container)
	// register service
	server := grpc.NewServer()
	// register server
	rpcUser.RegisterUserServiceServer(server, handler.userHandler)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server... " + cfg.GrpcPort())
	return server.Serve(listen)
}
