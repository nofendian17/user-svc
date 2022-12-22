package user

import (
	domain "auth-svc/src/domain/user"
	psql "auth-svc/src/shared/database"
	"strings"
	"time"
)

type userRepository struct {
	database *psql.DatabaseClient
}

func NewRepository(database *psql.DatabaseClient) *userRepository {
	u := &userRepository{database: database}
	if u.database == nil {
		panic("please provide database connection.")
	}
	return u
}

func (u *userRepository) FindByID(ID int64) (d domain.User, err error) {
	qb := strings.Builder{}
	qb.WriteString("SELECT id, username, email, password, is_active, created_at, updated_at FROM users WHERE id = $1")
	err = u.database.Client.Get(&d, qb.String(), ID)
	return
}

func (u *userRepository) FindByEmail(email string) (d domain.User, err error) {
	qb := strings.Builder{}
	qb.WriteString("SELECT id, username, email, password, is_active, created_at, updated_at FROM users WHERE email = $1")
	err = u.database.Client.Get(&d, qb.String(), email)
	return
}

func (u *userRepository) FindByUsername(username string) (d domain.User, err error) {
	qb := strings.Builder{}
	qb.WriteString("SELECT id, username, email, password, is_active, created_at, updated_at FROM users WHERE active = TRUE AND username = $1")
	err = u.database.Client.Get(&d, qb.String(), username)
	return
}

func (u *userRepository) Create(d domain.User) error {
	qb := strings.Builder{}
	qb.WriteString("INSERT INTO ")
	qb.WriteString(u.database.SchemaName())
	qb.WriteString(".users ")
	qb.WriteString("(username, email, password, is_active, created_at) ")
	qb.WriteString("VALUES ")
	qb.WriteString("($1, $2, $3, $4, $5) ")
	qb.WriteString("RETURNING id")

	now := time.Now()
	err := u.database.Client.Get(&d.ID, qb.String(), d.Username, d.Email, d.Password, false, now)
	return err
}

func (u *userRepository) Update(ID int64, d domain.User) (domain.User, error) {
	qb := strings.Builder{}
	qb.WriteString("UPDATE ")
	qb.WriteString(u.database.SchemaName())
	qb.WriteString(".users ")
	qb.WriteString("SET username = $1, email = $2, password = $3, is_active = $4 ")
	qb.WriteString("WHERE id = $5 ")
	qb.WriteString("RETURNING id")

	now := time.Now()

	_, err := u.database.Client.Exec(qb.String(), d.Username, d.Email, d.Password, d.IsActive, now, ID)
	if err != nil {
		return d, err
	}

	d.ID = ID

	return d, err
}

func (u *userRepository) Delete(ID int64) error {
	qb := strings.Builder{}
	qb.WriteString("DELETE FROM ")
	qb.WriteString(u.database.SchemaName())
	qb.WriteString(" WHERE id = $1")

	_, err := u.database.Client.Exec(qb.String(), ID)
	if err != nil {
		return err
	}

	return err
}
