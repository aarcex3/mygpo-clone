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

type authservice struct {
	userRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *authservice {
	return &authservice{userRepository: repo}
}

func (as *authservice) Register(ctx *gin.Context, form *schemas.Registration) error {
	hashedPassword, err := as.HashPassword(form.Password)
	if err != nil {
		return err
	}
	if as.userRepository.UserExists(ctx, form.Username, form.Email) {
		return errors.New("user already exists")
	}

	if err := as.userRepository.Add(ctx, form.Username, hashedPassword, form.Email); err != nil {
		return err

	}
	return nil

}
func (as *authservice) Authenticate(cxt *gin.Context, form *schemas.Login) (string, error) {
	user, err := as.userRepository.FindUserByUsername(cxt, form.Username)
	if err != nil {
		return "", err
	}
	if err := as.VerifyPassword(user.Password, form.Password); err != nil {
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

func (as *authservice) HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (as *authservice) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
