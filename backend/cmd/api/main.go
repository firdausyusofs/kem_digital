package main

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/server"
	"firdausyusofs/kem_digital/pkg/db/mysql"
	"firdausyusofs/kem_digital/pkg/logger"
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

	// Initialize logger
	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	// Connect to MySQL
	db, err := mysql.NewMySQLDB(cfg)
	if err != nil {
		logger.Fatalf("Error connecting to MySQL: %s", err)
	} else {
		logger.Infof("Connected to MySQL")
	}

	// Start server
	s := server.NewServer(cfg, db, logger)
	if err := s.Start(); err != nil {
		logger.Fatalf("Error starting server: %s", err)
	}
}
