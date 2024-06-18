package server

import (
	"database/sql"
	"log"

	"github.com/abroudoux/tonotdoapp/internal/services"
	"github.com/gin-gonic/gin"
)

func InitSever(port string) {
	r := gin.Default()

	// r.POST("/auth/register/:usernmame/:password", auth.Register)
	// r.POST("/auth/login/:usernmame/:password", auth.Login)
	db, err := sql.Open("sqlite3", "../../database.db")

	if err != nil {
		log.Fatal(err)
	}

	services.CreateUsersTable(db)
	services.CreateTodosTable(db)

	r.GET("/todos", services.GetTodos(db.Db))
	r.GET("/users", services.GetUsers(db.DB))

	r.Run(":" + port)
}