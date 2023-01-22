package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastNam   string `json:"last_name"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.Handlefunc("/movies", getMovies).Methods("GET")
	r.Handlefunc("/movies/{id}", getMovies).Methods("Get")
	r.Handlefunc("/movies", createMovie).Methods("POST")
	r.Handlefunc("/movies/{id}", updateMovie).Methods("PUT")
	r.Handlefunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
