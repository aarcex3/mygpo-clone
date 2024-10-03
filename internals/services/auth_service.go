package services

import (
	"errors"

	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/aarcex3/mygpo-clone/internals/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx *gin.Context, user *schemas.User) error
	Authenticate(username, password string) (bool, error)
}

type service struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *service {
	return &service{UserRepository: repo}
}

func (s *service) Register(ctx *gin.Context, user *schemas.User) error {
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return err
	}
	if s.UserRepository.UserExists(ctx, user.Username, user.Email) {
		return errors.New("user already exists")
	}

	if err := s.UserRepository.Add(ctx, user.Username, hashedPassword, user.Email); err != nil {
		return err

	}
	return nil

}
func (s *service) Authenticate(username, password string) (bool, error) {
	//var user database.User = s.UserRepository.FindUser(username)
	return true, nil
}

// HashPassword hashes the given password using bcrypt
func (s *service) HashPassword(password string) (string, error) {
	// Generate a hashed password with a cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
