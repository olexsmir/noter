package repository

import (
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository/psql"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user domain.User) error
}

type Repositorys struct {
	User Users
}

func NewRepositorys(db *sqlx.DB) *Repositorys {
	return &Repositorys{
		User: psql.NewUsersRepo(db),
	}
}
