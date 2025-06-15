package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// Priority 1: Environment variable
	configPath = os.Getenv("CONFIG_PATH")

	// Priority 2: CLI flag if ENV not set
	if configPath == "" {
		flagPath := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flagPath
		if configPath == "" {
			log.Fatal("config path is not set (use CONFIG_PATH env var or --config flag)")
		}
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	// Read YAML config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	return &cfg
}
