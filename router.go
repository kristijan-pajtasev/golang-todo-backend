package main

import (
	"encoding/json"
	"golang-todo-backend-project/config"
	"golang-todo-backend-project/helpers"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
)
import (
	"golang-todo-backend-project/routes/todo"
)

type ApiResponse struct {
	Status bool `json:"status"`
}

func Router(r *mux.Router) {
	r.HandleFunc("/validate-user", func(w http.ResponseWriter, r *http.Request) {
		userId := helpers.GetCurrentUserKey(r)
		if slices.Contains(config.Config.AllowedUsersList, userId) {
			json.NewEncoder(w).Encode(ApiResponse{Status: true})
		} else {
			json.NewEncoder(w).Encode(ApiResponse{Status: false})
		}
	}).Methods("GET")

	todo.ToDoRoutes(r)
}
