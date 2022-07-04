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

	r.HandleFunc("/animes", server.GetAnimes).Methods("GET")                         // gets all animes
	r.HandleFunc("/anime-episodes", server.GetAnimeEpisodes).Methods("GET")          // gets all anime episodes
	r.HandleFunc("/anime-episodes/{id}", server.GetAnimeEpisodesById).Methods("GET") // gets anime episodes by anime_id
	r.HandleFunc("/add-anime", server.AddAnime).Methods("POST")                      // adds anime into anime collection
	r.HandleFunc("/delete-anime/{id}", server.DeleteAnime).Methods("DELETE")         // deletes the selected anime and it's episodes from collections
	r.HandleFunc("/anime-episodes/add", server.AddAnimeEpisode).Methods("POST")      // adds anime episode into anime_episodes collection

	r.HandleFunc("/movies", server.GetMovies).Methods("GET")                         // gets all movies
	r.HandleFunc("/movie-episodes", server.GetMovieEpisodes).Methods("GET")          // gets all movie episodes
	r.HandleFunc("/movie-episodes/{id}", server.GetMovieEpisodesById).Methods("GET") // gets movie episodes by movie_id
	r.HandleFunc("/add-movie", server.AddMovie).Methods("POST")                      // adds movie into movie collection
	r.HandleFunc("/movie-episodes/add", server.AddMovieEpisode).Methods("POST")      // adds movie episode into movie_episodes collection

	r.HandleFunc("/sports", server.GetSports).Methods("GET")                         // gets all sport shows
	r.HandleFunc("/sport-episodes", server.GetSportEpisodes).Methods("GET")          // gets all ssport shows episodes
	r.HandleFunc("/sport-episodes/{id}", server.GetSportEpisodesById).Methods("GET") // gets sport show's all episodes by sport_id
	r.HandleFunc("/add-sport", server.AddSport).Methods("POST")                      // adds sport show
	r.HandleFunc("/sport-episodes/add", server.AddSportEpisode).Methods("POST")      // adds sport episode into sport_episodes collection

	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	handler := cors.AllowAll().Handler(r)
	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
