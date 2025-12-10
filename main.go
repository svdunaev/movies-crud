package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Movie struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Id       string `json:"id"`
	Genre    string `json:"genre"`
}

var movies = []Movie{
	{
		Title:    "Terminator",
		Director: "James Cameron",
		Id:       uuid.NewString(),
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

	fmt.Println(movies)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getMovieById(id string) (idx int, movie *Movie) {
	for i := range movies {
		if movies[i].Id == id {
			return i, &movies[i]
		}
	}

	return -1, nil
}

func getMoviesHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from get handler", movies)
	json.NewEncoder(res).Encode(movies)
}

func getMovieByIdHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]
	_, movie := getMovieById(id)

	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from get by id handler", movie)
	json.NewEncoder(res).Encode(movie)
}

// todo: add duplicate validation
func postMovieHandler(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var movie Movie
	err := decoder.Decode(&movie)
	if err != nil {
		panic(err)
	}

	movie.Id = uuid.NewString()

	movies = append(movies, movie)

	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from post handler", movies)
	json.NewEncoder(res).Encode(movies)
}

func deleteMovieHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	idx, _ := getMovieById(id)
	if idx >= 0 {
		movies = slices.Delete(movies, idx, idx+1)
	} else {
		log.Fatal("this movie does not exist")
	}

	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from delete handler", movies)
	json.NewEncoder(res).Encode(movies)
}

func updateMovieHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := parts[2]

	idx, movie := getMovieById(id)

	var updated Movie
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&updated)
	if err != nil {
		panic(err)
	}

	updated.Id = movie.Id

	movies[idx] = updated

	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from put handler", movies)
	json.NewEncoder(res).Encode(movies)
}
