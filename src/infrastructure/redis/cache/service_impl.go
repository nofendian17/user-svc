package cache

import (
	cacheContainer "auth-svc/src/infrastructure/redis"
	"auth-svc/src/shared/redis"
	"context"
	"encoding/json"
	"time"
)

type cacheService struct {
	redis *redis.CacheWrapper
}

func NewCache(redis *redis.CacheWrapper) *cacheService {
	c := &cacheService{
		redis: redis,
	}

	if c.redis == nil {
		panic("please provide redis connection.")
	}

	return c
}

func (c *cacheService) Set(ctx context.Context, key string, cache cacheContainer.CacheContainer, ttl int64) error {
	now := time.Now()
	t := time.Now().Add(time.Minute * time.Duration(ttl)).Unix()
	expire := time.Unix(t, 0)

	bytes, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	err = c.redis.RedisClient.WithContext(ctx).
		Set(key, bytes, expire.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheService) Get(ctx context.Context, key string) (*cacheContainer.CacheContainer, error) {
	var (
		cache *cacheContainer.CacheContainer
	)
	res, err := c.redis.RedisClient.WithContext(ctx).Get(key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &cache)
	if err != nil {
		return nil, err
	}
	return cache, err
}

func (c *cacheService) Delete(ctx context.Context, key string) error {
	err := c.redis.RedisClient.WithContext(ctx).Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
