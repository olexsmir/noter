package repository

import (
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository/psql"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user domain.User) error
	GetByCredentials(email, password string) (domain.User, error)
	GetByRefreshToken(refreshToken string) (domain.User, error)
	SetSession(session domain.Session) error
}

type Notes interface {
	Create(note domain.Note) error
	GetByID(id int) (domain.Note, error)
}

type Repositorys struct {
	User Users
	Note Notes
}

func NewRepositorys(db *sqlx.DB) *Repositorys {
	return &Repositorys{
		User: psql.NewUsersRepo(db),
		Note: psql.NewNotesRepo(db),
	}
}
