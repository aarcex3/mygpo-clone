package podcast

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type PodcastService interface {
	FindPodcastData(ctx *gin.Context, url string) (database.GetPodcastByUrlRow, error)
}

type service struct {
	repo podcastRepository
}

func Service(repo podcastRepository) *service {
	return &service{repo: repo}
}

func (s *service) FindPodcastData(ctx *gin.Context, url string) (database.GetPodcastByUrlRow, error) {
	return s.repo.GetPodcastByUrl(ctx, url)
}
