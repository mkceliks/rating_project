package main

import (
	"fmt"
	"log"
	"net/http"

	"rating_project/server"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/animes", server.GetAnimes).Methods("GET")
	r.HandleFunc("/anime-episodes", server.GetAnimeEpisodes).Methods("GET")
	r.HandleFunc("/movies", server.GetMovies).Methods("GET")
	r.HandleFunc("/movie-episodes", server.GetMovieEpisodes).Methods("GET")
	r.HandleFunc("/sports", server.GetSports).Methods("GET")
	r.HandleFunc("/sport-episodes", server.GetSportEpisodes).Methods("GET")
	r.HandleFunc("/sport-episodes/{id}", server.GetSportEpisodesById).Methods("GET")
	r.HandleFunc("/add-anime", server.AddAnime).Methods("POST")
	r.HandleFunc("/anime-episodes/add", server.AddAnimeEpisode).Methods("POST")
	r.HandleFunc("/add-movie", server.AddMovie).Methods("POST")
	r.HandleFunc("/movie-episodes/add", server.AddMovieEpisode).Methods("POST")
	r.HandleFunc("/add-sport", server.AddSport).Methods("POST")
	r.HandleFunc("/sport-episodes/add", server.AddSportEpisode).Methods("POST")

	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	handler := cors.Default().Handler(r)
	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
