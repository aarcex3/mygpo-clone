package controllers

import (
	"net/http"

	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
)

type DirectoryController struct {
	tagService services.TagService
}

func NewDirectoryController(service services.TagService) *DirectoryController {
	return &DirectoryController{tagService: service}
}

func (dc *DirectoryController) RetrieveTopTags(ctx *gin.Context) {
	tags, err := dc.tagService.GetTopTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something failed"})
	}

	ctx.JSON(http.StatusOK, gin.H{"top_tags": tags})
}
