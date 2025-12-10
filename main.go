package main

import (
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	title    string
	director string
	id       int64
	genre    string
}

var movies []Movie

func main() {
	fmt.Println("Start server")
	defer fmt.Println("Stop server")

	movies = append(movies, Movie{
		title:    "Terminator",
		director: "James Cameron",
		id:       1,
		genre:    "Sci-Fi",
	})

	fmt.Println(movies)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
