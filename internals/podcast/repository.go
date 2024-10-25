package podcast

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type podcastRepository interface {
	GetPodcastByUrl(ctx *gin.Context, url string) (database.GetPodcastByUrlRow, error)
}

type repository struct {
	db database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{db: queries}
}

func (r *repository) GetPodcastByUrl(ctx *gin.Context, url string) (database.GetPodcastByUrlRow, error) {
	return r.db.GetPodcastByUrl(ctx, url)
}
