package service

import (
	"github.com/flof-ik/noter/internal/domain"
	"github.com/flof-ik/noter/internal/repository"
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
		AuthorID:   input.AuthorID,
		NotebookID: input.NotebookID,
		Title:      input.Title,
		Content:    input.Content,
		CreatedAt:  input.CreatedAt,
		UpdatedAt:  input.UpdatedAt,
	})
}

func (s *NotesService) GetByID(id int) (domain.Note, error) {
	return s.repo.GetByID(id)
}

func (s *NotesService) GetAll(authorID, notebookID int) ([]domain.Note, error) {
	return s.repo.GetAll(authorID, notebookID)
}

func (s *NotesService) Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error {
	if err := inp.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, authorID, notebookID, domain.UpdateNoteInput{
		Title:     inp.Title,
		Content:   inp.Content,
		Pinted:    inp.Pinted,
		UpdatedAt: inp.UpdatedAt,
	})
}

func (s *NotesService) Delete(id, authorId int) error {
	return s.repo.Delete(id, authorId)
}
