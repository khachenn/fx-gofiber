package config

import "fmt"

type APIConfig struct {
	Port      string `mapstructure:"PORT" validate:"required"`
	LogLevel  string `mapstructure:"LOG_LEVEL" validate:"required"`
	JWTSecret string `mapstructure:"JWT_SECRET" validate:"required"`
	MongoURI  string `mapstructure:"MONGO_URI" validate:"required"`
}

func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	AutoLoadConfig(config)
	return config
}

func (cfg APIConfig) ListenAddr() string {
	return fmt.Sprintf(":%s", cfg.Port)
}
