package users

import (
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Add(ctx *gin.Context, username, password, email string) error
	UserExists(ctx *gin.Context, username, email string) bool
	FindUserByUsername(ctx *gin.Context, username string) (database.User, error)
}

type repository struct {
	queries database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{queries: queries}
}

func (repo *repository) Add(ctx *gin.Context, username string, password string, email string) error {
	return repo.queries.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email})
}

func (repo *repository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := repo.queries.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (repo *repository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = repo.queries.GetUserByUsername(ctx, username)
	return user, err
}
