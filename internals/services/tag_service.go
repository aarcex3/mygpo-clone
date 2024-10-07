package services

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/gin-gonic/gin"
)

type TagService interface {
	GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error)
}

type tagService struct {
	tagRepository repositories.TagRepository
}

func NewTagService(repo repositories.TagRepository) *tagService {
	return &tagService{tagRepository: repo}
}

func (ts *tagService) GetTopTags(ctx *gin.Context) ([]database.TopTagsRow, error) {
	tags, err := ts.tagRepository.GetTopTags(ctx)
	return tags, err
}
