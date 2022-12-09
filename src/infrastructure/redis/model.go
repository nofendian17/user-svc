package redis

type CacheContainer struct {
	TokenCache *TokenCache `json:"token_cache,omitempty"`
}

type TokenCache struct {
	UserID string `json:"user_id"`
}
