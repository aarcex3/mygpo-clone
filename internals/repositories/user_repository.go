package repositories

import (
	"net/http"

	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Add(ctx *gin.Context, username, password, email string) error
	FindUserByUserName(username string)
	FindUserById(ctx *gin.Context, id int64) database.User
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

func (r *repository) FindUserByUserName(username string) {}

func (r *repository) FindUserById(ctx *gin.Context, id int64) database.User {
	var user, err = r.Queries.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
	return user
}
