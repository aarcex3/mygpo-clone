package directory

import (
	"net/http"

	"github.com/aarcex3/mygpo-clone/internals/tags"
	"github.com/gin-gonic/gin"
)

type controller struct {
	tagService tags.TagService
}

func Controller(service tags.TagService) *controller {
	return &controller{tagService: service}
}

func (c *controller) RetrieveTopTags(ctx *gin.Context) {
	tags, err := c.tagService.GetTopTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something failed"})
	}

	ctx.JSON(http.StatusOK, tags)
}
