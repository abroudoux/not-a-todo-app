package services

import (
	"net/http"

	"github.com/abroudoux/tonotdoapp/internal/db"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.Todos)
}