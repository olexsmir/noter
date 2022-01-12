package domain

import (
	"errors"
	"time"
)

var (
	ErrNotebookAlreadyExists = errors.New("notebook already exists")
	ErrNotebookNotFound      = errors.New("notebook doesn't exists")
)

type Notebook struct {
	ID          int       `json:"id" db:"id"`
	AuthorID    int       `json:"author_id" db:"author_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateNotebookInput struct {
	Name        *string
	Description *string
	UpdatedAt   time.Time
}

func (v UpdateNotebookInput) Validate() error {
	if v.Name == nil && v.Description == nil {
		return errors.New("update input has no values")
	}

	return nil
}
