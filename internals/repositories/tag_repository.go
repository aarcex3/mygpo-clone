package repositories

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagRepository interface {
	GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error)
}

type tagRepository struct {
	Queries database.Queries
}

func NewTagRepository(queries database.Queries) *tagRepository {
	return &tagRepository{Queries: queries}
}

func (tr *tagRepository) GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error) {
	tags, err := tr.Queries.TopTags(ctx)
	return tags, err
}
