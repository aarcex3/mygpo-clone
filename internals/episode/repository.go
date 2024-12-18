package episode

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type episodeRepository interface {
	GetEpisodeByUrl(ctx *gin.Context, url string) (database.GetEpisodeByUrlRow, error)
}

type repository struct {
	db database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{db: queries}
}

func (r *repository) GetEpisodeByUrl(ctx *gin.Context, url string) (database.GetEpisodeByUrlRow, error) {
	return r.db.GetEpisodeByUrl(ctx, url)
}
