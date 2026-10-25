package main

import (
	"flag"
	"golang-todo-backend-project/config"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	allowedUsers := flag.String("allowedUsers", "user", "Comma separated allowed user list")
	port := flag.String("port", "8088", "Listening port")
	flag.Parse()
	config.SetAllowedUsers(strings.Split(*allowedUsers, ","))
	r := mux.NewRouter()
	Router(r)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":"+*port, r)
}
