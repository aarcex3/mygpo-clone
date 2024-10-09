package tags

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagService interface {
	GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error)
}

type service struct {
	tagRepository TagRepository
}

func Service(repo TagRepository) *service {
	return &service{tagRepository: repo}
}

func (s *service) GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error) {
	tags, err := s.tagRepository.GetTopTags(ctx)
	return tags, err
}
