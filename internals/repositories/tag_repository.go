package repositories

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagRepository interface {
	GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error)
}

type tagRepository struct {
	queries database.Queries
}

func NewTagRepository(queries database.Queries) *tagRepository {
	return &tagRepository{queries: queries}
}

func (tr *tagRepository) GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error) {
	tags, err := tr.queries.TopTags(ctx)
	return tags, err
}
