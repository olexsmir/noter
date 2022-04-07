package psql

import (
	"database/sql"
	"errors"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/flof-ik/noter/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type NotesRepo struct {
	db       *sqlx.DB
	pageSize int
}

func NewNotesRepo(db *sqlx.DB, pageSize int) *NotesRepo {
	return &NotesRepo{
		db:       db,
		pageSize: pageSize,
	}
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

func (r *NotesRepo) GetAll(authorID, notebookID, pageNumber int) ([]domain.Note, error) {
	offset := (pageNumber * r.pageSize) - r.pageSize
	logger.Error(offset)

	var notes []domain.Note
	err := r.db.Select(&notes, `
		SELECT * FROM notes
		WHERE author_id=$1 AND notebook_id=$2
		ORDER BY id ASC
		LIMIT $3 OFFSET $4`,
		authorID, notebookID, r.pageSize, offset)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNoteNotFound
	}

	return notes, err
}

func (r *NotesRepo) Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error {
	res, err := r.db.Exec(`
        UPDATE notes SET
            title = COALESCE($1, title),
            content = COALESCE($2, content),
            pinted = COALESCE($3, pinted),
            updated_at = $4
        WHERE id=$5 AND author_id=$6 AND notebook_id=$7`,
		inp.Title, inp.Content, inp.Pinted, inp.UpdatedAt, id, authorID, notebookID)
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

func (r *NotesRepo) DeleteAll(notebookID, authorID int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE notebook_id=$1 AND author_id=$2",
		notebookID, authorID)

	return err
}
