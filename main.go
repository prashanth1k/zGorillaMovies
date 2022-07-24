package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", postMovie).Methods("POST")
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(customErrorHandler).GetHandler()

	fmt.Println("Starting server at Port 5000..\n")
	log.Fatal(http.ListenAndServe(":5000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetMovie..")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func postMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// generate random, hopefully unique string
	randId := make([]byte, 4)
	rand.Read(randId)

	movie.ID = hex.EncodeToString(randId)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie := movies[0]
	for index, item := range movies {
		if item.ID == params["id"] {
			movie = movies[index]
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func customErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not handled")
}
