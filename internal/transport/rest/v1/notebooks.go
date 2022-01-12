package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initNotebooksRoutes(api *gin.RouterGroup) {
	notebooks := api.Group("/notebook")
	{
		authenticated := notebooks.Group("/", h.userIdentity)
		{
			authenticated.POST("/", h.notebooksCreate)
      authenticated.GET("/", h.notebookGetAll)
		}
	}
}

type notebooksCreateInput struct {
	Name        string `json':"name" bindings:"required,min=3"`
	Description string `json':"description"`
}

func (h *Handler) notebooksCreate(c *gin.Context) {
	var inp notebooksCreateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Notebook.Create(domain.Notebook{
		AuthorID:    userID,
		Name:        inp.Name,
		Description: inp.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}); err != nil {
		if errors.Is(err, domain.ErrNotebookAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) notebookGetAll(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	notebooks, err := h.services.Notebook.GetAll(userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotebookNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, notebooks)
}
