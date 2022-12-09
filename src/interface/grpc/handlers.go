package grpc

import "auth-svc/src/interface/container"

type handler struct {
	userHandler *userGrpcHandler
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		userHandler: newGrpcHandler(container.UserService),
	}
}

func (h *handler) Validate() *handler {
	if h.userHandler == nil {
		panic("userHandler is nil")
	}
	return h
}
