package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Smirnov-O/noter/internal/config"
	"github.com/Smirnov-O/noter/internal/server"
	"github.com/Smirnov-O/noter/pkg/database"
	"github.com/Smirnov-O/noter/pkg/logger"
)

func main() {
	cfg, err := config.New("configs")
	if err != nil {
		logger.Error(err)
	}

	db, err := database.NewConnection(database.ConnInfo{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logger.Error(err)
	}

	// Server
	srv := server.NewServer(cfg, nil)
	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Error(err)
	}

	if err := db.Close(); err != nil {
		logger.Error(err)
	}
}
