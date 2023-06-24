package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

type Database struct {
	db *pgxpool.Pool
}

var (
	dbInstance *pgxpool.Pool
	dbOnce     sync.Once
)

func NewDatabase(ctx context.Context, url string) (*pgxpool.Pool, error) {
	dbOnce.Do(func() {
		db, err := pgxpool.New(ctx, url)
		if err != nil {
			log.Fatalf("unable to create connection pool: %v", err)
		}

		dbInstance = db
	})

	return dbInstance, nil
}

func (db *Database) Ping(ctx context.Context) error {
	return db.db.Ping(ctx)
}

func (db *Database) Close() {
	db.db.Close()
}
