package user

import (
	"database/sql"
)

type User struct {
	ID        int64        `db:"id"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	IsActive  bool         `db:"is_active"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
