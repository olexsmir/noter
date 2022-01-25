package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

type noteCreateInput struct {
	Title   string `json:"title" binding:"required,min=12"`
	Content string `json:"content" binding:"required,min=24"`
}

func (h *Handler) noteCreate(c *gin.Context) {
	var inp noteCreateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userID := getUserId(c)
	notebookID := getNotebookID(c)

	if err := h.services.Note.Create(domain.Note{
		AuthorID:   userID,
		Title:      inp.Title,
		Content:    inp.Content,
		NotebookID: notebookID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) noteGetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	note, err := h.services.Note.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.Note{
		ID:         note.ID,
		AuthorID:   note.AuthorID,
		NotebookID: note.NotebookID,
		Title:      note.Title,
		Content:    note.Content,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	})
}

func (h *Handler) noteGetAll(c *gin.Context) {
	userID := getUserId(c)
	notebookID := getNotebookID(c)

	notes, err := h.services.Note.GetAll(userID, notebookID)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, notes)
}

type noteUpdateInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Pinted  bool    `json:"pinted"`
}

func (h *Handler) noteUpdate(c *gin.Context) {
	var inp noteUpdateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID := getUserId(c)
	notebookID := getNotebookID(c)

	if err := h.services.Note.Update(id, userID, notebookID, domain.UpdateNoteInput{
		Title:     inp.Title,
		Content:   inp.Content,
		Pinted:    inp.Pinted,
		UpdatedAt: time.Now(),
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) noteDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID := getUserId(c)

	if err := h.services.Note.Delete(id, userID); err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
