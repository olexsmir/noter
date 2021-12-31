package service

import (
	"errors"
	"time"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
	"github.com/Smirnov-O/noter/pkg/hash"
)

type UsersService struct {
	repo   repository.Users
	hasher hash.PasswordHasher
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher) *UsersService {
	return &UsersService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *UsersService) SignUp(user domain.UserSignUp) error {
	passwordHash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	if err := s.repo.Create(domain.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    passwordHash,
		RegistredAt: time.Now(),
		LastVisitAt: time.Now(),
	}); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return err
		}
		return err
	}

	return err
}
