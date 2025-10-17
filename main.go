// main.go
package main

import (
	"context"
	"log"

	"github.com/DeleMike/AIpply/api"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/DeleMike/AIpply/api/stringutil"
	"github.com/spf13/viper"
)

func main() {
	// load configurations
	LoadConfig()

	// Initialize LLM service once
	ctx := context.Background()
	apiKey := viper.GetString("api_key")
	log.Printf("API Key loaded: %v", stringutil.MaskString(apiKey))

	err := service.InitLLMService(ctx, apiKey)
	if err != nil {
		log.Fatalf("Failed to initialize LLM service: %v", err)
	}

	log.Printf("LLM Client initialized: %v", service.LLMClient != nil)

	// start server
	log.Println("🚀 Server starting up...")
	api.StartUpServer()
}
