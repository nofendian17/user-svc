package psql

import (
	"auth-svc/src/shared/config"
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Database interface {
	Conn() (*sqlx.DB, error)
	Ping(ctx context.Context) error
	Disconnect() error
}

type DatabaseClient struct {
	Client *sqlx.DB
	Schema string
}

func NewDatabase(cfg *config.DatabaseConfig) *DatabaseClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable search_path=%s",
			cfg.Username,
			cfg.Password,
			cfg.Name,
			cfg.Host,
			cfg.Port,
			cfg.Schema))

	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("failed to ping database")
		panic(err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)

	return &DatabaseClient{
		Client: db,
		Schema: cfg.Schema,
	}
}

func (d *DatabaseClient) SchemaName() string {
	return d.Schema
}

func (d *DatabaseClient) Conn() (*sqlx.DB, error) {
	if d.Client == nil {
		return nil, errors.New("config not yet configured")
	}

	return d.Client, nil
}

func (d *DatabaseClient) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}

func (d *DatabaseClient) Disconnect() error {
	return d.Client.Close()
}
