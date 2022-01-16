package domain

import (
	"errors"
	"time"
)

var ErrNoteNotFound = errors.New("note doesn't exists")

type Note struct {
	ID         int       `json:"id" db:"id"`
	AuthorID   int       `json:"author_id" db:"author_id"`
	NotebookID int       `json:"notebook_id" db:"notebook_id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateNoteInput struct {
	Title     *string
	Content   *string
	UpdatedAt time.Time
}

func (v UpdateNoteInput) Validate() error {
	if v.Title == nil && v.Content == nil {
		return errors.New("update input has no values")
	}

	return nil
}
