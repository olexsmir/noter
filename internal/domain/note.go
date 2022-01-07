package domain

import (
	"errors"
	"time"
)

var ErrNoteNotFound = errors.New("note doesn't exists")

type Note struct {
	ID        int       `json:"id" db:"id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
