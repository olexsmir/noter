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
