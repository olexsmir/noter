package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flof-ik/noter/internal/config"
	"github.com/flof-ik/noter/internal/repository"
	"github.com/flof-ik/noter/internal/server"
	"github.com/flof-ik/noter/internal/service"
	"github.com/flof-ik/noter/internal/transport/rest"
	"github.com/flof-ik/noter/pkg/cache"
	"github.com/flof-ik/noter/pkg/database"
	"github.com/flof-ik/noter/pkg/hash"
	"github.com/flof-ik/noter/pkg/logger"
	"github.com/flof-ik/noter/pkg/token"
)

// @title Noter API
// @version 1.0
// @description REST API for Noter

// @host localhost:8000
// @BasePath /api/v1/

// @securityDefinitions.apikey UserAuth
// @in header
// @name Authorization

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

    _ = cache.NewMemoryCache()
	repos := repository.NewRepositorys(db)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
	})
	handlers := rest.NewHandler(services, tokenManager)

	// Server
	srv := server.NewServer(cfg, handlers.InitRoutes(cfg))
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
