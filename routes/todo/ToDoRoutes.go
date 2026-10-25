package todo

import (
	"encoding/json"
	"golang-todo-backend-project/models/todo"
	"net/http"

	"github.com/gorilla/mux"
)

func ToDoRoutes(mainRouter *mux.Router) {
	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(todo.GetAllTodos())
	}).Methods("GET")

	todoRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
}
