package config

import "os"

type Config struct {
	Postgres  PostgresConfig
	Redis     RedisConfig
	JWTSecret string
}

type RedisConfig struct {
	Addr string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	return &Config{
		Postgres: PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
		Redis: RedisConfig{
			Addr: os.Getenv("REDIS_ADDR"),
		},
	}
}
