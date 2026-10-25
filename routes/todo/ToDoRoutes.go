package todo

import (
	"encoding/json"
	"golang-todo-backend-project/models/todo"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ToDoRoutes(mainRouter *mux.Router) {
	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(todo.GetAllTodos())
	}).Methods("GET")

	todoRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		newTodo := todo.AddTodo("New todo title", "New todo content", false)
		json.NewEncoder(w).Encode(newTodo)
	}).Methods("POST")

	todoRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoId := vars["id"]
		id, _ := strconv.Atoi(todoId)
		json.NewEncoder(w).Encode(todo.GetTodoById(id))
	}).Methods("GET")
}
