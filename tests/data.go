package tests

import (
	"time"

	"github.com/flof-ik/noter/internal/domain"
)

var (
	user = domain.User{
        ID: 1,
		Name:        "The test user",
		Email:       "uniqueTestEmail@test.com",
		Password:    "the strong password",
		RegistredAt: time.Now(),
		LastVisitAt: time.Now(),
	}

	notebook = domain.Notebook{
        ID: 1,
		AuthorID:    1,
		Name:        "Notebook for tests",
		Description: "The test notebook for tests",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	note = domain.Note{
        ID: 1,
		AuthorID:   1,
		NotebookID: 1,
		Pinted:     false,
		Title:      "The test note",
		Content:    "I'm create this note only for tests",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
)
