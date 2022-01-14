package v1

import (
	"errors"
	"net/http"
	"strconv"
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
			authenticated.GET("/:notebook_id", h.notebookGetById)
			authenticated.PUT("/:notebook_id", h.notebookUpdate)

			notes := notebooks.Group("/:notebook_id/note", h.userIdentity, h.setNotebookCtx)
			{
				notes.POST("/", h.noteCreate)
				notes.GET("/", h.noteGetAll)
				notes.GET("/:id", h.noteGetByID)
				notes.PUT("/:id", h.noteUpdate)
				notes.DELETE("/:id", h.noteDelete)
			}
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

	userID := getUserId(c)

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

func (h *Handler) notebookGetById(c *gin.Context) {
	userID := getUserId(c)

	idParam := c.Param("notebook_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	notebook, err := h.services.Notebook.GetById(id, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotebookNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, notebook)
}

func (h *Handler) notebookGetAll(c *gin.Context) {
	userID := getUserId(c)

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

type notebookUpdateInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (h *Handler) notebookUpdate(c *gin.Context) {
	var inp notebookUpdateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	idParam := c.Param("notebook_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID := getUserId(c)

	if err := h.services.Notebook.Update(id, userID, domain.UpdateNotebookInput{
		Name:        inp.Name,
		Description: inp.Description,
		UpdatedAt:   time.Now(),
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) notebookDelete(c *gin.Context) {
	idParam := c.Param("notebook_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID := getUserId(c)

	if err := h.services.Notebook.Delete(id, userID); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
