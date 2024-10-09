package directory

import (
	"net/http"
	"strconv"

	"github.com/aarcex3/mygpo-clone/internals/tags"
	"github.com/gin-gonic/gin"
)

type controller struct {
	tagService tags.TagService
}

func Controller(service tags.TagService) *controller {
	return &controller{tagService: service}
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
