package authorization

import (
	domain "auth-svc/src/domain/user"
	cacheContainer "auth-svc/src/infrastructure/redis"
	"context"
	"github.com/dgrijalva/jwt-go"
)

type AuthorizationService interface {
	GenerateToken(domain.User) (*TokenDetail, error)
	ExtractToken(signedToken string) string
	VerifyAccessToken(signedToken string) (*jwt.Token, error)
	VerifyRefreshToken(signedToken string) (*jwt.Token, error)
	AccessTokenValid(signedToken string) error
	RefreshTokenValid(signedToken string) error
	ExtractAccessTokenMetaData(signedToken string) (*AccessDetail, error)
	ExtractRefreshTokenMetaData(signedToken string) (*AccessDetail, error)
	SaveMetaData(ctx context.Context, userID string, tokenDetail *TokenDetail) error
	DeleteMetaData(ctx context.Context, accessDetail *AccessDetail) error
	GetMetaData(ctx context.Context, accessDetail *AccessDetail) (*cacheContainer.CacheContainer, error)
}
