package jwt

import (
	domain "auth-svc/src/domain/user"
	"auth-svc/src/infrastructure/authorization"
	cacheContainer "auth-svc/src/infrastructure/redis"
	"auth-svc/src/shared/config"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type authorizationService struct {
	authorization.JWTConfig
	cacheRepository cacheContainer.CacheRepository
}

func NewJWT(cfg *config.JWTConfig, cacheRepository cacheContainer.CacheRepository) *authorizationService {
	return &authorizationService{
		JWTConfig: authorization.JWTConfig{
			AccessSecretKey:         cfg.AccessSecretKey,
			AccessExpirationMinute:  cfg.AccessExpirationMinute,
			RefreshSecretKey:        cfg.RefreshSecretKey,
			RefreshExpirationMinute: cfg.RefreshExpirationMinute,
			Issuer:                  cfg.Issuer,
		},
		cacheRepository: cacheRepository,
	}
}

func (s *authorizationService) GenerateToken(user domain.User) (*authorization.TokenDetail, error) {
	var (
		err error
	)
	tokenDetail := &authorization.TokenDetail{}
	tokenDetail.AccessTokenExpires = s.JWTConfig.AccessExpirationMinute
	tokenDetail.AccessUUID = uuid.Must(uuid.NewRandom()).String()
	tokenDetail.RefreshTokenExpires = s.JWTConfig.RefreshExpirationMinute
	tokenDetail.RefreshUUID = uuid.Must(uuid.NewRandom()).String()

	// Access Token
	accessTokenClaim := jwt.MapClaims{}
	accessTokenClaim["accessTokenUUID"] = tokenDetail.AccessUUID
	accessTokenClaim["userID"] = user.ID
	accessTokenClaim["expireAt"] = tokenDetail.AccessTokenExpires
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(s.AccessSecretKey))
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshTokenClaim := jwt.MapClaims{}
	refreshTokenClaim["refreshTokenUUID"] = tokenDetail.RefreshUUID
	refreshTokenClaim["userID"] = user.ID
	refreshTokenClaim["expireAt"] = tokenDetail.RefreshTokenExpires
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)
	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(s.RefreshSecretKey))
	if err != nil {
		return nil, err
	}
	return tokenDetail, err
}

func (s *authorizationService) VerifyAccessToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.JWTConfig.AccessSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authorizationService) VerifyRefreshToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.JWTConfig.RefreshSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authorizationService) ExtractToken(signedToken string) string {
	strArr := strings.Split(signedToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return signedToken
}

func (s *authorizationService) AccessTokenValid(signedToken string) error {
	jwtToken, err := s.VerifyAccessToken(signedToken)
	if err != nil {
		return err
	}
	if _, ok := jwtToken.Claims.(jwt.Claims); !ok || !jwtToken.Valid {
		return err
	}
	return nil
}

func (s *authorizationService) RefreshTokenValid(signedToken string) error {
	jwtToken, err := s.VerifyRefreshToken(signedToken)
	if err != nil {
		return err
	}
	if _, ok := jwtToken.Claims.(jwt.Claims); !ok || !jwtToken.Valid {
		return err
	}
	return nil
}

func (s *authorizationService) ExtractAccessTokenMetaData(signedToken string) (*authorization.AccessDetail, error) {
	jwtToken, err := s.VerifyAccessToken(signedToken)
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		accessUuid, ok := claims["accessTokenUUID"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["userID"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &authorization.AccessDetail{
			AccessUUID:  accessUuid,
			RefreshUUID: "",
			UserID:      userId,
		}, nil
	}
	return nil, err
}

func (s *authorizationService) ExtractRefreshTokenMetaData(signedToken string) (*authorization.AccessDetail, error) {
	jwtToken, err := s.VerifyRefreshToken(signedToken)
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		refreshUUID, ok := claims["refreshTokenUUID"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["userID"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &authorization.AccessDetail{
			AccessUUID:  "",
			RefreshUUID: refreshUUID,
			UserID:      userId,
		}, nil
	}
	return nil, err
}

func (s *authorizationService) SaveMetaData(ctx context.Context, userID string, tokenDetail *authorization.TokenDetail) error {
	var (
		err error
	)

	cache := cacheContainer.CacheContainer{
		TokenCache: &cacheContainer.TokenCache{
			UserID: userID,
		},
	}

	err = s.cacheRepository.Set(ctx, tokenDetail.AccessUUID, cache, tokenDetail.AccessTokenExpires)
	if err != nil {
		return err
	}

	err = s.cacheRepository.Set(ctx, tokenDetail.RefreshUUID, cache, tokenDetail.RefreshTokenExpires)
	if err != nil {
		return err
	}

	return err
}

func (s *authorizationService) DeleteMetaData(ctx context.Context, accessDetail *authorization.AccessDetail) error {
	if accessDetail.AccessUUID != "" {
		err := s.cacheRepository.Delete(ctx, accessDetail.AccessUUID)
		if err != nil {
			return err
		}
	}

	if accessDetail.RefreshUUID != "" {
		err := s.cacheRepository.Delete(ctx, accessDetail.RefreshUUID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *authorizationService) GetMetaData(ctx context.Context, accessDetail *authorization.AccessDetail) (*cacheContainer.CacheContainer, error) {

	cache, err := s.cacheRepository.Get(ctx, accessDetail.AccessUUID)
	if err != nil {
		return cache, err
	}
	return cache, err
}
