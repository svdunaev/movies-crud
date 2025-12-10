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

	fmt.Println(movies)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// TODO: modify to find by id
func getMovieById(id string) *Movie {
	for i := range movies {
		if movies[i].Title == id {
			return &movies[i]
		}
	}

	return nil
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
	movie := getMovieById(id)

	fmt.Println(movie)

	res.Header().Set("Content-Type", "application/json")
	fmt.Println("This is from get by id handler", movie)
	json.NewEncoder(res).Encode(movie)
}
