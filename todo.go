package todo

//TodoList struct
type TodoList struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
