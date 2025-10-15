// config.go
package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	Server struct {
		Port    int
		GinMode string
	}
	App struct {
		Name        string
		Environment string
	}
	Logging struct {
		Level string
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("app.name", "APP_NAME")
	viper.BindEnv("server.gin_mode", "GIN_MODE")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; fallback to ENV
			if !viper.IsSet("server.port") || !viper.IsSet("app.name") {
				panic(fmt.Errorf("config file not found and required environment variables missing: %w", err))
			}
		} else {
			// Config file exists but other error (parse error, etc.)
			panic(fmt.Errorf("error reading config file: %w", err))
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// Print loaded config for verification
	fmt.Println("Loaded Configuration:")
	fmt.Println("Server Gin Mode:", cfg.Server.GinMode)
	fmt.Println("Server port:", cfg.Server.Port)
	fmt.Println("App Name:", cfg.App.Name)
}
