package tags

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type TagService interface {
	FindTopTags(ctx *gin.Context, limit int64) ([]database.ListTopTagsRow, error)
}

type service struct {
	tagRepository TagRepository
}

func Service(repo TagRepository) *service {
	return &service{tagRepository: repo}
}

func (s *service) FindTopTags(ctx *gin.Context, limit int64) ([]database.ListTopTagsRow, error) {
	tags, err := s.tagRepository.ListTopTags(ctx, limit)
	return tags, err
}
