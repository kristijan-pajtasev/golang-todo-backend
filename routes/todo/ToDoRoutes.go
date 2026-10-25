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
		var todoVal todo.ToDo
		err := json.NewDecoder(r.Body).Decode(&todoVal)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid object"}`))
			return
		}

		newTodo := todo.AddTodo(todoVal.Title, todoVal.Content, todoVal.IsCompleted)
		json.NewEncoder(w).Encode(newTodo)
	}).Methods("POST")

	todoRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoId := vars["id"]
		id, _ := strconv.Atoi(todoId)
		json.NewEncoder(w).Encode(todo.GetTodoById(id))
	}).Methods("GET")
}
