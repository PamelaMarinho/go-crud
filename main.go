package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	return
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
	return
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Movie 1", Director: &Director{Firstname: "maria", Lastname: "Auhs"}})
	movies = append(movies, Movie{ID: "2", Isbn: "11124", Title: "Movie 2", Director: &Director{Firstname: "Jhonny", Lastname: "Deep"}})
	r.HandleFunc("/movie", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
