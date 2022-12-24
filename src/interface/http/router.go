package http

import (
	_ "auth-svc/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"
)

// @title Auth Service API
// @version 1.0
// @description This is a authentication server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email bennie.anware@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://user-svc-hnv6.onrender.com
// @BasePath /v1

func SetupRouter(e *echo.Echo, h *handler) {
	// docs
	e.GET("/docs/*", echoSwagger.WrapHandler)
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
