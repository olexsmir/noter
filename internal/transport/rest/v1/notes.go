package v1

import (
	"net/http"
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
