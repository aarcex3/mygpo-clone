package repositories

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

func NewUserRepository(queries database.Queries) *userRepository {
	return &userRepository{queries: queries}
}

func (ur *userRepository) Add(ctx *gin.Context, username string, password string, email string) error {
	if err := ur.queries.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email}); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := ur.queries.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (ur *userRepository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = ur.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}
