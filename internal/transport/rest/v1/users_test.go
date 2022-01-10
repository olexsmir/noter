package v1

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/Smirnov-O/noter/internal/domain"
	"github.com/Smirnov-O/noter/internal/service"
	mock_service "github.com/Smirnov-O/noter/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_userSignUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, inp domain.UserSignUp)

	tests := []struct {
		name         string
		requestBody  string
		input        domain.UserSignUp
		statusCode   int
		mockBehavior mockBehavior
		responseBody string
	}{
		{
			name:        "ok",
			requestBody: "{\"name\": \"Bob Test\", \"email\": \"test@test.com\", \"password\": \"testpassword\"}",
			input: domain.UserSignUp{
				Name:     "Bob Test",
				Email:    "test@test.com",
				Password: "testpassword",
			},
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {
				r.EXPECT().SignUp(inp).Return(nil)
			},
			statusCode: 201,
		},
		{
			name:        "missing name",
			requestBody: "{\"name\": \"\", \"email\": \"test@test.com\", \"password\": \"testpassword\"}",
			input: domain.UserSignUp{
				Name:     "",
				Email:    "test@test.com",
				Password: "testpassword",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
		{
			name:        "invalid name",
			requestBody: "{\"name\": \"a\", \"email\": \"test@test.com\", \"password\": \"testpassword\"}",
			input: domain.UserSignUp{
				Name:     "a",
				Email:    "test@test.com",
				Password: "testpassword",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
		{
			name:        "missing email",
			requestBody: "{\"name\": \"Bob Test\", \"email\": \"\", \"password\": \"testpassword\"}",
			input: domain.UserSignUp{
				Name:     "Bob Test",
				Email:    "",
				Password: "testpassword",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
		{
			name:        "invalid email",
			requestBody: "{\"name\": \"Bob Test\", \"email\": \"test\", \"password\": \"testpassword\"}",
			input: domain.UserSignUp{
				Name:     "Bob Test",
				Email:    "test",
				Password: "testpassword",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
		{
			name:        "missing password",
			requestBody: "{\"name\": \"Bob Test\", \"email\": \"test@test.com\", \"password\": \"\"}",
			input: domain.UserSignUp{
				Name:     "Bob Test",
				Email:    "test@test.com",
				Password: "",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
		{
			name:        "invalid password",
			requestBody: "{\"name\": \"Bob Test\", \"email\": \"test@test.com\", \"password\": \"qwerty\"}",
			input: domain.UserSignUp{
				Name:     "Bob Test",
				Email:    "test@test.com",
				Password: "qwerty",
			},
			statusCode:   400,
			responseBody: "{\"message\":\"invalid input body\"}",
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignUp) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockUsers(c)
			tt.mockBehavior(s, tt.input)

			services := &service.Services{User: s}
			handler := Handler{services: services}

			// Init endpoint
			r := gin.New()
			r.GET("/sign-up", func(c *gin.Context) {
			}, handler.userSignUp)

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sign-up", bytes.NewBufferString(tt.requestBody))

			// Make request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}

func TestHandler_userSignIn(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, inp domain.UserSignIn)

	tests := []struct {
		name         string
		requestBody  string
		input        domain.UserSignIn
		statusCode   int
		mockBehavior mockBehavior
		responseBody string
	}{
		{
			name:        "ok",
			requestBody: "{\"email\": \"test@test.com\", \"password\": \"testpassword\"}",
			input: domain.UserSignIn{
				Email:    "test@test.com",
				Password: "testpassword",
			},
			statusCode: 200,
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignIn) {
				r.EXPECT().SignIn(inp).Return(domain.Tokens{Access: "access", Refresh: "refresh"}, nil)
			},
			responseBody: "{\"access_token\":\"access\",\"refresh_token\":\"refresh\"}",
		},
		{
			name:        "missing email",
			requestBody: "{\"email\": \"\", \"password\": \"testpassword\"}",
			input: domain.UserSignIn{
				Email:    "",
				Password: "testpassword",
			},
			statusCode:   400,
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignIn) {},
			responseBody: "{\"message\":\"invalid input body\"}",
		},
		{
			name:        "invalid email",
			requestBody: "{\"email\": \"test\", \"password\": \"testpassword\"}",
			input: domain.UserSignIn{
				Email:    "test",
				Password: "testpassword",
			},
			statusCode:   400,
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignIn) {},
			responseBody: "{\"message\":\"invalid input body\"}",
		},
		{
			name:        "missing password",
			requestBody: "{\"email\": \"test@test.com\", \"password\": \"\"}",
			input: domain.UserSignIn{
				Email:    "test@test.com",
				Password: "",
			},
			statusCode:   400,
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignIn) {},
			responseBody: "{\"message\":\"invalid input body\"}",
		},
		{
			name:        "invalid password",
			requestBody: "{\"email\": \"test@test.com\", \"password\": \"a\"}",
			input: domain.UserSignIn{
				Email:    "test@test.com",
				Password: "a",
			},
			statusCode:   400,
			mockBehavior: func(r *mock_service.MockUsers, inp domain.UserSignIn) {},
			responseBody: "{\"message\":\"invalid input body\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockUsers(c)
			tt.mockBehavior(s, tt.input)

			services := &service.Services{User: s}
			handler := Handler{services: services}

			// Init endpoint
			r := gin.New()
			r.GET("/sign-in", func(c *gin.Context) {}, handler.userSignIn)

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sign-in", bytes.NewBufferString(tt.requestBody))

			// Make request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}
