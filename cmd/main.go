package main

import (
	"github.com/Smirnov-O/noter/internal/config"
	"github.com/Smirnov-O/noter/pkg/database"
)

func main() {
	cfg, err := config.New("configs")
	if err != nil {
		panic(err)
	}

	_, err = database.NewConnection(database.ConnInfo{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		panic(err)
	}
}
