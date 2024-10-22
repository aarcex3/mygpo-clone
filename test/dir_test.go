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

	t.Run("Get Podcast Data", func(t *testing.T) {
		url := "/v1/data/podcast?url=" + "http://feeds.feedburner.com/coverville"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		expectedContent := `{
				"Title": "Coverville",
				"Website": "http://coverville.com",
				"MygpoLink": "http://www.gpodder.net/podcast/16124",
				"Description": "The best cover songs, delivered to your ears two to three times a week!",
				"Subscribers": 19,
				"Author": "Brian Ibbott",
				"Url": "http://feeds.feedburner.com/coverville",
				"LogoUrl": "http://www.coverville.com/art/coverville_iTunes300.jpg"
			}`

		actualContent := res.Body.String()

		assert.JSONEq(t, expectedContent, actualContent, "Response body should match expected JSON content")

	})
	t.Run("Get Podcast Data Not Found", func(t *testing.T) {
		url := "/v1/data/podcast?url=" + "http://no.url.com"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		assert.Contains(t, res.Body.String(), "Podcast not found")

	})
	t.Run("Get Podcast Data No URL", func(t *testing.T) {
		url := "/v1/data/podcast"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)

		assert.Contains(t, res.Body.String(), "Podcast url required")

	})

	t.Run("Get Episode Data", func(t *testing.T) {
		url := "/v1/data/episode?url=" + "http://www.podtrac.com/pts/redirect.mp3/aolradio.podcast.aol.com/twit/twit0245.mp3"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		expectedContent := `{
					"Title": "TWiT 245: No Hitler For You",
					"Url": "http://www.podtrac.com/pts/redirect.mp3/aolradio.podcast.aol.com/twit/twit0245.mp3",
					"PodcastTitle": "this WEEK in TECH - MP3 Edition",
					"PodcastUrl": "http://leo.am/podcasts/twit",
					"Description": "A roundtable discussion about the latest trends in technology.",
					"Website": "http://www.podtrac.com/pts/redirect.mp3/aolradio.podcast.aol.com/twit/twit0245.mp3",
					"Released": "2010-12-25T00:30:00Z",
					"MygpoLink": "http://gpodder.net/episode/1046492"
				}`

		actualContent := res.Body.String()

		assert.JSONEq(t, expectedContent, actualContent, "Response body should match expected JSON content")

	})

	t.Run("Get Episode Data Not Found", func(t *testing.T) {
		url := "/v1/data/episode?url=" + "http://no.url.com"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		assert.Contains(t, res.Body.String(), "Episode not found")

	})
	t.Run("Get Episode Data No URL", func(t *testing.T) {
		url := "/v1/data/episode"

		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)

		assert.Contains(t, res.Body.String(), "Episode url required")

	})
}
