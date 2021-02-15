package repository

import (
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (idList int, err error)
	GetAll(userId int) (lists []todo.TodoList, err error)
	GetById(userId, listId int) (list todo.TodoList, err error)
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
