package repository

import (
	"fmt"
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (idList int, err error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createListQuery = fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)

	_, err = tx.Exec(createListQuery, userId, id)
	if err != nil {
		_ = tx.Rollback()
	}

	return id, tx.Commit()

}

func (s *TodoListPostgres) GetAll(userId int) (lists []todo.TodoList, err error) {
	var l []todo.TodoList

	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 ", todoListTable, usersListTable)
	err = s.db.Select(&l, query, userId)

	return l, err
}

func (s *TodoListPostgres) GetById(userId, listId int) (list todo.TodoList, err error) {
	var l todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListTable, usersListTable)
	err = s.db.Get(&l, query, userId, listId)

	return l, err
}
