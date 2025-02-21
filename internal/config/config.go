package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
	GRPCClient `yaml:"grpc_client"`
}

type HTTPServer struct {
	ResponseTimeout time.Duration `yaml:"response_timeout"`
}

type GRPCClient struct {
	Address    string        `yaml:"address"`
	Timeout    time.Duration `yaml:"timeout"`
	RetryCount int64         `yaml:"retry_count"`
}

func LoadConfig() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "develop"
	}

	configPath := fmt.Sprintf("configs/%s/config.yaml", appEnv)
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %s", err.Error())
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %s", err.Error())
	}

	return &cfg, nil
}
