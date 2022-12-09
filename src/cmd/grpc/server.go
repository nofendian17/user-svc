package grpc

import (
	"auth-svc/src/interface/container"
	"auth-svc/src/interface/grpc"
	"auth-svc/src/shared/config"
	"context"
	"fmt"
)

func Start() {

	//-- config file and port
	cfg := config.InitConfig()

	//-- setup container
	co := container.Setup(cfg)

	// start server grpc
	err := grpc.RunServer(context.Background(), co, cfg.Apps)
	if err != nil {
		fmt.Printf("error %s", err.Error())
	}
}
