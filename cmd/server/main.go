package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/update/{type}/{key}/{value}", handler).Methods("POST")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "%v %v %v\n", vars["type"], vars["key"], vars["value"])
}
