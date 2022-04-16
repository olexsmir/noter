package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/olexsmir/noter/internal/config"
	"github.com/olexsmir/noter/internal/repository"
	"github.com/olexsmir/noter/internal/service"
	"github.com/olexsmir/noter/internal/transport/rest"
	"github.com/olexsmir/noter/pkg/cache"
	"github.com/olexsmir/noter/pkg/database"
	"github.com/olexsmir/noter/pkg/hash"
	"github.com/olexsmir/noter/pkg/logger"
	"github.com/olexsmir/noter/pkg/token"
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

	memCache := cache.NewMemoryCache()
	repos := repository.NewRepositorys(db, cfg.Pagination.PageSize)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		Cache:           memCache,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
		CacheTTL:        int64(cfg.CacheTTL.Seconds()),
	})
	handlers := rest.NewHandler(services, tokenManager)

	// Server
	srv := rest.NewServer(cfg, handlers.InitRoutes(cfg))
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
