package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/aarcex3/mygpo-clone/internals"
	"github.com/aarcex3/mygpo-clone/test/testconfig"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper to perform a request
func performRequest(app *gin.Engine, method, path string, form url.Values) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(res, req)
	return res
}

// Common setup for the tests
func setupAppWithDB() (*gin.Engine, *sql.DB, func()) {
	app := gin.Default()
	db, cleanup := testconfig.SetupDatabase()
	internals.SetUpApp(app, db)
	return app, db, cleanup
}

func TestRegistration(t *testing.T) {
	t.Run("Normal registration", func(t *testing.T) {
		app, db, cleanup := setupAppWithDB()
		defer cleanup()

		form := url.Values{
			"username": {"johndoe"},
			"email":    {"john@example.com"},
			"password": {"supersecretpassword"},
		}

		res := performRequest(app, "POST", "/v1/auth/registration", form)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.JSONEq(t, `{"message":"Registration successful"}`, res.Body.String())

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "johndoe").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("Try to register an existing user", func(t *testing.T) {
		app, _, cleanup := setupAppWithDB()
		defer cleanup()

		form := url.Values{
			"username": {"johndoe"},
			"email":    {"john@example.com"},
			"password": {"supersecretpassword"},
		}

		_ = performRequest(app, "POST", "/v1/auth/registration", form)

		// Try registering the same user again
		res := performRequest(app, "POST", "/v1/auth/registration", form)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.JSONEq(t, `{"message":"User already exists"}`, res.Body.String())
	})
}

func TestLogin(t *testing.T) {
	app, _, cleanup := setupAppWithDB()
	defer cleanup()

	// Register the user first
	registrationForm := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}
	_ = performRequest(app, "POST", "/v1/auth/registration", registrationForm)

	t.Run("Successful login", func(t *testing.T) {
		form := url.Values{
			"username": {"johndoe"},
			"password": {"supersecretpassword"},
		}
		res := performRequest(app, "POST", "/v1/auth/login", form)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, `{"message": "Login successful"}`, res.Body.String())

		authHeader := res.Header().Get("Authorization")
		assert.NotEmpty(t, authHeader)
		assert.Contains(t, authHeader, "Bearer")
	})

	t.Run("Login with wrong username", func(t *testing.T) {
		form := url.Values{
			"username": {"johndo"},
			"password": {"supersecretpassword"},
		}
		res := performRequest(app, "POST", "/v1/auth/login", form)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

		authHeader := res.Header().Get("Authorization")
		assert.Empty(t, authHeader)
	})

	t.Run("Login with wrong password", func(t *testing.T) {
		form := url.Values{
			"username": {"johndoe"},
			"password": {"notasecretpassword"},
		}
		res := performRequest(app, "POST", "/v1/auth/login", form)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

		authHeader := res.Header().Get("Authorization")
		assert.Empty(t, authHeader)
	})
}
