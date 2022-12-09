package http

import (
	"auth-svc/src/interface/container"
	"auth-svc/src/shared/config"
	"context"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartHttpServer(co *container.Container, cfg *config.AppsConfig) {
	//-- construct echo
	e := echo.New()
	e.Debug = true
	InitMiddleware(e, co, cfg)

	// setup Handler
	handler := SetupHandlers(co)

	//-- setup router
	SetupRouter(e, handler)

	// start server
	go func() {
		if err := e.Start(cfg.AppPort()); err != nil {
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
