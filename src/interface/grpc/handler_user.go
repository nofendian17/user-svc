package grpc

import (
	"auth-svc/src/interface/usecase/user"
	rpcUser "auth-svc/src/shared/grpc/user"
	"context"
	"github.com/go-playground/validator/v10"
)

type userGrpcHandler struct {
	service user.UserService
	rpcUser.UnimplementedUserServiceServer
}

func newGrpcHandler(service user.UserService) *userGrpcHandler {
	h := &userGrpcHandler{service: service}
	if h.service == nil {
		panic("please provide user service")
	}
	return h
}

func (h *userGrpcHandler) Register(ctx context.Context, request *rpcUser.RegisterRequest) (*rpcUser.RegisterResponse, error) {
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
		return nil, err
	}

	resp, err := h.service.Register(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *userGrpcHandler) Login(ctx context.Context, request *rpcUser.LoginRequest) (*rpcUser.LoginResponse, error) {
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
		return nil, err
	}

	resp, err := h.service.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (h *userGrpcHandler) Refresh(ctx context.Context, request *rpcUser.RefreshRequest) (*rpcUser.RefreshResponse, error) {
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
		return nil, err
	}
	resp, err := h.service.Refresh(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (h *userGrpcHandler) User(ctx context.Context, request *rpcUser.UserRequest) (*rpcUser.UserResponse, error) {
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
		return nil, err
	}
	resp, err := h.service.FindByID(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}
