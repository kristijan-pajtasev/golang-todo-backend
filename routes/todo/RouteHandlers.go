package todo

import (
	"encoding/json"
	"golang-todo-backend-project/models/todo"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllToDoHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todo.GetAllTodos())
}

func createToDoHandler(w http.ResponseWriter, r *http.Request) {
	var todoVal todo.ToDo
	err := json.NewDecoder(r.Body).Decode(&todoVal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid object"}`))
		return
	}

	newTodo := todo.AddTodo(todoVal.Title, todoVal.Content, todoVal.IsCompleted)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	id, _ := strconv.Atoi(todoId)
	json.NewEncoder(w).Encode(todo.GetTodoById(id))
}

func deleteTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	id, _ := strconv.Atoi(todoId)
	success := todo.DeleteTodoById(id)
	json.NewEncoder(w).Encode(ApiResponse{Status: success})
}
