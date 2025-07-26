package main

import (
	"log"

	"github.com/shshimamo/knowledge/main/app/config"
	"github.com/shshimamo/knowledge/main/app/server"
)

func main() {
	cfg := config.LoadConfig()
	
	if err := server.Start(cfg); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}