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

// Register
// @Summary     Register client
// @Description Register client with email, username, password, retype_password
// @Accept application/json
// @Tags        User
// @Produce     json
// @Param  Register body rpcUser.RegisterRequest true "register request"
// @Success     200 {object} rpcUser.RegisterResponse
// @Router      /v1/user/register [post]
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

// Login
// @Summary     Login client
// @Description Auth client with email and password
// @Accept application/json
// @Tags        User
// @Produce     json
// @Param  Login body rpcUser.LoginRequest true "register request"
// @Success     200 {object} rpcUser.LoginResponse
// @Router      /v1/user/auth [post]
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

// Refresh
// @Summary     RefreshToken client
// @Description RefreshToken Auth client
// @Accept application/json
// @Tags        User
// @Produce     json
// @Param  Refresh body rpcUser.RefreshRequest true "refresh token request"
// @Success     200 {object} rpcUser.RefreshResponse
// @Router      /v1/user/refresh [post]
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

// User
// @Summary     Get auth detail
// @Description Get Auth client
// @Accept application/json
// @Tags        User
// @Produce     json
// @Param Authorization header string true "With the bearer access_token"
// @Security BearerAuth
// @Success     200 {object} rpcUser.UserResponse
// @Router      /v1/user/me [get]
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
