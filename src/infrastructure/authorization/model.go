package authorization

type JWTConfig struct {
	AccessSecretKey         string
	AccessExpirationMinute  int64
	RefreshSecretKey        string
	RefreshExpirationMinute int64
	Issuer                  string
}
