package http

import "auth-svc/src/interface/container"

type handler struct {
	userHandler *userHttpHandler
}

func SetupHandlers(container *container.Container) *handler {
	return &handler{
		userHandler: newHttpHandler(container.UserService),
	}
}

func (h *handler) Validate() *handler {
	if h.userHandler == nil {
		panic("userHandler is nil")
	}
	return h
}
