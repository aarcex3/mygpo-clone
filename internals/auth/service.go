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

type AuthService interface {
	Register(ctx *gin.Context, form *RegistrationForm) error
	Authenticate(cxt *gin.Context, form *LoginForm) (string, error)
}

type service struct {
	userRepository users.UserRepository
	config         *config.Config
}

func Service(repo users.UserRepository, config *config.Config) *service {
	return &service{
		userRepository: repo,
		config:         config,
	}
}

func (s *service) Register(ctx *gin.Context, form *RegistrationForm) error {
	hashedPassword, err := s.HashPassword(form.Password)
	if err != nil {
		return err
	}
	if s.userRepository.UserExists(ctx, form.Username, form.Email) {
		return errors.New("user already exists")
	}

	if err := s.userRepository.Add(ctx, form.Username, hashedPassword, form.Email); err != nil {
		return err

	}
	return nil

}
func (s *service) Authenticate(cxt *gin.Context, form *LoginForm) (string, error) {
	user, err := s.userRepository.FindUserByUsername(cxt, form.Username)
	if err != nil {
		return "", err
	}
	if err := s.VerifyPassword(user.Password, form.Password); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.config.SecretKey)
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

func (s *service) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
