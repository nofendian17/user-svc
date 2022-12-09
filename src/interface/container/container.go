package container

import (
	"auth-svc/src/infrastructure/authorization/jwt"
	dbUser "auth-svc/src/infrastructure/psql/user"
	"auth-svc/src/infrastructure/redis/cache"
	"auth-svc/src/interface/usecase/user"
	"auth-svc/src/shared/config"
	database "auth-svc/src/shared/database"
	"auth-svc/src/shared/redis"
)

type Container struct {
	UserService user.UserService
}

func Setup(cfg *config.DefaultConfig) *Container {
	db := database.NewDatabase(cfg.Database)
	c := redis.NewCache(cfg.Cache)

	userRepository := dbUser.NewRepository(db)
	cacheService := cache.NewCache(c)

	authorizationService := jwt.NewJWT(cfg.Jwt, cacheService)

	userService := user.NewService(cfg, userRepository, authorizationService, cacheService)
	return &Container{
		UserService: userService,
	}
}
