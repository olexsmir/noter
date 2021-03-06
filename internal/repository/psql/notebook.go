package psql

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/olexsmir/noter/internal/domain"
	"github.com/olexsmir/noter/pkg/database"
)

type NotebooksRepo struct {
	db       *sqlx.DB
	pageSize int
}

func NewNotebooksRepo(db *sqlx.DB, pageSize int) *NotebooksRepo {
	return &NotebooksRepo{
		db:       db,
		pageSize: pageSize,
	}
}

func (r *NotebooksRepo) Create(notebook domain.Notebook) error {
	_, err := r.db.Exec(`INSERT INTO notebooks (author_id, name, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		notebook.AuthorID, notebook.Name, notebook.Description, notebook.CreatedAt, notebook.UpdatedAt)

	if database.IsDuplicate(err) {
		return domain.ErrNotebookAlreadyExists
	}

	return err
}

func (r *NotebooksRepo) GetAll(authorID, pageNumber int) ([]domain.Notebook, error) {
	offset := (pageNumber * r.pageSize) - r.pageSize

	var notebooks []domain.Notebook
	err := r.db.Select(&notebooks, `
		SELECT * FROM notebooks
		WHERE author_id=$1
		ORDER BY id ASC
		LIMIT $2 OFFSET $3`,
		authorID, r.pageSize, offset)

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
	res, err := r.db.Exec(`UPDATE notebooks SET
      name = COALESCE($1, name),
      description = COALESCE($2, description),
      updated_at = $3
    WHERE id=$4 AND author_id=$5`,
		inp.Name, inp.Description, inp.UpdatedAt, id, authorID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return domain.ErrNotebookNotFound
	}

	return err
}

func (r *NotebooksRepo) Delete(id, authorID int) error {
	res, err := r.db.Exec("DELETE FROM notebooks WHERE id=$1 AND author_id=$2", id, authorID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return domain.ErrNotebookNotFound
	}

	return err
}
