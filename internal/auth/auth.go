package auth

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (token string, err error) {
	username := c.Param("username")
	password := c.Param("password")

	token = "token" + username + password

	return "token", nil
}

func Login(c *gin.Context) (authorized bool, err error) {
	username := c.Param("username")
	password := c.Param("password")

	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, nil
}