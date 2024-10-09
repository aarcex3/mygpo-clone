package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/app"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Test Case structure
type testCase struct {
	name           string
	form           url.Values
	expectedStatus int
	expectedBody   string
}

func TestRegistration(t *testing.T) {

	cfg := config.LoadConfig("test")
	router := gin.Default()
	db, cleanup := SetupTestDatabase(cfg)
	defer cleanup()

	_ = app.New(router, db, cfg)

	testCases := []testCase{
		{
			name: "Successful Registration",
			form: url.Values{
				"username": {"testuser"},
				"password": {"testpassword"},
				"email":    {"testuser@example.com"},
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "Registration successful",
		},
		{
			name: "Failed Registration",
			form: url.Values{
				"username": {"testuser"},
				"password": {"testpassword"},
				"email":    {"testuser@example.com"},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "User already exists",
		},
		{
			name: "Invalid Request Data",
			form: url.Values{
				"username": {""},
				"password": {},
				"email":    {"@example.com"},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request data",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/auth/registration", bytes.NewBufferString(tc.form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			assert.NoError(t, err)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.Contains(t, res.Body.String(), tc.expectedBody)

		})
	}
}
