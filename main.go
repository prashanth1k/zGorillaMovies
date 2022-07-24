package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies/", getMovies).Methods("GET")
	router.HandleFunc("/movies/id", getMovie).Methods("GET")
	router.HandleFunc("/movies/", postMovie).Methods("POST")
	router.HandleFunc("/movies/", putMovie).Methods("PUT")
	router.HandleFunc("/movies/", delMovie).Methods("DELETE")

	fmt.Printf("String server at Port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
