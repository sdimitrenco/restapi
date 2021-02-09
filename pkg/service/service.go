package service

import (
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
)

type Authorisation interface {
	CreateUser(user todo.User) (id int, err error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
	}
}
