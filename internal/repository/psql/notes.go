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
		return domain.Note{}, domain.ErrNoteNotFound
	}

	return note, err
}

func (r *NotesRepo) GetAll(authorID int) ([]domain.Note, error) {
	var notes []domain.Note
	err := r.db.Select(&notes, "SELECT * FROM notes WHERE author_id=$1", authorID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNoteNotFound
	}

	return notes, err
}

func (r *NotesRepo) Update(id, authorID int, inp domain.UpdateNoteInput) error {
	_, err := r.db.Exec(
		`UPDATE notes SET
       title = COALESCE($1, title),
       content = COALESCE($2, content),
       updated_at = $3
     WHERE id=$4 AND author_id=$5`,
		inp.Title, inp.Content, inp.UpdatedAt, id, authorID)

	return err
}

func (r *NotesRepo) Delete(id, authorID int) error {
	res, err := r.db.Exec("DELETE FROM notes WHERE id=$1 AND author_id=$2", id, authorID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return domain.ErrNoteNotFound
	}

	return err
}
