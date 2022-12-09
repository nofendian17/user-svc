package user

import (
	rpcUser "auth-svc/src/shared/grpc/user"
	"context"
)

type UserService interface {
	Login(ctx context.Context, request *rpcUser.LoginRequest) (*rpcUser.LoginResponse, error)
	Register(ctx context.Context, request *rpcUser.RegisterRequest) (*rpcUser.RegisterResponse, error)
	Refresh(ctx context.Context, request *rpcUser.RefreshRequest) (*rpcUser.RefreshResponse, error)
	FindByID(ctx context.Context, request *rpcUser.UserRequest) (*rpcUser.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (*rpcUser.UserResponse, error)
}
