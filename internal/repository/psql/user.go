package psql

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/olexsmir/noter/internal/domain"
	"github.com/olexsmir/noter/pkg/database"
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

func (r *UsersRepo) GetByCredentials(email, password string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1 AND password=$2", email, password)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}

	return user, err
}

func (r *UsersRepo) GetByRefreshToken(refreshToken string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id IN (SELECT user_id FROM sessions WHERE refresh_token = $1)", refreshToken)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrSessionNotFound
	}

	return user, err
}

func (r *UsersRepo) SetSession(session domain.Session) error {
	_, err := r.db.Exec("INSERT INTO sessions (user_id, refresh_token, expires_at) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET refresh_token = $2",
		session.UserID, session.RefreshToken, session.ExpiresAt)

	return err
}

func (r *UsersRepo) RemoveSession(userID int) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE user_id = $1", userID)

	return err
}
