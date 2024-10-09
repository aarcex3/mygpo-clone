package test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/app"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var db *sql.DB
var cleanup func()

func TestMain(m *testing.M) {

	cfg := config.LoadConfig("test")

	db, cleanup = SetupTestDatabase(cfg)

	router = gin.Default()
	_ = app.New(router, db, cfg)

	code := m.Run()

	cleanup()

	os.Exit(code)
}

type testCase struct {
	name           string
	form           url.Values
	expectedStatus int
	expectedBody   string
}

func TestRegistration(t *testing.T) {
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

func TestLogin(t *testing.T) {
	form := url.Values{
		"username": {"testuser2"},
		"password": {"testpassword2"},
		"email":    {"testuser2@example.com"},
	}

	req, err := http.NewRequest("POST", "/v1/auth/registration", bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.NoError(t, err)

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	testCases := []testCase{
		{
			name: "Failed Login",
			form: url.Values{
				"username": {"testuser"},
				"password": {"wrongpassword"},
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Login error",
		},
		{
			name: "Invalid Request Data",
			form: url.Values{
				"username": {},
				"password": {},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request data",
		},
		{
			name: "Successful Login",
			form: url.Values{
				"username": {"testuser2"},
				"password": {"testpassword2"},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Login successful",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/auth/login", bytes.NewBufferString(tc.form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			assert.NoError(t, err)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.Contains(t, res.Body.String(), tc.expectedBody)
			if tc.expectedBody == "Successful Login" {
				assert.Contains(t, res.Header(), "Bearer")
			}
		})
	}
}
