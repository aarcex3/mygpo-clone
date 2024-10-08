package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService interface {
	Register(ctx *gin.Context, form *RegistrationForm) error
	Authenticate(cxt *gin.Context, form *LoginForm) (string, error)
}

type service struct {
	userRepository users.UserRepository
}

func Service(repo users.UserRepository) *service {
	return &service{userRepository: repo}
}

func (as *service) Register(ctx *gin.Context, form *RegistrationForm) error {
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
func (as *service) Authenticate(cxt *gin.Context, form *LoginForm) (string, error) {
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
	secretKey := config.LoadConfig().SecretKey
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func (as *service) HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (as *service) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
