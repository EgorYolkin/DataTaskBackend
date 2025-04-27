package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type configType string

const (
	ymlConfigType configType = "yaml"
)

// NewConfig Config provider
func NewConfig(envFilePath, ymlFilePath, ymlFileName string) (*Config, error) {
	if err := godotenv.Load(envFilePath); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	viper.SetConfigType(string(ymlConfigType))
	viper.SetConfigName(ymlFileName)
	viper.AddConfigPath(ymlFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
