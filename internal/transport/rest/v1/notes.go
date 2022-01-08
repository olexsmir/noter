package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initNotesRoutes(api *gin.RouterGroup) {
	notes := api.Group("/note")
	{
		authenticated := notes.Group("/", h.userIdentity)
		{
			authenticated.POST("/", h.noteCreate)
			authenticated.GET("/", h.noteGetAll)
			authenticated.GET("/:id", h.noteGetByID)
			authenticated.DELETE("/:id", h.noteDelete)
		}
	}
}

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

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Note.Create(domain.Note{
		AuthorID:  userID,
		Title:     inp.Title,
		Content:   inp.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
		AuthorID:  note.AuthorID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	})
}

func (h *Handler) noteGetAll(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	notes, err := h.services.Note.GetAll(userID)
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

func (h *Handler) noteDelete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Note.Delete(id); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
