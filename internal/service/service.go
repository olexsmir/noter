package service

import (
	"time"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/flof-ik/noter/internal/repository"
	"github.com/flof-ik/noter/pkg/cache"
	"github.com/flof-ik/noter/pkg/hash"
	"github.com/flof-ik/noter/pkg/token"
)

//go:generate mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go

type Users interface {
	SignUp(user domain.UserSignUp) error
	SignIn(input domain.UserSignIn) (domain.Tokens, error)
	RefreshTokens(refreshToken string) (domain.Tokens, error)
	Logout(userID int) error
}

type Notes interface {
	Create(input domain.Note) error
	GetByID(id int) (domain.Note, error)
	GetAll(authorID, notebookID, page int) ([]domain.Note, error)
	Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error
	Delete(id, authorID int) error
	DeleteAll(notebookID, userID int) error
}

type Notebooks interface {
	Create(input domain.Notebook) error
	GetAll(userID, page int) ([]domain.Notebook, error)
	GetById(id, userID int) (domain.Notebook, error)
	Update(id, userID int, inp domain.UpdateNotebookInput) error
	Delete(id, userID int) error
}

type Services struct {
	User     Users
	Note     Notes
	Notebook Notebooks
}

type Deps struct {
	Repos        *repository.Repositorys
	Hasher       hash.PasswordHasher
	TokenManager token.TokenManager
	Cache        cache.Cacher

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	CacheTTL        int64
}

func NewServices(deps Deps) *Services {
	return &Services{
		User:     NewUsersService(deps.Repos.User, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		Note:     NewNotesService(deps.Repos.Note, deps.Cache, deps.CacheTTL),
		Notebook: NewNotebooksSerivce(deps.Repos.Notebook),
	}
}
