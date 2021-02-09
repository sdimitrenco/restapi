package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
)

const salt = "jhgghendgfhfp[267akjhfnv"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (id int, err error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
