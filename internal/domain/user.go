package domain

import (
	"errors"
	"time"
)

var (
	ErrUserAlreadyExists = errors.New("user already exeists")
	ErrUserNotFound      = errors.New("user doesn't exists")
)

type User struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	RegistredAt time.Time `json:"registred_at" db:"registred_at"`
	LastVisitAt time.Time `json:"last_visit_at" db:"last_visit_at"`
}

type UserSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
