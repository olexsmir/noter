package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
	"github.com/Smirnov-O/noter/pkg/hash"
	"github.com/Smirnov-O/noter/pkg/token"
)

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager token.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager token.TokenManager, accessTokenTTL, refreshTokenTTL time.Duration) *UsersService {
	return &UsersService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
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

func (s *UsersService) SignIn(input domain.UserSignIn) (domain.Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return domain.Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(input.Email, passwordHash)
	if err != nil {
		return domain.Tokens{}, err
	}

	return s.createSession(user.ID)
}

func (s *UsersService) createSession(userID int) (domain.Tokens, error) {
	var (
		tokens domain.Tokens
		err    error
	)

	tokens.Access, err = s.tokenManager.NewJWT(strconv.Itoa(userID), s.accessTokenTTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	tokens.Refresh, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return domain.Tokens{}, err
	}

	err = s.repo.SetSession(domain.Session{
		UserID:       userID,
		RefreshToken: tokens.Refresh,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	})

	return tokens, err
}
