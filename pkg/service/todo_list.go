package service

import (
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (idList int, err error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) (lists []todo.TodoList, err error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (list todo.TodoList, err error) {
	return s.repo.GetById(userId, listId)
}
