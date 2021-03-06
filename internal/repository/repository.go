package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olexsmir/noter/internal/domain"
	"github.com/olexsmir/noter/internal/repository/psql"
)

//go:generate mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go

type Users interface {
	Create(user domain.User) error
	GetByCredentials(email, password string) (domain.User, error)
	GetByRefreshToken(refreshToken string) (domain.User, error)
	SetSession(session domain.Session) error
	RemoveSession(userID int) error
}

type Notes interface {
	Create(note domain.Note) error
	GetByID(id int) (domain.Note, error)
	GetAll(authorID, notebookID, pageNumber int) ([]domain.Note, error)
	Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error
	Delete(id, authorID int) error
	DeleteAll(notebookID, authorID int) error
}

type Notebooks interface {
	Create(notebook domain.Notebook) error
	GetAll(authorID, pageNumber int) ([]domain.Notebook, error)
	GetById(id, authorID int) (domain.Notebook, error)
	Update(id, authorID int, inp domain.UpdateNotebookInput) error
	Delete(id, authorID int) error
}

type Repositorys struct {
	User     Users
	Note     Notes
	Notebook Notebooks
}

func NewRepositorys(db *sqlx.DB, pageSize int) *Repositorys {
	return &Repositorys{
		User:     psql.NewUsersRepo(db),
		Note:     psql.NewNotesRepo(db, pageSize),
		Notebook: psql.NewNotebooksRepo(db, pageSize),
	}
}
