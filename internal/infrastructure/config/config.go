package config

import (
	"exchange-rates-api/internal/adapters"
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
}

func (s *ServerConfig) SocketAddr() string { return fmt.Sprintf("%s:%d", s.Addr, s.Port) }

type Config struct {
	Server               ServerConfig                         `mapstructure:"server"`
	OpenExchangeRatesAPI adapters.OpenExchangeRatesHTTPConfig `mapstructure:"open_exchange_rates_api"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read the config file: %w", err)
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file content: %w", err)
	}
	return &config, nil
}
