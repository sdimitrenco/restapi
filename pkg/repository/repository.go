package repository

import (
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user todo.User) (id int, err error)
}

type TodoList interface {
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
	}
}
