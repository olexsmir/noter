package psql

import (
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/pkg/database"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct{ db *sqlx.DB }

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (r *UsersRepo) Create(user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password, registred_at, last_visit_at) VALUES ($1, $2, $3, $4, $5)",
		user.Name, user.Email, user.Password, user.RegistredAt, user.LastVisitAt)

	if database.IsDuplicate(err) {
		return domain.ErrUserAlreadyExists
	}

	return err
}
