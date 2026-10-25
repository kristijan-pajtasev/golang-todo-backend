package todo

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ToDoRoutes(mainRouter *mux.Router) {
	todoRouter := mainRouter.PathPrefix("/todos").Subrouter()
	todoRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	todoRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
}
