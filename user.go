package todo

//User type
type User struct {
	ID       int    `json: "-"`
	Name     string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
}