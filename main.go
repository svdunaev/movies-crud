package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Movie struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
}

var movies = map[string]Movie{
	uuid.NewString(): {
		Title:    "Terminator",
		Director: "James Cameron",
		Genre:    "Sci-Fi",
	},
}

func main() {
	fmt.Println("Start server")
	defer fmt.Println("Stop server")
	http.HandleFunc("GET /movies", getMoviesHandler)
	http.HandleFunc("GET /movies/{id}", getMovieByIdHandler)
	http.HandleFunc("POST /movies", postMovieHandler)
	http.HandleFunc("DELETE /movies/{id}", deleteMovieHandler)
	http.HandleFunc("PUT /movies/{id}", updateMovieHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getMoviesHandler(res http.ResponseWriter, req *http.Request) {
	if len(movies) < 1 {
		fmt.Fprintf(res, "There are no movies in this list, you can add the first one")
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovieByIdHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]
	movie, ok := movies[id]
	if !ok {
		http.Error(res, "No such movie exist", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movie)
}

func postMovieHandler(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var movie Movie
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}

	if movieExists(movies, movie.Title) {
		http.Error(res, "This movie already exists", http.StatusConflict)
	} else {
		movies[uuid.NewString()] = movie
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func deleteMovieHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	_, ok := movies[id]
	if ok {
		delete(movies, id)
	} else {
		http.Error(res, "No such movie exist", http.StatusNotFound)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func updateMovieHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	var updated Movie
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&updated)
	if err != nil {
		http.Error(res, "Bad reques", http.StatusBadRequest)
	}

	_, ok := movies[id]

	if !ok {
		http.Error(res, "No such movie exist", http.StatusNotFound)
	}

	movies[id] = updated

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func movieExists(movies map[string]Movie, title string) bool {
	for _, movie := range movies {
		if movie.Title == title {
			return true
		}
	}

	return false
}
