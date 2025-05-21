package config

import (
	"DataTask/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
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

	_ = viper.BindEnv("database.DB_HOST", "DB_HOST")
	_ = viper.BindEnv("database.DB_PORT", "DB_PORT")
	_ = viper.BindEnv("database.DB_USER", "DB_USER")
	_ = viper.BindEnv("database.DB_PASS", "DB_PASS")
	_ = viper.BindEnv("database.DB_BASE", "DB_BASE")

	_ = viper.BindEnv("rabbitmq.RABBITMQ_HOST", "RABBITMQ_HOST")
	_ = viper.BindEnv("rabbitmq.RABBITMQ_PORT", "RABBITMQ_PORT")
	_ = viper.BindEnv("rabbitmq.RABBITMQ_USER", "RABBITMQ_USER")
	_ = viper.BindEnv("rabbitmq.RABBITMQ_PASS", "RABBITMQ_PASS")

	_ = viper.BindEnv("jwt.JWT_SECRET", "JWT_SECRET")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	logger.Log.Info(cfg.JWT.Secret)
	logger.Log.Info(fmt.Sprintf("%+v", cfg.RabbitMQ.Host))

	return &cfg, nil
}
