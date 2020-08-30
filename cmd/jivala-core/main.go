package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/users", user.)

	log.Fatal(http.ListenAndServe(":8000", r))
}
