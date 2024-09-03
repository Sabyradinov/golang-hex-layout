package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Configs application configs
type Configs struct {
	DB            DatabaseConfig
	Cache         Redis
	StopTimeoutMS int
	Logger        LoggerConfig
	WebServer     WebServerConfig
	MarketUrls    MarketUrls
	SwaggerUI     SwaggerConfig
}

// Redis cache client configs
type Redis struct {
	Host           string
	Port           int
	DatabaseNumber int
	Password       string
	Keys           struct {
		RedisKey RedisKey
	}
}

// GetHostPort returns <Host>:<Port>
func (r Redis) GetHostPort() string {
	return fmt.Sprintf("%v:%v", r.Host, r.Port)
}

// RedisKey model for redis key
type RedisKey struct {
	Key        string        `json:"key"`
	Expiration time.Duration `json:"expiration"`
}

// DatabaseConfig db connection configs
type DatabaseConfig struct {
	ConnectionString string
	Name             string
	LogMode          bool
}

// LoggerConfig logger config
type LoggerConfig struct {
	Component string
	MinLevel  string
	Writer    struct {
		Brokers []string
		Topic   string
	}
}

// MarketUrls config for market service
type MarketUrls struct {
	BaseAddress UrlConfig
	Pay         string
}

// ShipmentUrls config for shipment service
type ShipmentUrls struct {
	BaseAddress UrlConfig
	Pay         string
}

// UrlConfig base url configs
type UrlConfig struct {
	Url         string
	HttpTimeout time.Duration
}

// SwaggerConfig swagger configs
type SwaggerConfig struct {
	PageTitle   string
	Host        string
	Description string
}

// WebServerConfig server configs
type WebServerConfig struct {
	Port int
	GIN  GINConfig
}

// GINConfig gin configs
type GINConfig struct {
	ReleaseMode bool
	UseLogger   bool
	UseRecovery bool
}

// Init function to read the config file
func Init(path string) (config *Configs, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config file reading error, err : %v", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("config file unmarshalling error, err : %v", err)
	}

	return
}
