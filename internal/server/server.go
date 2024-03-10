package server

import (
	"context"
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
	db     *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logger.Logger) *Server {
	return &Server{
		echo:   echo.New(),
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Start() error {
	// Create server
	server := &http.Server{
		Addr: s.cfg.Server.Port,
	}

	// Start server
	go func() {
		s.logger.Infof("Starting server on port %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting server: %s", err)
		}
	}()

	// Make handlers
	if err := s.MakeHandlers(s.echo); err != nil {
		return err
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Infof("Shutting down server")
	return s.echo.Server.Shutdown(ctx)
}
