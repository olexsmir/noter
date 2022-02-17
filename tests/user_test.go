package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

func (s *APITestSuite) TestUserSignUP() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	name, email, password := "Test User", "test@test.com", "qwerty123"
	signUpData := fmt.Sprintf(`{"name":"%s","email":"%s","password":"%s"}`, name, email, password)

	req, _ := http.NewRequest("POST", "/api/v1/auth/sign-up", bytes.NewBuffer([]byte(signUpData)))
	req.Header.Set("Content-type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusCreated, resp.Result().StatusCode)

	var user domain.User
	err := s.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	r.NoError(err)

	passwordHash, err := s.hasher.Hash(password)
	s.NoError(err)

	r.Equal(name, user.Name)
	r.Equal(passwordHash, user.Password)
	r.Equal(email, user.Email)
}
