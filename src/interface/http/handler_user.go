package http

import (
	"auth-svc/src/interface/usecase/user"
	rpcUser "auth-svc/src/shared/grpc/user"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHttpHandler struct {
	service user.UserService
}

func newHttpHandler(service user.UserService) *userHttpHandler {
	h := &userHttpHandler{service: service}
	if h.service == nil {
		panic("please provide user service")
	}
	return h
}

func (h *userHttpHandler) Register(c echo.Context) error {
	request := &rpcUser.RegisterRequest{}

	if err := c.Bind(request); err != nil {
		return err
	}

	customValidator := func() error {
		v := validator.New()
		rules := map[string]string{
			"Username":        "required",
			"Email":           "required,email",
			"Password":        "required",
			"ConfirmPassword": "required,eqfield=Password",
		}

		v.RegisterStructValidationMapRules(
			rules,
			request,
		)
		return v.Struct(request)
	}

	if err := customValidator(); err != nil {
		return err
	}

	//if err := c.Validate(request); err != nil {
	//	return err
	//}

	resp, err := h.service.Register(context.Background(), request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *userHttpHandler) Login(c echo.Context) error {
	request := &rpcUser.LoginRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}

	customValidator := func() error {
		v := validator.New()
		rules := map[string]string{
			"Email":    "required,email",
			"Password": "required",
		}

		v.RegisterStructValidationMapRules(
			rules,
			request,
		)
		return v.Struct(request)
	}

	if err := customValidator(); err != nil {
		return err
	}
	resp, err := h.service.Login(context.Background(), request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *userHttpHandler) Refresh(c echo.Context) error {
	request := &rpcUser.RefreshRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}

	customValidator := func() error {
		v := validator.New()
		rules := map[string]string{
			"RefreshToken": "required",
		}

		v.RegisterStructValidationMapRules(
			rules,
			request,
		)
		return v.Struct(request)
	}

	if err := customValidator(); err != nil {
		return err
	}
	resp, err := h.service.Refresh(context.Background(), request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *userHttpHandler) User(c echo.Context) error {
	request := &rpcUser.UserRequest{}
	request.AccessToken = c.Request().Header.Get("Authorization")
	if err := c.Bind(request); err != nil {
		return err
	}

	customValidator := func() error {
		v := validator.New()
		rules := map[string]string{
			"AccessToken": "required",
		}

		v.RegisterStructValidationMapRules(
			rules,
			request,
		)
		return v.Struct(request)
	}

	if err := customValidator(); err != nil {
		return err
	}
	resp, err := h.service.FindByID(context.Background(), request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
