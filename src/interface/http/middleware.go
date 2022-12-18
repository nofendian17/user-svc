package http

import (
	"auth-svc/src/shared/constant"
	rpcUser "auth-svc/src/shared/grpc/user"
	"auth-svc/src/shared/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitMiddleware(e *echo.Echo) {
	// - setup cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderAcceptEncoding,
			echo.HeaderXCSRFToken, echo.HeaderXRequestID, "token"},
		ExposeHeaders: []string{echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials, echo.HeaderContentType},
		AllowCredentials: true,
	}))

	e.HTTPErrorHandler = errorHandler

	e.Validator = &DataValidator{ValidatorData: validator.New()}

	// - panic recover
	e.Use(middleware.Recover())
}

func errorHandler(err error, c echo.Context) {
	code := constant.CodeProcessingError

	response := &rpcUser.ErrorResponse{
		Message: "error",
		Code:    code,
	}

	if he, ok := err.(*helper.ApplicationError); ok {
		response.Code = he.Code
		response.Message = he.Error()
	} else if he, ok := err.(*echo.HTTPError); ok {
		response.Message = he.Error()
	} else {
		response.Message = err.Error()
	}
	err = c.JSON(http.StatusOK, response)
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}
