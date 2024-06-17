package server

import (
	"github.com/abroudoux/tonotdoapp/internal/auth"
	"github.com/abroudoux/tonotdoapp/internal/services"
	"github.com/gin-gonic/gin"
)

func NewServer(port string) {
	r := gin.Default()

	r.POST("/auth/register/:usernmame/:password", auth.Register)
	r.POST("/auth/login/:usernmame/:password", auth.Login)


	r.GET("/todos", services.GetTodos)

	r.Run(":" + port)
}