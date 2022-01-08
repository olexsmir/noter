package service

import (
	"database/sql"
	"errors"
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
	note, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Note{}, domain.ErrNoteNotFound
		}

		return domain.Note{}, err
	}

	return note, nil
}

func (s *NotesService) GetAll(authorID int) ([]domain.Note, error) {
	return s.repo.GetAll(authorID)
}

func (s *NotesService) Delete(id int) error {
	return s.repo.Delete(id)
}
