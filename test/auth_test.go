package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/aarcex3/mygpo-clone/test/testconfig"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper to perform a request
func performFormRequest(app *gin.Engine, method, path string, form url.Values) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(res, req)
	return res
}

func TestRegistrationNormal(t *testing.T) {
	app, db, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()

	form := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}

	res := performFormRequest(app, "POST", "/v1/auth/registration", form)
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.JSONEq(t, `{"message":"Registration successful"}`, res.Body.String())

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "johndoe").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

}

func TestRegistrationUserExists(t *testing.T) {
	app, _, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()

	form := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}

	_ = performFormRequest(app, "POST", "/v1/auth/registration", form)

	// Try registering the same user again
	res := performFormRequest(app, "POST", "/v1/auth/registration", form)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.JSONEq(t, `{"message":"User already exists"}`, res.Body.String())
}

func TestLoginSuccess(t *testing.T) {
	app, _, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()

	registrationForm := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}
	_ = performFormRequest(app, "POST", "/v1/auth/registration", registrationForm)

	form := url.Values{
		"username": {"johndoe"},
		"password": {"supersecretpassword"},
	}
	res := performFormRequest(app, "POST", "/v1/auth/login", form)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, `{"message": "Login successful"}`, res.Body.String())

	authHeader := res.Header().Get("Authorization")
	assert.NotEmpty(t, authHeader)
	assert.Contains(t, authHeader, "Bearer")

}

func TestLoginWrongUsername(t *testing.T) {
	app, _, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()

	registrationForm := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}
	_ = performFormRequest(app, "POST", "/v1/auth/registration", registrationForm)

	form := url.Values{
		"username": {"johndo"},
		"password": {"supersecretpassword"},
	}
	res := performFormRequest(app, "POST", "/v1/auth/login", form)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

	authHeader := res.Header().Get("Authorization")
	assert.Empty(t, authHeader)

}

func TestLoginWrongPassword(t *testing.T) {
	app, _, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()

	registrationForm := url.Values{
		"username": {"johndoe"},
		"email":    {"john@example.com"},
		"password": {"supersecretpassword"},
	}
	_ = performFormRequest(app, "POST", "/v1/auth/registration", registrationForm)

	form := url.Values{
		"username": {"johndoe"},
		"password": {"notasecretpassword"},
	}
	res := performFormRequest(app, "POST", "/v1/auth/login", form)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

	authHeader := res.Header().Get("Authorization")
	assert.Empty(t, authHeader)
}
