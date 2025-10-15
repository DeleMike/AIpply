// main.go
package main

import (
	"fmt"

	"github.com/DeleMike/AIpply/api"
)

func main() {
	// load configurations
	LoadConfig()

	// start server
	api.StartUpServer()

	fmt.Println("Hello, World!")
}
