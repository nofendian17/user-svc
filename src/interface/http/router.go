package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SetupRouter(e *echo.Echo, h *handler) {
	// health check
	e.GET("/health", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})

	v1 := e.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", h.userHandler.Register)
			user.POST("/auth", h.userHandler.Login)
			user.POST("/refresh", h.userHandler.Refresh)
			user.GET("/me", h.userHandler.User)
		}
	}
}
