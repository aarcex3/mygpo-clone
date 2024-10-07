package test

import (
	"net/http"
	"testing"

	"github.com/aarcex3/mygpo-clone/test/testconfig"
	"github.com/stretchr/testify/assert"
)

func TestGetTopTags(t *testing.T) {
	app, _, cleanup := testconfig.SetupAppWithDB()
	defer cleanup()
	res := testconfig.PerformRequest(app, "GET", "/v1/tags")
	assert.Equal(t, http.StatusOK, res.Code)
}
