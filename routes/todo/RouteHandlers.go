package todo

import (
	"encoding/json"
	"golang-todo-backend-project/helpers"
	"golang-todo-backend-project/models/todo"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllToDoHandler(w http.ResponseWriter, r *http.Request) {
	userId := helpers.GetCurrentUserKey(r)
	json.NewEncoder(w).Encode(todo.GetAllTodos(userId))
}

func createToDoHandler(w http.ResponseWriter, r *http.Request) {
	var todoVal todo.ToDo
	err := json.NewDecoder(r.Body).Decode(&todoVal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid object"}`))
		return
	}

	userId := helpers.GetCurrentUserKey(r)

	newTodo := todo.AddTodo(userId, todoVal.Title, todoVal.Content, todoVal.IsCompleted)
	json.NewEncoder(w).Encode(newTodo)
}

func updateToDoHandler(w http.ResponseWriter, r *http.Request) {
	var todoVal todo.ToDo
	vars := mux.Vars(r)
	todoId := vars["id"]
	id, _ := strconv.Atoi(todoId)
	err := json.NewDecoder(r.Body).Decode(&todoVal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid object"}`))
	}

	userId := helpers.GetCurrentUserKey(r)
	updatedTodo := todo.UpdateTodoById(userId, id, todoVal.Title, todoVal.Content, todoVal.IsCompleted)
	json.NewEncoder(w).Encode(updatedTodo)
}

func getTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	id, _ := strconv.Atoi(todoId)

	userId := helpers.GetCurrentUserKey(r)
	json.NewEncoder(w).Encode(todo.GetTodoById(userId, id))
}

func deleteTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	id, _ := strconv.Atoi(todoId)

	userId := helpers.GetCurrentUserKey(r)
	success := todo.DeleteTodoById(userId, id)
	json.NewEncoder(w).Encode(ApiResponse{Status: success})
}
