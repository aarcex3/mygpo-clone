package services

import (
	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/aarcex3/mygpo-clone/internals/schemas"
)

type AuthService interface {
	Register(user *schemas.User) error
	Authenticate(username, password string) (bool, error)
}

type service struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *service {
	return &service{UserRepository: repo}
}

func (s *service) Register(user *schemas.User) error {
	// if err := s.UserRepository.Add(user.Username, user.Password, user.Email); err != nill {
	//
	//}
	return nil

}
func (s *service) Authenticate(username, password string) (bool, error) {
	//var user database.User = s.UserRepository.FindUser(username)
	return true, nil
}
