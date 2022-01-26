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

// @Summary User SignUp
// @Tags user-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept json
// @Produce json
// @Param input body userSignUpInput true "sign up info"
// @Success 201
// @Failure 400,500 {object} response
// @Failure default {object} response
// @Router /auth/sign-up [post]
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

// @Summary User SignIn
// @Tags user-auth
// @Description user sign in
// @ModuleID authSignIn
// @Accept json
// @Produce json
// @Param input body userSignInInput true "sign up info"
// @Success 200 {object} domain.Tokens
// @Failure 400,500 {object} response
// @Failure default {object} response
// @Router /auth/sign-in [post]
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

// @Summary User Refresh Tokens
// @Tags user-auth
// @Description user refresh tokens
// @Accept json
// @Produce json
// @Param input body userRefreshTokensInput true "sign up info"
// @Success 200 {object} domain.Tokens
// @Failure 400,500 {object} response
// @Failure default {object} response
// @Router /auth/refresh-tokens [post]
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

// @Summary User logout
// @Security user-auth
// @Tags user-auth
// @Description user logout
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,500 {object} response
// @Failure default {object} response
// @Router /auth/logout [post]
func (h *Handler) userLogout(c *gin.Context) {
	userID := getUserId(c)

	if err := h.services.User.Logout(userID); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
