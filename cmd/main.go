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
	"github.com/Smirnov-O/noter/internal/repository"
	"github.com/Smirnov-O/noter/internal/server"
	"github.com/Smirnov-O/noter/internal/service"
	"github.com/Smirnov-O/noter/internal/transport/rest"
	"github.com/Smirnov-O/noter/pkg/database"
	"github.com/Smirnov-O/noter/pkg/hash"
	"github.com/Smirnov-O/noter/pkg/logger"
	"github.com/Smirnov-O/noter/pkg/token"
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

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
  tokenManager, err := token.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
	}

	repos := repository.NewRepositorys(db)
	services := service.NewServices(repos, hasher, tokenManager, cfg)
	handlers := rest.NewHandler(services)

	// Server
	srv := server.NewServer(cfg, handlers.InitRoutes())
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
