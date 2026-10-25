package main

import "github.com/gorilla/mux"
import (
	"golang-todo-backend-project/routes/todo"
)

func Router(r *mux.Router) {
	todo.ToDoRoutes(r)
}
