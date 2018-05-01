package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
	router := mux.NewRouter()
	router.HandleFunc("/Topics", GetTopics).Methods("GET")
	
}
