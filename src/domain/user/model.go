package user

import "time"

type UserModel struct {
	ID        int64
	Username  string
	Email     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
