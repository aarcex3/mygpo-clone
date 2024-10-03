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
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/registration", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(w, req)
	return w
}

func TestRegistration(t *testing.T) {
	app := gin.Default()
	db, cleanup := testconfig.SetupDatabase()
	defer cleanup()

	internals.SetUpApp(app, db)

	// Prepare the registration form data
	form := url.Values{}
	form.Set("username", "johndoe")
	form.Set("email", "john@example.com")
	form.Set("password", "supersecretpassword")

	// Perform the registration request
	w := performRegistrationRequest(app, form)

	// Assert the response status and body
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"message\":\"Registration successful\"}", w.Body.String())

	// Verify the user is created in the database
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "johndoe").Scan(&count); err != nil {
		t.Fatalf("Failed to count users: %v", err)
	}
	assert.Equal(t, 1, count)
}
