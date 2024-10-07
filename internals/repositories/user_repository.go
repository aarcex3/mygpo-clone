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
	Queries database.Queries
}

func NewUserRepository(queries database.Queries) *userRepository {
	return &userRepository{Queries: queries}
}

func (ur *userRepository) Add(ctx *gin.Context, username string, password string, email string) error {
	if err := ur.Queries.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email}); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := ur.Queries.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (ur *userRepository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = ur.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}
