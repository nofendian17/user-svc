package http

import (
	"auth-svc/src/interface/container"
	"auth-svc/src/interface/http"
	"auth-svc/src/shared/config"
)

func Start() {

	//-- config file and port
	cfg := config.InitConfig()

	//-- setup container
	co := container.Setup(cfg)

	// start server http
	http.StartHttpServer(co, cfg.Apps)

}
