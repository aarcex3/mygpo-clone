package tags

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagRepository interface {
	GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error)
}

type repository struct {
	queries database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{queries: queries}
}

func (tr *repository) GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error) {
	tags, err := tr.queries.TopTags(ctx)
	return tags, err
}
