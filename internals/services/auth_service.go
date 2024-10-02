package services

import "github.com/aarcex3/mygpo-clone/internals/schemas"

type IAuthService interface {
	Register(user *schemas.User) error
	Authenticate(username, password string) (bool, error)
}

type AuthService struct {
	// UserRepository repositores.UserRepository
}

// func NewAuthService(repo repositories.UserRepository) *AuthService {
//     return &AuthService{UserRepository: repo}
// }

func (s *AuthService) Register(user *schemas.User) error {
	// if err := s.UserRepository.CreateUser(user.Username, user.Password, user.Email); err != nill {
	//
	//}
	return nil

}
func (s *AuthService) Authenticate(username, password string) (bool, error) {
	//var user models.User = s.UserRepository.FindUser(username)
	return true, nil
}
