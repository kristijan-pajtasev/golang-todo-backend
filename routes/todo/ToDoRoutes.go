package todo

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ToDoRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
}
