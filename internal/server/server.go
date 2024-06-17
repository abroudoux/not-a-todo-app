package server

import (
	"github.com/abroudoux/tonotdoapp/internal/services"
	"github.com/gin-gonic/gin"
)

func NewServer(port string) {
	r := gin.Default()

	r.GET("/todos", services.GetTodos)

	r.Run(":" + port)
}