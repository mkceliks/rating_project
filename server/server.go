package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"rating_project/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var animes []models.Anime
var anime_episodes []models.Anime_episode
var movies []models.Movie
var movie_episodes []models.Movie_episode
var sports []models.Sport
var sport_episodes []models.Sport_episode

func InitiateMongoClient() *mongo.Client {

	var err error
	var client *mongo.Client
	uri := "mongodb://localhost:27017"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

var conn = InitiateMongoClient()

var letmewatch = conn.Database("letmewatch")

func GetAnimes(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	animeCollection := letmewatch.Collection("animes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := animeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &animes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(animes)
	log.Printf("Binomial took %s", time.Since(start))

}

func GetMovies(w http.ResponseWriter, r *http.Request) {

	movieCollection := letmewatch.Collection("movies")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := movieCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movies)

}

func GetSports(w http.ResponseWriter, r *http.Request) {

	sportCollection := letmewatch.Collection("sports")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := sportCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &sports); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(sports)

}

func GetAnimeEpisodes(w http.ResponseWriter, r *http.Request) {

	animeEpisodeCollection := letmewatch.Collection("anime-episodes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := animeEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &anime_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(anime_episodes)

}

func GetMovieEpisodes(w http.ResponseWriter, r *http.Request) {

	movieEpisodeCollection := letmewatch.Collection("movie-episodes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := movieEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &movie_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie_episodes)

}

func GetSportEpisodes(w http.ResponseWriter, r *http.Request) {

	sportEpisodeCollection := letmewatch.Collection("sport-episodes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := sportEpisodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &sport_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(sport_episodes)

}

func GetAnimeEpisodesById(w http.ResponseWriter, r *http.Request) {

	animeEpisodeCollection := letmewatch.Collection("anime-episodes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := animeEpisodeCollection.Find(ctx, bson.M{"anime_id": objId})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &anime_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(anime_episodes)

}

func GetMovieEpisodesById(w http.ResponseWriter, r *http.Request) {

	movieEpisodeCollection := letmewatch.Collection("movie-episodes")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := movieEpisodeCollection.Find(ctx, bson.M{"movie_id": objId})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &movie_episodes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie_episodes)

}

func GetSportEpisodesById(w http.ResponseWriter, r *http.Request) {

	sportEpisodeCollection := letmewatch.Collection("sport-episodes")

	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
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

func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	animesCollection := letmewatch.Collection("animes")
	animeEpisodeCollection := letmewatch.Collection("anime-episodes")

	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	animesCollection.DeleteOne(ctx, bson.M{"_id": objId})
	animeEpisodeCollection.DeleteMany(ctx, bson.M{"anime_id": objId})

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	moviesCollection := letmewatch.Collection("movies")
	movieEpisodeCollection := letmewatch.Collection("movie-episodes")

	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	moviesCollection.DeleteOne(ctx, bson.M{"_id": objId})
	movieEpisodeCollection.DeleteMany(ctx, bson.M{"movie_id": objId})

}

func DeleteSport(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	sportsCollection := letmewatch.Collection("sports")
	sportEpisodeCollection := letmewatch.Collection("sport-episodes")

	id := mux.Vars(r)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	sportsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	sportEpisodeCollection.DeleteMany(ctx, bson.M{"sport_id": objId})

}

func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	animeCollection := letmewatch.Collection("animes")
	if r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		var anime models.Anime
		_ = json.NewDecoder(r.Body).Decode(&anime)
		id := mux.Vars(r)["id"]
		objId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{"_id", objId}}
		result, err := animeCollection.UpdateOne(ctx, filter, bson.D{
			{"rating", anime.Rating},
			{"title", anime.Title},
		})
		json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func AddAnime(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	animeCollection := letmewatch.Collection("animes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var anime models.Anime
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

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	movieCollection := letmewatch.Collection("movies")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var movie models.Movie
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

func AddSport(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	sportCollection := letmewatch.Collection("sports")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var sport models.Sport
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

func AddAnimeEpisode(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	animeEpisodeCollection := letmewatch.Collection("anime-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var anime_episode models.Anime_episode
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

func AddMovieEpisode(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	movieEpisodeCollection := letmewatch.Collection("movie-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var movie_episode models.Movie_episode
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

func AddSportEpisode(w http.ResponseWriter, r *http.Request) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	sportEpisodeCollection := letmewatch.Collection("sport-episodes")

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var sport_episode models.Sport_episode
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
