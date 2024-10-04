package test

import (
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

func performRegistrationRequest(app *gin.Engine, form url.Values) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/registration", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(res, req)
	return res
}

func performLoginRequest(app *gin.Engine, form url.Values) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/login", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(res, req)
	return res
}

func TestRegistration(t *testing.T) {
	// Set up
	app := gin.Default()
	db, cleanup := testconfig.SetupDatabase()
	defer cleanup()

	internals.SetUpApp(app, db)

	form := url.Values{}
	form.Set("username", "johndoe")
	form.Set("email", "john@example.com")
	form.Set("password", "supersecretpassword")

	res := performRegistrationRequest(app, form)

	assert.Equal(t, http.StatusCreated, res.Code)
	assert.JSONEq(t, `{"message":"Registration successful"}`, res.Body.String())

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "johndoe").Scan(&count); err != nil {
		t.Fatalf("Failed to count users: %v", err)
	}
	assert.Equal(t, 1, count)

	// Try to register an existing user
	form = url.Values{}
	form.Set("username", "johndoe")
	form.Set("email", "john@example.com")
	form.Set("password", "supersecretpassword")

	res = performRegistrationRequest(app, form)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.JSONEq(t, `{"message":"User already exists"}`, res.Body.String())
}

func TestLogin(t *testing.T) {
	// Set up
	app := gin.Default()
	db, cleanup := testconfig.SetupDatabase()
	defer cleanup()

	internals.SetUpApp(app, db)

	form := url.Values{}
	form.Set("username", "johndoe")
	form.Set("email", "john@example.com")
	form.Set("password", "supersecretpassword")
	_ = performRegistrationRequest(app, form)

	// Normal login
	form = url.Values{}
	form.Set("username", "johndoe")
	form.Set("password", "supersecretpassword")

	res := performLoginRequest(app, form)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, `{"message": "Login successful"}`, res.Body.String())

	authHeader := res.Header().Get("Authorization")
	assert.NotEmpty(t, authHeader, "Authorization header should be set")
	assert.Contains(t, authHeader, "Bearer", "Authorization header should contain Bearer token")

	// Wrong username
	form = url.Values{}
	form.Set("username", "johndo")
	form.Set("password", "supersecretpassword")

	res = performLoginRequest(app, form)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

	authHeader = res.Header().Get("Authorization")
	assert.Empty(t, authHeader, "Authorization header should  not be set")
	assert.NotContains(t, authHeader, "Bearer", "Authorization header should  not contain Bearer token")

	// Wrong password
	form = url.Values{}
	form.Set("username", "johndoe")
	form.Set("password", "notasecretpassword")

	res = performLoginRequest(app, form)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.JSONEq(t, `{"message": "Login error"}`, res.Body.String())

	authHeader = res.Header().Get("Authorization")
	assert.Empty(t, authHeader, "Authorization header should  not be set")
	assert.NotContains(t, authHeader, "Bearer", "Authorization header should  not contain Bearer token")
}
