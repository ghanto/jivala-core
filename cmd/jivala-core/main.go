package main

import (
	"log"
	"net/http"

	"github.com/ghanto/jivala-core/pkg/user/interfaces/http/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", user.UsersHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
