package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initNotebooksRoutes(api *gin.RouterGroup) {
	notebooks := api.Group("/notebook")
	{
		authenticated := notebooks.Group("/", h.userIdentity)
		{
			authenticated.GET("/", h.notebookGetAll)
			authenticated.GET("/:notebook_id", h.notebookGetById)
			authenticated.POST("/", h.notebooksCreate)
			authenticated.PUT("/:notebook_id", h.notebookUpdate)
			authenticated.DELETE("/:notebook_id", h.notebookDelete)

			notes := authenticated.Group("/:notebook_id")
			{
				h.initNotesRoutes(notes)
			}
		}
	}
}

type notebooksCreateInput struct {
	Name        string `json':"name" bindings:"required,min=3"`
	Description string `json':"description"`
}

// @Summary Create notebook
// @Security user-auth
// @Tags notebooks
// @Description create new notebook
// @Accept json
// @Produce json
// @Success 201
// @Param input body notebooksCreateInput true "create notebook input"
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook [post]
func (h *Handler) notebooksCreate(c *gin.Context) {
	var inp notebooksCreateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
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

// @Summary Get notebook by id
// @Security user-auth
// @Tags notebooks
// @Description get notebook by id
// @Accept json
// @Produce json
// @Param id path string true "notebook_id"
// @Success 200 {object} domain.Notebook
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{id} [get]
func (h *Handler) notebookGetById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

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

// @Summary Get all notebooks
// @Security user-auth
// @Tags notebooks
// @Description get all notebooks
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Notebook
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook [get]
func (h *Handler) notebookGetAll(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
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

type notebookUpdateInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// @Summary Update notebook
// @Security user-auth
// @Tags notebooks
// @Description update notebook
// @Accept json
// @Produce json
// @Param id path string true "notebook_id"
// @Param input body notebookUpdateInput true "update info"
// @Success 200
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{id} [put]
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

	userID, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "userCtx is of invalid type")
		return
	}

	if err := h.services.Notebook.Update(id, userID, domain.UpdateNotebookInput{
		Name:        inp.Name,
		Description: inp.Description,
		UpdatedAt:   time.Now(),
	}); err != nil {
		if errors.Is(err, domain.ErrNotebookNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Update notebook
// @Security user-auth
// @Tags notebooks
// @Description delete notebook
// @Accept json
// @Produce json
// @Param id path string true "notebook_id"
// @Success 200
// @Failure 400,401,404,500 {object} response
// @Failure default {object} response
// @Router /notebook/{id} [delete]
func (h *Handler) notebookDelete(c *gin.Context) {
	idParam := c.Param("notebook_id")
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

	if err := h.services.Notebook.DeleteAllNotes(id, userID); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Notebook.Delete(id, userID); err != nil {
		if errors.Is(err, domain.ErrNotebookNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
