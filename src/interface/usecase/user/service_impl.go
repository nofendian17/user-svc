package user

import (
	domain "auth-svc/src/domain/user"
	"auth-svc/src/infrastructure/authorization"
	"auth-svc/src/infrastructure/redis"
	"auth-svc/src/shared/config"
	"auth-svc/src/shared/constant"
	rpcUser "auth-svc/src/shared/grpc/user"
	"auth-svc/src/shared/helper"
	"context"
	"fmt"
	"strconv"
)

type userService struct {
	config *config.DefaultConfig
	repo   domain.Repository
	auth   authorization.AuthorizationService
	cache  redis.CacheRepository
}

func NewService(config *config.DefaultConfig, repo domain.Repository, auth authorization.AuthorizationService, cache redis.CacheRepository) UserService {
	s := &userService{
		config: config,
		repo:   repo,
		auth:   auth,
		cache:  cache,
	}
	if s.repo == nil {
		panic("user repository is nil")
	}
	if s.cache == nil {
		panic("cache repository is nil")
	}
	if s.auth == nil {
		panic("authorization service is nil")
	}

	return s
}

func (u *userService) Login(ctx context.Context, request *rpcUser.LoginRequest) (*rpcUser.LoginResponse, error) {
	user, err := u.repo.FindByEmail(request.GetEmail())
	if err != nil {
		err = helper.NewError(constant.CodeProcessingError, "user not found")
		return nil, err
	}

	match := helper.CheckPasswordHash(request.GetPassword(), user.Password)
	if !match {
		err = helper.NewError(constant.CodeProcessingError, "wrong password")
		return nil, err
	}

	token, err := u.auth.GenerateToken(user)
	if err != nil {
		err = helper.NewError(constant.CodeInternalServerError, "failed generate token")
		return nil, err
	}

	userID := fmt.Sprintf("%d", user.ID)
	err = u.auth.SaveMetaData(ctx, userID, token)
	if err != nil {
		err = helper.NewError(constant.CodeInternalServerError, "failed save authorization data to cache")
		return nil, err
	}

	res := &rpcUser.LoginResponse{
		Code:    constant.CodeSuccess,
		Message: "success",
		Data: &rpcUser.TokenData{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		},
	}
	return res, err
}

func (u *userService) Register(ctx context.Context, request *rpcUser.RegisterRequest) (*rpcUser.RegisterResponse, error) {
	var model domain.User
	model.Username = request.GetUsername()
	model.Email = request.GetEmail()
	model.Password = helper.HashPassword(request.GetPassword())

	_, err := u.repo.FindByEmail(request.GetEmail())
	if err == nil {
		err = helper.NewError(constant.CodeProcessingError, "email already exist")
		return nil, err
	}

	_, err = u.repo.FindByUsername(request.GetUsername())
	if err == nil {
		err = helper.NewError(constant.CodeProcessingError, "username already exist")
		return nil, err
	}

	err = u.repo.Create(model)
	if err != nil {
		err = helper.NewError(constant.CodeProcessingError, "failed create new user")
		return nil, err
	}

	res := &rpcUser.RegisterResponse{
		Message: constant.CodeSuccess,
		Code:    "success",
	}
	return res, err
}

func (u *userService) Refresh(ctx context.Context, request *rpcUser.RefreshRequest) (*rpcUser.RefreshResponse, error) {
	refreshToken := request.GetRefreshToken()
	metadata, err := u.auth.ExtractRefreshTokenMetaData(refreshToken)
	if err != nil {
		err = helper.NewError(constant.CodeInvalidRequest, err.Error())
		return nil, err
	}

	accessDetail := &authorization.AccessDetail{
		AccessUUID:  metadata.AccessUUID,
		RefreshUUID: metadata.RefreshUUID,
		UserID:      metadata.UserID,
	}

	err = u.auth.DeleteMetaData(ctx, accessDetail)
	if err != nil {
		err = helper.NewError(constant.CodeInvalidRequest, "unauthorized request")
		return nil, err
	}

	user, err := u.repo.FindByID(accessDetail.UserID)
	if err != nil {
		err = helper.NewError(constant.CodeDataNotFound, "user not found")
		return nil, err
	}

	token, err := u.auth.GenerateToken(user)
	if err != nil {
		err = helper.NewError(constant.CodeInvalidRequest, "unauthorized request")
		return nil, err
	}

	err = u.auth.SaveMetaData(ctx, strconv.FormatInt(user.ID, 10), token)
	if err != nil {
		err = helper.NewError(constant.CodeInternalServerError, err.Error())
		return nil, err
	}

	res := &rpcUser.RefreshResponse{
		Message: "success",
		Code:    constant.CodeSuccess,
		Data: &rpcUser.TokenData{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		},
	}

	return res, nil
}

func (u *userService) FindByID(ctx context.Context, request *rpcUser.UserRequest) (*rpcUser.UserResponse, error) {
	accessToken := u.auth.ExtractToken(request.GetAccessToken())
	metadata, err := u.auth.ExtractAccessTokenMetaData(accessToken)
	if err != nil {
		err = helper.NewError(constant.CodeInvalidRequest, "unauthorized request")
		return nil, err
	}

	_, err = u.auth.GetMetaData(ctx, metadata)
	if err != nil {
		err = helper.NewError(constant.CodeInvalidRequest, "unauthorized request")
		return nil, err
	}

	entity, err := u.repo.FindByID(metadata.UserID)
	if err != nil {
		err = helper.NewError(constant.CodeDataNotFound, "user not found")
		return nil, err
	}

	createdAt := func() string {
		if entity.CreatedAt.Valid {
			return entity.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		return ""
	}
	updatedAt := func() string {
		if entity.UpdatedAt.Valid {
			return entity.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		return ""
	}

	res := &rpcUser.UserResponse{
		Code:    constant.CodeSuccess,
		Message: "success",
		Data: &rpcUser.UserData{
			Id:        entity.ID,
			Username:  entity.Username,
			Email:     entity.Email,
			IsActive:  entity.IsActive,
			CreatedAt: createdAt(),
			UpdatedAt: updatedAt(),
		},
	}
	return res, err
}

func (u *userService) FindByEmail(ctx context.Context, email string) (*rpcUser.UserResponse, error) {
	entity, err := u.repo.FindByEmail(email)
	if err != nil {
		err = helper.NewError(constant.CodeDataNotFound, "user not found")
		return nil, err
	}

	createdAt := func() string {
		if entity.CreatedAt.Valid {
			return entity.CreatedAt.Time.GoString()
		}
		return ""
	}
	updatedAt := func() string {
		if entity.UpdatedAt.Valid {
			return entity.UpdatedAt.Time.GoString()
		}
		return ""
	}

	res := &rpcUser.UserResponse{
		Code:    constant.CodeSuccess,
		Message: "success",
		Data: &rpcUser.UserData{
			Id:        entity.ID,
			Username:  entity.Username,
			Email:     entity.Email,
			Password:  entity.Password,
			IsActive:  entity.IsActive,
			CreatedAt: createdAt(),
			UpdatedAt: updatedAt(),
		},
	}
	return res, err
}
