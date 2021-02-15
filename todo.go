package todo

//TodoList struct
type TodoList struct {
	ID          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

//UserList struct
type UserList struct {
	ID     int
	UserID int
	ListID int
}

//TodoItem struct
type TodoItem struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

//ListItem struct
type ListItem struct {
	ID     int
	ListID int
	ItemID int
}
