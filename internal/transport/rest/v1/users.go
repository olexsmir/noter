package v1

import (
	"errors"
	"net/http"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	user := api.Group("/auth")
	{
		user.POST("/sign-up", h.userSignUp)
		user.POST("/sign-in", h.userSignIn)
		user.POST("/refresh-tokens", h.userRefreshTokens)

		authenticated := user.Group("/", h.userIdentity)
		{
			authenticated.POST("/logout", h.userLogout)
		}
	}
}

type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp userSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.User.SignUp(domain.UserSignUp{
		Name:     inp.Name,
		Email:    inp.Email,
		Password: inp.Password,
	}); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type userSignInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) userSignIn(c *gin.Context) {
	var inp userSignInInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.User.SignIn(domain.UserSignIn{
		Email:    inp.Email,
		Password: inp.Password,
	})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	})
}

type userRefreshTokensInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (h *Handler) userRefreshTokens(c *gin.Context) {
	var inp userRefreshTokensInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.User.RefreshTokens(inp.RefreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrSessionNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	})
}

func (h *Handler) userLogout(c *gin.Context) {
	userID := getUserId(c)

	if err := h.services.User.Logout(userID); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
