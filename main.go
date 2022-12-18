package main

import (
	"auth-svc/src/cmd/grpc"
	"auth-svc/src/cmd/http"
	"auth-svc/src/shared/config"
	"fmt"
	"github.com/soheilhy/cmux"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
)

func main() {
	fmt.Println("program started...")
	cfg := config.InitConfig()
	listener, err := net.Listen("tcp", ":"+cast.ToString(cfg.Apps.Port))
	if err != nil {
		log.Fatal(err)
	}
	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpc.Start(grpcListener) })
	g.Go(func() error { return http.Start(httpListener) })
	g.Go(func() error { return m.Serve() })

	log.Println("run server:", g.Wait())
}
