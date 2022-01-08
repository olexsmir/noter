package service

import (
	"time"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/repository"
)

type NotesService struct {
	repo repository.Notes
}

func NewNotesService(repo repository.Notes) *NotesService {
	return &NotesService{
		repo: repo,
	}
}

func (s *NotesService) Create(input domain.Note) error {
	return s.repo.Create(domain.Note{
		AuthorID:  input.AuthorID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (s *NotesService) GetByID(id int) (domain.Note, error) {
	return s.repo.GetByID(id)
}

func (s *NotesService) GetAll(authorID int) ([]domain.Note, error) {
	return s.repo.GetAll(authorID)
}

func (s *NotesService) Delete(id, authorId int) error {
	return s.repo.Delete(id, authorId)
}
