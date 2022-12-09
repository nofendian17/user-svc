package authorization

type TokenDetail struct {
	AccessToken         string `json:"access_token"`
	AccessUUID          string `json:"access_uuid"`
	AccessTokenExpires  int64  `json:"access_token_expires"`
	RefreshToken        string `json:"refresh_token"`
	RefreshUUID         string `json:"refresh_uuid"`
	RefreshTokenExpires int64  `json:"refresh_token_expires"`
}

type AccessDetail struct {
	AccessUUID  string `json:"access_uuid"`
	RefreshUUID string `json:"refresh_uuid"`
	UserID      int64  `json:"user_id"`
}
