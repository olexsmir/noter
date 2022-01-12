package psql

import (
	"database/sql"
	"errors"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/pkg/database"
	"github.com/jmoiron/sqlx"
)

type NotebooksRepo struct{ db *sqlx.DB }

func NewNotebooksRepo(db *sqlx.DB) *NotebooksRepo {
	return &NotebooksRepo{db}
}

func (r *NotebooksRepo) Create(notebook domain.Notebook) error {
	_, err := r.db.Exec(`INSERT INTO notebooks (author_id, name, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		notebook.AuthorID, notebook.Name, notebook.Description, notebook.CreatedAt, notebook.UpdatedAt)

	if database.IsDuplicate(err) {
		return domain.ErrNotebookAlreadyExists
	}

	return err
}

func (r *NotebooksRepo) GetAll(authorID int) ([]domain.Notebook, error) {
	var notebooks []domain.Notebook
	err := r.db.Select(&notebooks, "SELECT * FROM notebooks WHERE author_id=$1", authorID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotebookNotFound
	}

	return notebooks, err
}

func (r *NotebooksRepo) GetById(id, authorID int) (domain.Notebook, error) {
	var notebook domain.Notebook
	err := r.db.Get(&notebook, "SELECT * FROM notebooks WHERE id=$1 AND author_id=$2",
		id, authorID)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Notebook{}, domain.ErrNotebookNotFound
	}

	return notebook, err
}

func (r *NotebooksRepo) Update(id, authorID int, inp domain.UpdateNotebookInput) error {
	_, err := r.db.Exec(`UPDATE notebooks SET
      name = COALESCE($1, name),
      description = COALESCE($2, description),
      updated_at = $3
    WHERE id=$4 AND author_id=$5`,
		inp.Name, inp.Description, inp.UpdatedAt, id, authorID)

	return err
}
