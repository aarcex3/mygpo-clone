package directory

import (
	"net/http"
	"strconv"

	"github.com/aarcex3/mygpo-clone/internals/podcast"
	"github.com/aarcex3/mygpo-clone/internals/tags"
	"github.com/gin-gonic/gin"
)

type controller struct {
	tagService     tags.TagService
	podcastService podcast.PodcastService
}

func Controller(tagService tags.TagService, podcastService podcast.PodcastService) *controller {
	return &controller{tagService: tagService, podcastService: podcastService}
}

func (c *controller) GetTopTags(ctx *gin.Context) {
	limitStr := ctx.Param("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit intput"})
		return
	}
	tags, err := c.tagService.FindTopTags(ctx, int64(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something failed"})
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (c *controller) GetPodcastData(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Podcast url required"})
		return
	}
	podcast, err := c.podcastService.FindPodcastData(ctx, url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Podcast not found"})
		return
	}
	ctx.JSONP(http.StatusOK, podcast)
}
