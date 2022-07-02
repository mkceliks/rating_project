package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Anime struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rating string             `bson:"rating,omitempty" json:"rating"`
	Title  string             `bson:"title,omitempty" json:"title"`
}

type Anime_episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Anime       primitive.ObjectID `bson:"anime,omitempty" json:"anime_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	Duration    int32              `bson:"duration,omitempty" json:"duration"`
}

type Movie struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rating string             `bson:"rating,omitempty" json:"rating"`
	Title  string             `bson:"title,omitempty" json:"title"`
}

type Movie_episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Movie       primitive.ObjectID `bson:"movie,omitempty" json:"movie_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	Duration    int32              `bson:"duration,omitempty" json:"duration"`
}

type Sport struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rating string             `bson:"rating,omitempty" json:"rating"`
	Title  string             `bson:"title,omitempty" json:"title"`
}

type Sport_episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Sport       primitive.ObjectID `bson:"sport,omitempty" json:"sport_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	Duration    int32              `bson:"duration,omitempty" json:"duration"`
}

var animes []Anime
var anime_episodes []Anime_episode
var movies []Movie
var movie_episodes []Movie_episode
var sports []Sport
var sport_episodes []Sport_episode

func InitiateMongoClient() *mongo.Client {

	var err error
	var client *mongo.Client
	uri := "mongodb+srv://mkceliks:159987741@loginsystem.0j8qw.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

func getAnimes(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	animeCollection := letmewatch.Collection("animes")

	cursor, err := animeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &animes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(animes)

}

func getMovies(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	movieCollection := letmewatch.Collection("movies")

	cursor, err := movieCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movies)

}

func getSports(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	sportCollection := letmewatch.Collection("sports")

	cursor, err := sportCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &sports); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(sports)

}

func getAnimeEpisodes(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	animeEpisodeCollection := letmewatch.Collection("anime-episodes")

	cursor, err := animeEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &anime_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(anime_episodes)

}

func getMovieEpisodes(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	movieEpisodeCollection := letmewatch.Collection("movie-episodes")

	cursor, err := movieEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &movie_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie_episodes)

}

func getSportEpisodes(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	sportEpisodeCollection := letmewatch.Collection("sport-episodes")

	cursor, err := sportEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &sport_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(sport_episodes)

}

func getSportEpisodesById(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	sportEpisodeCollection := letmewatch.Collection("sport-episodes")
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := sportEpisodeCollection.Find(ctx, bson.M{"sport_id": objId})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &sport_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(sport_episodes)

}

func addAnime(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	animeCollection := letmewatch.Collection("animes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var anime Anime
		_ = json.NewDecoder(r.Body).Decode(&anime)

		animeResult, err := animeCollection.InsertOne(ctx, bson.D{
			{"rating", anime.Rating},
			{"title", anime.Title},
		})
		json.NewEncoder(w).Encode(animeResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	movieCollection := letmewatch.Collection("movies")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)

		movieResult, err := movieCollection.InsertOne(ctx, bson.D{
			{"rating", movie.Rating},
			{"title", movie.Title},
		})
		json.NewEncoder(w).Encode(movieResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addSport(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	sportCollection := letmewatch.Collection("sports")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var sport Sport
		_ = json.NewDecoder(r.Body).Decode(&sport)

		sportResult, err := sportCollection.InsertOne(ctx, bson.D{
			{"rating", sport.Rating},
			{"title", sport.Title},
		})
		json.NewEncoder(w).Encode(sportResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addAnimeEpisode(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	animeEpisodeCollection := letmewatch.Collection("anime-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var anime_episode Anime_episode
		_ = json.NewDecoder(r.Body).Decode(&anime_episode)

		animeEpisodeResult, err := animeEpisodeCollection.InsertOne(ctx, bson.D{
			{"anime_id", anime_episode.Anime},
			{"title", anime_episode.Title},
			{"description", anime_episode.Description},
			{"duration", anime_episode.Duration},
		})
		json.NewEncoder(w).Encode(animeEpisodeResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addMovieEpisode(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	movieEpisodeCollection := letmewatch.Collection("movie-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var movie_episode Movie_episode
		_ = json.NewDecoder(r.Body).Decode(&movie_episode)

		movieEpisodeResult, err := movieEpisodeCollection.InsertOne(ctx, bson.D{
			{"movie_id", movie_episode.Movie},
			{"title", movie_episode.Title},
			{"description", movie_episode.Description},
			{"duration", movie_episode.Duration},
		})
		json.NewEncoder(w).Encode(movieEpisodeResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addSportEpisode(w http.ResponseWriter, r *http.Request) {

	conn := InitiateMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	letmewatch := conn.Database("letmewatch")
	sportEpisodeCollection := letmewatch.Collection("sport-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var sport_episode Sport_episode
		_ = json.NewDecoder(r.Body).Decode(&sport_episode)

		sportEpisodeResult, err := sportEpisodeCollection.InsertOne(ctx, bson.D{
			{"sport_id", sport_episode.Sport},
			{"title", sport_episode.Title},
			{"description", sport_episode.Description},
			{"duration", sport_episode.Duration},
		})
		json.NewEncoder(w).Encode(sportEpisodeResult)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/animes", getAnimes).Methods("GET")
	r.HandleFunc("/anime-episodes", getAnimeEpisodes).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie-episodes", getMovieEpisodes).Methods("GET")
	r.HandleFunc("/sports", getSports).Methods("GET")
	r.HandleFunc("/sport-episodes", getSportEpisodes).Methods("GET")
	r.HandleFunc("/sport-episodes/{id}", getSportEpisodesById).Methods("GET")
	r.HandleFunc("/add-anime", addAnime).Methods("POST")
	r.HandleFunc("/anime-episodes/add", addAnimeEpisode).Methods("POST")
	r.HandleFunc("/add-movie", addMovie).Methods("POST")
	r.HandleFunc("/movie-episodes/add", addMovieEpisode).Methods("POST")
	r.HandleFunc("/add-sport", addSport).Methods("POST")
	r.HandleFunc("/sport-episodes/add", addSportEpisode).Methods("POST")

	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	handler := cors.Default().Handler(r)
	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
