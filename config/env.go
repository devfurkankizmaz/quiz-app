package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	RedisAddress   string `mapstructure:"REDIS_ADDRESS"`
	RedisPass      string `mapstructure:"REDIS_PASS"`
	RedisDb        int    `mapstructure:"REDIS_DB"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file config.yaml : ", err)
		return &env, err
	}
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return &env, err
	}
	return &env, nil
}
