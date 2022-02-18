package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/flof-ik/noter/internal/domain"
	"github.com/gin-gonic/gin"
)

type notebookResponse struct {
	ID          int       `json:"id"`
	AuthorID    int       `json:"author_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *APITestSuite) TestNotebookGetByID() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	token, err := s.getJWT(1)
	s.NoError(err)

	req, _ := http.NewRequest("GET", "/api/v1/notebook/1", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusOK, resp.Result().StatusCode)

	respData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var notebookData notebookResponse
	err = json.Unmarshal(respData, &notebookData)
	s.NoError(err)

	r.Equal(notebook.ID, notebookData.ID)
	r.Equal(notebook.AuthorID, notebookData.AuthorID)
	r.Equal(notebook.Name, notebookData.Name)
	r.Equal(notebook.Description, notebookData.Description)
}

func (s *APITestSuite) TestNotebookGetAll() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	token, err := s.getJWT(1)
	s.NoError(err)

	req, _ := http.NewRequest("GET", "/api/v1/notebook/", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusOK, resp.Result().StatusCode)

	respData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var notebookData []notebookResponse
	err = json.Unmarshal(respData, &notebookData)
	s.NoError(err)

	r.NotNil(notebookData)
}

func (s *APITestSuite) TestNotebookCreate() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	userID := 1
	token, err := s.getJWT(userID)
	s.NoError(err)

	name, description := "testing notebook", "the testing notebook for test"
	createData := fmt.Sprintf(`{"name":"%s", "description":"%s"}`, name, description)

	req, _ := http.NewRequest("POST", "/api/v1/notebook/", bytes.NewBuffer([]byte(createData)))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusCreated, resp.Result().StatusCode)

	var dbNotebook domain.Notebook
	err = s.db.Get(&dbNotebook, "SELECT * FROM notebooks WHERE name = $1", name)
	r.NoError(err)

	r.Equal(userID, dbNotebook.AuthorID)
	r.Equal(name, dbNotebook.Name)
	r.Equal(description, dbNotebook.Description)
}

func (s *APITestSuite) TestNotebookUpdate() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	token, err := s.getJWT(1)
	s.NoError(err)

	name, description := "the new name", "the new testing notebook description"
	updateData := fmt.Sprintf(`{"name":"%s", "description":"%s"}`, name, description)

	req, _ := http.NewRequest("PUT", "/api/v1/notebook/1", bytes.NewBuffer([]byte(updateData)))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusOK, resp.Result().StatusCode)

	var dbNotebook domain.Notebook
	err = s.db.Get(&dbNotebook, "SELECT * FROM notebooks WHERE name = $1", name)
	r.NoError(err)

	r.Equal(name, dbNotebook.Name)
	r.Equal(description, dbNotebook.Description)
}

func (s *APITestSuite) TestNotebookDelete() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	token, err := s.getJWT(1)
	s.NoError(err)

	var notebookID int
	err = s.db.QueryRow("INSERT INTO notebooks (author_id, name, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		1, "notebook for tests", "the testing notebook for test", time.Now(), time.Now()).Scan(&notebookID)
	r.NoError(err)

	idStr := strconv.Itoa(notebookID)

	req, _ := http.NewRequest("DELETE", "/api/v1/notebook/"+idStr, nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusOK, resp.Result().StatusCode)

	var isExists bool
	err = s.db.QueryRow("SELECT exists(SELECT 1 FROM notebooks WHERE id=$1)", notebookID).Scan(&isExists)
	r.NoError(err)

	r.Equal(false, isExists)
}
