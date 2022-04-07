package service

import (
	"github.com/flof-ik/noter/internal/domain"
	"github.com/flof-ik/noter/internal/repository"
	"github.com/flof-ik/noter/pkg/cache"
)

type NotesService struct {
	repo  repository.Notes
	cache cache.Cacher
	ttl   int64
}

func NewNotesService(repo repository.Notes, cache cache.Cacher, ttl int64) *NotesService {
	return &NotesService{
		repo:  repo,
		cache: cache,
		ttl:   ttl,
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
	//nolint:nilerr
	if v, err := s.cache.Get(id); err == nil {
		return v.(domain.Note), err
	}

	note, err := s.repo.GetByID(id)
	if err != nil {
		return domain.Note{}, err
	}

	err = s.cache.Set(id, note, s.ttl)

	return note, err
}

func (s *NotesService) GetAll(authorID, notebookID, page int) ([]domain.Note, error) {
	return s.repo.GetAll(authorID, notebookID, page)
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

func (s *NotesService) DeleteAll(notebookID, userID int) error {
	return s.repo.DeleteAll(notebookID, userID)
}
