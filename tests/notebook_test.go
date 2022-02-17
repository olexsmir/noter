package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
)

type notebookResponce struct {
	ID          int       `json:"id"`
	AuthorID    int       `json:"author_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *APITestSuite) TestNotebookGetAll() {
	router := gin.New()
	s.handler.Init(router.Group("/api"))
	r := s.Require()

	req, _ := http.NewRequest("GET", "/api/v1/notebook", nil)
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(http.StatusOK, resp.Result().StatusCode)

	var respNotebooks struct {
		Data []notebookResponce `json:"[0]"`
	}

	respData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	err = json.Unmarshal(respData, &respNotebooks)
	s.NoError(err)

	r.Equal(1, len(respNotebooks.Data))
}
