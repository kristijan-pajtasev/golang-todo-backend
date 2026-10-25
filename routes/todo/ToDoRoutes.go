package todo

import (
	"golang-todo-backend-project/middleware"

	"github.com/gorilla/mux"
)

type ApiResponse struct {
	Status bool `json:"status"`
}

func ToDoRoutes(mainRouter *mux.Router) {
	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("", middleware.Chain(getAllToDoHandler, middleware.UserValidator())).Methods("GET")

	todoRouter.HandleFunc("", createToDoHandler).Methods("POST")

	todoRouter.HandleFunc("/{id}", getTodoByIdHandler).Methods("GET")

	todoRouter.HandleFunc("/{id}", deleteTodoByIdHandler).Methods("DELETE")
}
