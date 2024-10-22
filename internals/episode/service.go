package episode

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type EpisodeService interface {
	FindEpisodeData(ctx *gin.Context, url string) (database.GetEpisodeByUrlRow, error)
}

type service struct {
	repo episodeRepository
}

func Service(repo episodeRepository) *service {
	return &service{repo: repo}
}

func (s *service) FindEpisodeData(ctx *gin.Context, url string) (database.GetEpisodeByUrlRow, error) {
	return s.repo.GetEpisodeByUrl(ctx, url)
}
