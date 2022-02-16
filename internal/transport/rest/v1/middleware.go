package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader  = "Authorization"
	userCtx     = "userID"
	notebookCtx = "notebookID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userID, err := h.tokenManager.Parse(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

func (h *Handler) setNotebookCtx(c *gin.Context) {
	id := c.Param("notebook_id")

	c.Set(notebookCtx, id)
}

func getNotebookID(c *gin.Context) (int, error) {
	id := c.GetString(notebookCtx)
	notebookID, err := strconv.Atoi(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "notebookCtx is of invalid type")
		return 0, err
	}

	return notebookID, nil
}

func getUserId(c *gin.Context) (int, error) {
	id := c.GetString(userCtx)
	userId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
