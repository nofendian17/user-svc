package http

import (
	"auth-svc/src/interface/container"
	"auth-svc/src/interface/http"
	"auth-svc/src/shared/config"
	"net"
)

func Start(listener net.Listener) error {

	//-- config file and port
	cfg := config.InitConfig()

	//-- setup container
	co := container.Setup(cfg)

	// start server http
	err := http.StartHttpServer(co, listener)

	return err
}
