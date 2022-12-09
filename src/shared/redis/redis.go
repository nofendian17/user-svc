package redis

import (
	"auth-svc/src/shared/config"
	"fmt"
	"github.com/go-redis/redis/v7"
)

type CacheWrapper struct {
	RedisClient *redis.Client
}

func NewCache(config *config.CacheConfig) *CacheWrapper {
	opt := &redis.Options{
		Network:            "",
		Addr:               fmt.Sprintf("%s:%d", config.Host, config.Port),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           config.Username,
		Password:           config.Password,
		DB:                 config.Db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           config.PoolSize,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	}
	client := redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("failed to ping redis", err.Error())
		panic(err)
	}
	return &CacheWrapper{
		RedisClient: client,
	}
}
