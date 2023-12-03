package main

import (
	"bonjour_nails/cmd/config"
	"bonjour_nails/cmd/server"
	"bonjour_nails/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	m, err := server.Run()
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}

	app.Run(cfg, m)
}
