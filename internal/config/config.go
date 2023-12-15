package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env            string        `yaml:"env" env-default:"local"`          // текущее окружение: local, dev, prod и т.п.
	StoragePath    string        `yaml:"storage_path" env-required:"true"` // SQLite, поэтому нужно указать путь до файла, где хранится наша БД
	GRPC           GRPCConfig    `yaml:"grpc"`                             // порт gRPC-сервиса и таймаут обработки запросов
	MigrationsPath string        // путь до директории с миграциями БД. Он будет использоваться утилитой migrator
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"` // время жизни выдаваемых токенов авторизации
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
