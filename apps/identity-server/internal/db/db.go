package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(c *Config) (*Database, error) {
	db, err := sqlx.Open("postgres", c.ConnectionString)
	if err != nil {
		return nil, err
	}

	// Set up connection pooling
	db.SetMaxOpenConns(c.ConnectionPoolSize)
	db.SetMaxIdleConns(c.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(time.Duration(c.ConnectionMaxLifetimeSec) * time.Second)

	return &Database{DB: db}, nil
}
