package main

import (
	"encoding/json"
	"fmt"
	"go-movies-crud/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at Port %v...\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetMovies())

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie := db.GetMovie(params["id"])
	if movie != nil {
		json.NewEncoder(w).Encode(*movie)
		return
	}
	json.NewEncoder(w).Encode("No Movie found")

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie db.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = db.GetNextID()
	db.AddMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie db.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	go db.UpdateMovie(params["id"], &movie)
	json.NewEncoder(w).Encode(movie)

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db.DeleteMovie(params["id"])
}
