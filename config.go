// config.go
package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	Server struct {
		Port    int
		GinMode string `mapstructure:"gin_mode"`
	}
	App struct {
		Name        string
		Environment string
	}
	Logging struct {
		Level string
	}

	APIKey string `mapstructure:"api_key"`

	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.BindEnv("api_key", "API_KEY")
	viper.BindEnv("redis.url", "REDIS_URL")
	viper.BindEnv("redis.addr", "REDIS_ADDR")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("app.name", "APP_NAME")
	viper.BindEnv("server.gin_mode", "GIN_MODE")

	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.password", "")

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found, but that's ok, we'll rely on ENVs/defaults
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	if redisURL := viper.GetString("redis.url"); redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, fmt.Errorf("error parsing REDIS_URL: %w", err)
		}
		cfg.Redis.Addr = opt.Addr
		cfg.Redis.Password = opt.Password
		cfg.Redis.DB = opt.DB
		fmt.Println("✓ Using Redis from REDIS_URL")
	}

	// Print loaded config for verification
	fmt.Println("Loaded Configuration:")
	fmt.Println("Server Gin Mode:", cfg.Server.GinMode)
	fmt.Println("Server port:", cfg.Server.Port)
	fmt.Println("App Name:", cfg.App.Name)
	fmt.Println("Redis Addr:", cfg.Redis.Addr)

	return &cfg, nil
}
