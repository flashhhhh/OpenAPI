package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	JWT JWTConfig `yaml:"jwt"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"db_name"`
}

type JWTConfig struct {
	SecretKey string `yaml:"secret_key"`
	TokenExpiry string `yaml:"token_expiry"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("configs/config_local.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *Config) GetJWTConfig() JWTConfig {
	return c.JWT
}