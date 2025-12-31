package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Env string `env:"ENV" envDefault:"local"`

	HTTP HTTPConfig

	Store StoreConfig
	Redis    RedisConfig
	Elastic  ElasticConfig

	Worker WorkerConfig

	Logger LoggerConfig
}

type LoggerConfig struct {
	Level   string `env:"LOG_LEVEL" envDefault:"info"`
	Service string `env:"SERVICE_NAME" envDefault:"inventory-api"`
}

type HTTPConfig struct {
	Address         string        `env:"HTTP_ADDR" envDefault:":8080"`
	ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"10s"`
	IdleTimeout     time.Duration `env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
}

type StoreConfig struct {
	URL             string        `env:"POSTGRES_URL,required"`
	MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS" envDefault:"10"`
	ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME" envDefault:"5m"`
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

type ElasticConfig struct {
	Addresses []string `env:"ELASTIC_ADDRESSES" envSeparator:"," required:"true"`
	Username  string   `env:"ELASTIC_USERNAME"`
	Password  string   `env:"ELASTIC_PASSWORD"`
	Index     string   `env:"ELASTIC_PRODUCT_INDEX" envDefault:"products"`
}

type WorkerConfig struct {
	PollInterval time.Duration `env:"WORKER_POLL_INTERVAL" envDefault:"1s"`
	BatchSize    int           `env:"WORKER_BATCH_SIZE" envDefault:"10"`
	MaxRetries   int           `env:"WORKER_MAX_RETRIES" envDefault:"5"`
}

func Load() Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return cfg
}