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

type repository struct {
	Queries database.Queries
}

func NewUserRepository(queries database.Queries) *repository {
	return &repository{Queries: queries}
}

func (r *repository) Add(ctx *gin.Context, username string, password string, email string) error {
	if err := r.Queries.CreateUser(ctx, database.CreateUserParams{Username: username, Password: password, Email: email}); err != nil {
		return err
	}
	return nil
}

func (r *repository) UserExists(ctx *gin.Context, username, email string) bool {
	count, _ := r.Queries.UserExists(ctx, database.UserExistsParams{Username: username, Email: email})
	return count > 0
}

func (r *repository) FindUserByUsername(ctx *gin.Context, username string) (database.User, error) {
	var user, err = r.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}
