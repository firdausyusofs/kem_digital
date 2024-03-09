package main

import (
	"firdausyusofs/kem_digital/config"
	"log"
)

func main() {
	// Load config from config.yml
	cfgFile, err := config.LoadConfig("./config/config")
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	// Parse config into Config struct
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}

	log.Printf("Config: %+v", cfg)
}
