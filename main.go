package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	Router(r)
	http.ListenAndServe(":8088", r)
}
