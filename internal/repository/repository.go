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
	GetAll(authorID int) ([]domain.Note, error)
	Update(id, authorID int, inp domain.UpdateNoteInput) error
	Delete(id, authorID int) error
}

type Notebooks interface {
	Create(notebook domain.Notebook) error
  GetAll(userId int) ([]domain.Notebook, error)
}

type Repositorys struct {
	User     Users
	Note     Notes
	Notebook Notebooks
}

func NewRepositorys(db *sqlx.DB) *Repositorys {
	return &Repositorys{
		User:     psql.NewUsersRepo(db),
		Note:     psql.NewNotesRepo(db),
		Notebook: psql.NewNotebooksRepo(db),
	}
}
