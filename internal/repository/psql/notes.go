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
	_, err := r.db.Exec("INSERT INTO notes (author_id, notebook_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		note.AuthorID, note.NotebookID, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	return err
}

func (r *NotesRepo) GetByID(id int) (domain.Note, error) {
	var note domain.Note
	err := r.db.Get(&note, "SELECT * FROM notes WHERE id=$1", id)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Note{}, domain.ErrNoteNotFound
	}

	return note, err
}

func (r *NotesRepo) GetAll(authorID, notebookID int) ([]domain.Note, error) {
	var notes []domain.Note
	err := r.db.Select(&notes, "SELECT * FROM notes WHERE author_id=$1 AND notebook_id=$2",
		authorID, notebookID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNoteNotFound
	}

	return notes, err
}

func (r *NotesRepo) Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error {
	_, err := r.db.Exec(`UPDATE notes SET
           title = COALESCE($1, title),
           content = COALESCE($2, content),
           updated_at = $3
         WHERE id=$4 AND author_id=$5 AND notebook_id=$6`,
		inp.Title, inp.Content, inp.UpdatedAt, id, authorID, notebookID)

	return err
}

func (r *NotesRepo) Delete(id, authorID int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id=$1 AND author_id=$2", id, authorID)
	if err != nil {
		return err
	}

	return err
}
