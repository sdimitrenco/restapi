package service

import (
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
)

type Authorisation interface {
	CreateUser(user todo.User) (id int, err error)
	GenerateToken(username, password string) (token string, err error)
	ParseToken(token string) (id int, err error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (idList int, err error)
	GetAll(userId int) (lists []todo.TodoList, err error)
	GetById(userId, listId int) (list todo.TodoList, err error)
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
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
