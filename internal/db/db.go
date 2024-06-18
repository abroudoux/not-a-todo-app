package db

import (
	"database/sql"
	"log"

	types "github.com/abroudoux/tonotdoapp/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

var Todos []types.Todo = []types.Todo{
	{ ID: 1, Title: "Go to the bank", Completed: false },
	{ ID: 2, Title: "Buy some bread", Completed: true },
	{ ID: 3, Title: "Learn Go", Completed: false },
}

const DATABASE_PATH = "../../database.db"

func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DATABASE_PATH)

	if err != nil {
		log.Fatal(err)
	}

	CreateUsersTable(db)
	CreateTodosTable(db)

	return db, nil
}

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