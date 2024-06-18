package main

import (
	"github.com/abroudoux/tonotdoapp/internal/db"
	"github.com/abroudoux/tonotdoapp/internal/server"
)

func main() {
	db.InitDatabase()
	server.InitSever("8080")
} 