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

	todoRouter.HandleFunc("", middleware.Chain(createToDoHandler, middleware.UserValidator())).Methods("POST")

	todoRouter.HandleFunc("/{id}", middleware.Chain(getTodoByIdHandler, middleware.UserValidator())).Methods("GET")

	todoRouter.HandleFunc("/{id}", middleware.Chain(deleteTodoByIdHandler, middleware.UserValidator())).Methods("DELETE")
}
