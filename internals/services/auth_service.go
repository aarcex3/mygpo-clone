package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/aarcex3/mygpo-clone/internals/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

type AuthService interface {
	Register(ctx *gin.Context, user *schemas.Registration) error
	Authenticate(cxt *gin.Context, form *schemas.Login) (string, error)
}

type service struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *service {
	return &service{UserRepository: repo}
}

func (s *service) Register(ctx *gin.Context, form *schemas.Registration) error {
	hashedPassword, err := s.HashPassword(form.Password)
	if err != nil {
		return err
	}
	if s.UserRepository.UserExists(ctx, form.Username, form.Email) {
		return errors.New("user already exists")
	}

	if err := s.UserRepository.Add(ctx, form.Username, hashedPassword, form.Email); err != nil {
		return err

	}
	return nil

}
func (s *service) Authenticate(cxt *gin.Context, form *schemas.Login) (string, error) {
	user, err := s.UserRepository.FindUserByUsername(cxt, form.Username)
	if err != nil {
		return "", err
	}
	if !s.VerifyPassword(user.Password, form.Password) {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,                               
		"username": user.Username,                        
		"exp":      time.Now().Add(time.Hour * 72).Unix(), 
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	
	return tokenString, nil
}

func (s *service) HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *service) VerifyPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
