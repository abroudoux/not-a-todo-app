package types

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin   bool   `json:"admin"`
}