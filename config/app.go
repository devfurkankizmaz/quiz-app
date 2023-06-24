package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log"
)

type Application struct {
	Env   *Env
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DBUserName, env.DBUserPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := NewDatabase(context.Background(), url)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := NewRedisCache(env)
	app.Redis = r
	app.Env = env
	app.DB = db
	return *app
}
