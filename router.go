package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	http.ListenAndServe(":8088", r)
}
