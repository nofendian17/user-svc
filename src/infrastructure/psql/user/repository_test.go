package user

import (
	domain "auth-svc/src/domain/user"
	psql "auth-svc/src/shared/database"
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type repositorySuite struct {
	suite.Suite
	databaseClient psql.DatabaseClient
	mock           sqlmock.Sqlmock
	repository     *userRepository
	user           *domain.User
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(repositorySuite))
}

func (r *repositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, r.mock, err = sqlmock.New()
	if err != nil {
		require.NoError(r.T(), err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	r.databaseClient = psql.DatabaseClient{
		Client: sqlxDB,
	}
}

func (r *repositorySuite) AfterTest(_, _ string) {
	require.NoError(r.T(), r.mock.ExpectationsWereMet())
}

func (r *repositorySuite) TestValidate() {
	//scenario: panic, db master & db slave is nil
	assert.Panics(r.T(), func() {
		NewRepository(nil)
	})

	//scenario: not panic
	assert.NotPanics(r.T(), func() {
		NewRepository(&r.databaseClient)
	})
}

func (r *repositorySuite) TestFindByID() {
	repo := NewRepository(&r.databaseClient)
	//scenario: return error
	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs(0).
		WillReturnError(errors.New("data not found"))

	_, err := repo.FindByID(0)
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "test", "test@local")

	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs(1).
		WillReturnRows(rows)

	//scenario: success return expected data
	entity, err := repo.FindByID(1)
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), "test", entity.Username)
	assert.Equal(r.T(), "test@local", entity.Email)
}

func (r *repositorySuite) TestFindByEmail() {
	repo := NewRepository(&r.databaseClient)
	//scenario: return error
	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs("test@local").
		WillReturnError(errors.New("data not found"))

	_, err := repo.FindByEmail("test@local")
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "test", "test@local")

	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs("test@local").
		WillReturnRows(rows)

	//scenario: success return expected data
	entity, err := repo.FindByEmail("test@local")
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), int64(1), entity.ID)
	assert.Equal(r.T(), "test", entity.Username)
	assert.Equal(r.T(), "test@local", entity.Email)
}

func (r *repositorySuite) TestFindByUsername() {
	repo := NewRepository(&r.databaseClient)
	//scenario: return error
	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs("test").
		WillReturnError(errors.New("data not found"))

	_, err := repo.FindByUsername("test")
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "test", "test@local")

	r.mock.ExpectQuery("^SELECT (.+) FROM t_user").
		WithArgs("test").
		WillReturnRows(rows)

	//scenario: success return expected data
	entity, err := repo.FindByUsername("test")
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), int64(1), entity.ID)
	assert.Equal(r.T(), "test", entity.Username)
	assert.Equal(r.T(), "test@local", entity.Email)
}

func (r *repositorySuite) TestCreate() {
	repo := NewRepository(&r.databaseClient)
	query := "^INSERT (.+)t_user*"
	//scenario: return error
	r.mock.ExpectQuery(query).
		WithArgs("test", "test@local", "secret", false, AnyTime{}).
		WillReturnError(errors.New("data not found"))

	user := domain.User{
		Username: "test",
		Email:    "test@local",
		Password: "secret",
		IsActive: false,
	}

	err := repo.Create(user)
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")

	r.mock.ExpectQuery(query).
		WithArgs("test", "test@local", "secret", false, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	//scenario: success return expected data

	err = repo.Create(user)
	assert.Nil(r.T(), err)
}

func (r *repositorySuite) TestUpdate() {
	repo := NewRepository(&r.databaseClient)
	query := "^UPDATE (.+)t_user*"

	user := domain.User{
		Username: "test",
		Email:    "test@local",
		Password: "secret",
		IsActive: true,
	}

	r.mock.ExpectExec(query).
		WithArgs("test", "test@local", "secret", true, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// scenario: success return expected data
	d, err := repo.Update(int64(1), user)
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), int64(1), d.ID)

	// scenario: return error exec
	r.mock.ExpectExec(query).
		WithArgs("test", "test@local", "secret", true, AnyTime{}, 1).
		WillReturnError(errors.New("data not found"))

	_, err = repo.Update(int64(1), user)
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")
}

func (r *repositorySuite) TestDelete() {
	repo := NewRepository(&r.databaseClient)
	query := "^DELETE FROM *"

	r.mock.ExpectExec(query).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// scenario: success return expected data
	err := repo.Delete(int64(1))
	assert.Nil(r.T(), err)

	// scenario: return error exec
	r.mock.ExpectExec(query).
		WithArgs(1).
		WillReturnError(errors.New("data not found"))

	err = repo.Delete(int64(1))
	assert.NotNil(r.T(), err)
	assert.EqualError(r.T(), err, "data not found")

}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
