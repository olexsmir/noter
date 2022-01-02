package service

import (
	"github.com/Smirnov-O/noter/internal/config"
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
	"github.com/Smirnov-O/noter/pkg/hash"
	"github.com/Smirnov-O/noter/pkg/token"
)

type Users interface {
	SignUp(user domain.UserSignUp) error
	SignIn(input domain.UserSignIn) (domain.Tokens, error)
	RefreshTokens(refreshToken string) (domain.Tokens, error)
}

type Services struct {
	User Users
}

func NewServices(repos *repository.Repositorys, hasher hash.PasswordHasher, tokenManager token.TokenManager, cfg *config.Config) *Services {
	return &Services{
		User: NewUsersService(repos.User, hasher, tokenManager, cfg.Auth.JWT.AccessTokenTTL, cfg.Auth.JWT.RefreshTokenTTL),
	}
}
