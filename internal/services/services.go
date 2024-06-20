package services

import (
	"database/sql"
	"log"

	"github.com/abroudoux/tonotdoapp/internal/types"
)

func CreateUsersTable(db *sql.DB) error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, is_admin BOOLEAN)")

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()

	return nil
}

func CreateTodosTable(db *sql.DB) error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, id_user INTEGER, title TEXT, is_completed BOOLEAN)")

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()

	return nil
}

func CreateAdminUser(db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO users (id, username, password, is_admin) VALUES (1, user, password, true)")

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec("admin", "admin", true)

	return nil
}

func GetUsers(db *sql.DB) ([]types.User, error) {
	rows, err := db.Query("SELECT id, username, password, is_admin FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	users := []types.User{}

	for rows.Next() {
		var user types.User

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.IsAdmin)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func GetTodos(db *sql.DB) ([]types.Todo, error) {
	rows, err := db.Query("SELECT id, user_id title, is-completed FROM todos")


	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	todos := []types.Todo{}

	for rows.Next() {
		var todo types.Todo

		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Completed)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}