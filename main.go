// main.go
package main

import (
	"context"
	"log"

	"github.com/DeleMike/AIpply/api"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// load configurations
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// metrics.InitRedis(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	service.InitLLMService(context.Background(), cfg.APIKey)
	gin.SetMode(cfg.Server.GinMode)

	// start server
	log.Println("🚀 Server starting up...")
	api.StartUpServer()
}
