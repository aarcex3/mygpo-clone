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

type userRepository struct {
	queries database.Queries
}

func Repository(queries database.Queries) *userRepository {
	return &userRepository{queries: queries}
}

func (ur *userRepository) Add(ctx *gin.Context, username string, password string, email string) error {
	return ur.queries.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email})
}

func (ur *userRepository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := ur.queries.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (ur *userRepository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = ur.queries.GetUserByUsername(ctx, username)
	return user, err
}
