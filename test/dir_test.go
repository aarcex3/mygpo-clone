package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/app"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type Tag struct {
	Title string `json:"Title"`
	Code  string `json:"Code"`
	Usage int    `json:"Usage"`
}

func TestDirectory(t *testing.T) {
	cfg := config.LoadConfig("test")

	db, cleanup := SetupTestDatabase(cfg)
	defer cleanup()
	router := gin.Default()
	_ = app.New(router, db, cfg)

	t.Run("Get Top Tags", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/v1/tags/%d", 4), nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		var responseBody []Tag
		err = json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		expectedLength := 4
		assert.Equal(t, expectedLength, len(responseBody))

	})

	t.Run("Get Top Tags Invalid Limit", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/tags/f", nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Contains(t, res.Body.String(), "Invalid limit intput")

	})
}
