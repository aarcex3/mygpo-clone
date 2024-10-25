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
	db database.Queries
}

func Repository(queries database.Queries) *repository {
	return &repository{db: queries}
}

func (repo *repository) Add(ctx *gin.Context, username string, password string, email string) error {
	return repo.db.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email})
}

func (repo *repository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := repo.db.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (repo *repository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = repo.db.GetUserByUsername(ctx, username)
	return user, err
}
