package psql

import (
	"database/sql"
	"errors"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/jmoiron/sqlx"
)

type NotesRepo struct{ db *sqlx.DB }

func NewNotesRepo(db *sqlx.DB) *NotesRepo {
	return &NotesRepo{db}
}

func (r *NotesRepo) Create(note domain.Note) error {
	_, err := r.db.Exec("INSERT INTO notes (author_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		note.AuthorID, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	return err
}

func (r *NotesRepo) GetByID(id int) (domain.Note, error) {
	var note domain.Note
	err := r.db.Get(&note, "SELECT * FROM notes WHERE id=$1", id)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Note{}, err
	}

	return note, err
}
