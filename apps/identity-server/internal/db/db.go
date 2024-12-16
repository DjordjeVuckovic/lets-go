package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Database struct {
	DB  *sqlx.DB
	cfg *Config
}

func NewDatabase(c *Config) (*Database, error) {
	db, err := sqlx.Open("postgres", c.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.ConnectionPoolSize)
	db.SetMaxIdleConns(c.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(time.Duration(c.ConnectionMaxLifetimeSec) * time.Second)

	return &Database{DB: db, cfg: c}, nil
}

func (d *Database) PingCtx(ctx context.Context) error {
	return d.DB.PingContext(ctx)
}
