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

// func deleteMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(movies)
// }

// func getMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range movies {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// }

// func updateMovie(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var movie Movie
// 			_ = json.NewDecoder(r.Body).Decode(&movie)
// 			movie.ID = params["id"]
// 			movies = append(movies, movie)
// 			json.NewEncoder(w).Encode(movie)
// 		}
// 	}

// }

// func createMovie(w http.ResponseWriter, r *http.Request) {

// }

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

func main() {

	// /////////////////////////////////////////// SEARCHING ALL THE DB AND RETURNS ONE WITH FindOne
	// var anime bson.M
	// if err = animeCollection.FindOne(ctx, bson.M{}).Decode(&anime); err != nil {
	// 	log.Fatal(err)
	// }
	// //fmt.Println(anime)
	// ///////////////////////////////////////////

	// /////////////////////////////////////////// FILTERING THE DB WITH Cursor and Find
	// filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var episodesFiltered []bson.M
	// if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	// //fmt.Println(episodesFiltered)
	// ///////////////////////////////////////////

	// /////////////////////////////////////////// SORTING THE DB WITH SetSort
	// opts := options.Find()
	// opts.SetSort(bson.D{{"duration", 1}})
	// sortCursor, err := episodesCollection.Find(ctx, bson.D{
	// 	{"duration", bson.D{
	// 		{"$gt", 26},
	// 	}},
	// }, opts)
	// var episodesSorted []bson.M
	// if err = sortCursor.All(ctx, &episodesSorted); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episodesSorted)

	/////////////////////////////////////////// INSERTING TO DB WITH InsertMany
	// episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
	// 	bson.D{
	// 		{"anime", animeResult.InsertedID},
	// 		{"title", "Episode #1"},
	// 		{"description", "This is the first episode."},
	// 		{"duration", 25},
	// 	},
	// 	bson.D{
	// 		{"anime", animeResult.InsertedID},
	// 		{"title", "Episode #2"},
	// 		{"description", "This is the second episode."},
	// 		{"duration", 32},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episodeResult.InsertedIDs)

	/////////////////////////////////////////// MY CODES
	r := mux.NewRouter()
	r.HandleFunc("/animes", getAnimes).Methods("GET")
	r.HandleFunc("/anime-episodes", getAnimeEpisodes).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/sports", getSports).Methods("GET")
	r.HandleFunc("/add-anime", addAnime).Methods("POST")
	r.HandleFunc("/add-movie", addMovie).Methods("POST")
	r.HandleFunc("/add-sport", addSport).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	handler := cors.Default().Handler(r)
	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
