package tags

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagRepository interface {
	ListTopTags(ctx *gin.Context, limit int64) ([]database.ListTopTagsRow, error)
}

type repository struct {
	db database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{db: queries}
}

func (r *repository) ListTopTags(ctx *gin.Context, limit int64) ([]database.ListTopTagsRow, error) {

	tags, err := r.db.ListTopTags(ctx, int64(limit))

	return tags, err
}
