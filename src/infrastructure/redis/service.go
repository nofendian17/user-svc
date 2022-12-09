package redis

import (
	"context"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, cache CacheContainer, ttl int64) error
	Get(ctx context.Context, key string) (*CacheContainer, error)
	Delete(ctx context.Context, key string) error
}
