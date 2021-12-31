package service

import (
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
	"github.com/Smirnov-O/noter/pkg/hash"
)

type Users interface {
	SignUp(user domain.UserSignUp) error
}

type Services struct {
	User Users
}

func NewServices(repos *repository.Repositorys, hasher hash.PasswordHasher) *Services {
	return &Services{
		User: NewUsersService(repos.User, hasher),
	}
}
