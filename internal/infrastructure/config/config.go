package config

import (
	"encoding/json"
	"exchange-rates-api/internal/adapters"
	"fmt"
	"strings"

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

func (c *Config) String() string {
	bb, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	sb.WriteByte('\n')
	sb.WriteString(strings.Repeat("~", 100))
	sb.WriteByte('\n')
	sb.WriteString("YAML configuration file:\n")
	sb.WriteString(strings.Repeat("~", 100))
	sb.WriteByte('\n')
	sb.Write(bb)
	sb.WriteByte('\n')
	sb.WriteString(strings.Repeat("~", 100))
	sb.WriteByte('\n')

	return sb.String()
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
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	return &config, nil
}
