package db

import types "github.com/abroudoux/tonotdoapp/internal/types"

var Todos []types.Todo = []types.Todo{
	{ ID: 1, Title: "Go to the bank", Completed: false },
	{ ID: 2, Title: "Buy some bread", Completed: true },
	{ ID: 3, Title: "Learn Go", Completed: false },
}

func CreateDatabase() {
	// This function is not implemented yet
}