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

// @Summary Create note in notebook
// @Security user-auth
// @Tags notes
// @Description create new note
// @Accept json
// @Produce json
// @Success 201
// @Param input body noteCreateInput true "create note input"
// @Param notebook_id path string true "notebook_id"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{notebook_id}/note [post]
func (h *Handler) noteCreate(c *gin.Context) {
	var inp noteCreateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

	notebookID, err := getNotebookID(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "notebookCtx is of invalid type")
		return
	}

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

// @Summary Get note from notebook
// @Security user-auth
// @Tags notes
// @Description get a note by id
// @Accept json
// @Produce json
// @Success 200 {object} domain.Note
// @Param notebook_id path string true "notebook_id"
// @Param id path string true "id"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{notebook_id}/note/{id} [get]
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

// @Summary Get all notes
// @Security user-auth
// @Tags notes
// @Description get all note from notebook
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Note
// @Param notebook_id path string true "notebook_id"
// @Param id path string true "id"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{notebook_id}/note [get]
func (h *Handler) noteGetAll(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

	notebookID, err := getNotebookID(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "notebookCtx is of invalid type")
		return
	}

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

// @Summary Update note
// @Security user-auth
// @Tags notes
// @Description update note from notebook by id
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Note
// @Param notebook_id path string true "notebook_id"
// @Param id path string true "id"
// @Param input body noteUpdateInput  true "update info"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{notebook_id}/note/{id} [put]
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

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

	notebookID, err := getNotebookID(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "notebookCtx is of invalid type")
		return
	}

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

// @Summary Delete note
// @Security user-auth
// @Tags notes
// @Description delete note from notebook by id
// @Accept json
// @Produce json
// @Success 200
// @Param notebook_id path string true "notebook_id"
// @Param id path string true "id"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{notebook_id}/note/{id} [delete]
func (h *Handler) noteDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

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
