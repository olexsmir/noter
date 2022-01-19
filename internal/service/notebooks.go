package service

import (
	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
)

type NotebooksService struct {
	repo repository.Notebooks
}

func NewNotebooksSerivce(repo repository.Notebooks) *NotebooksService {
	return &NotebooksService{
		repo: repo,
	}
}

func (s *NotebooksService) Create(input domain.Notebook) error {
	return s.repo.Create(domain.Notebook{
		AuthorID:    input.AuthorID,
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	})
}

func (s *NotebooksService) GetAll(userID int) ([]domain.Notebook, error) {
	return s.repo.GetAll(userID)
}

func (s *NotebooksService) GetById(id, userID int) (domain.Notebook, error) {
	return s.repo.GetById(id, userID)
}

func (s *NotebooksService) Update(id, userID int, inp domain.UpdateNotebookInput) error {
	if err := inp.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, userID, domain.UpdateNotebookInput{
		Name:        inp.Name,
		Description: inp.Description,
		UpdatedAt:   inp.UpdatedAt,
	})
}

func (s *NotebooksService) Delete(id, userID int) error {
	return s.repo.Delete(id, userID)
}

func (s *NotebooksService) DeleteAllNotes(id, userID int) error {
    return s.repo.DeleteAllNotes(id, userID)
}
