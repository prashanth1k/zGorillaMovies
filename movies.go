package main

import (
	"encoding/json"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var movies []Movie

func getMovies(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
