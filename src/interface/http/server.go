package http

import (
	"auth-svc/src/interface/container"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartHttpServer(co *container.Container, listener net.Listener) error {
	//-- construct echo
	e := echo.New()
	e.Debug = true
	InitMiddleware(e)

	// setup Handler
	handler := SetupHandlers(co)

	//-- setup router
	SetupRouter(e, handler)

	// start server

	errServe := make(chan error)
	go func() {
		if err := e.Server.Serve(listener); err != nil {
			errServe <- err
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return <-errServe
}
