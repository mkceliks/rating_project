package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Anime struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rating string             `bson:"rating,omitempty" json:"rating"`
	Title  string             `bson:"title,omitempty" json:"title"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Anime       primitive.ObjectID `bson:"anime,omitempty" json:"anime_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	Duration    int32              `bson:"duration,omitempty" json:"duration"`
}

var animes []Anime
var episodes []Episode

func getAnimes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mkceliks:159987741@loginsystem.0j8qw.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	animeDb := client.Database("animes")
	animeCollection := animeDb.Collection("anime")

	cursor, err := animeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// /////////////////////////////////////////// SEARCHING ALL THE DB
	// var animes []Anime
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&animes); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(animes[0].Title)
	// }
	// ///////////////////////////////////////////

	/////////////////////////////////////////  SEARCHING ALL THE DB

	if err = cursor.All(ctx, &animes); err != nil {
		log.Fatal(err)
	}
	/////////////////////////////////////////
	json.NewEncoder(w).Encode(animes)
}

func getEpisodes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mkceliks:159987741@loginsystem.0j8qw.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	animeDb := client.Database("animes")
	episodeCollection := animeDb.Collection("episodes")

	cursor, err := episodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// /////////////////////////////////////////// SEARCHING ALL THE DB
	// var animes []Anime
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&animes); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(animes[0].Title)
	// }
	// ///////////////////////////////////////////

	/////////////////////////////////////////  SEARCHING ALL THE DB

	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	/////////////////////////////////////////
	json.NewEncoder(w).Encode(episodes)
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

// func createMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var movie Movie
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID = strconv.Itoa(rand.Intn(10000000))
// 	movies = append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
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

	// /////////////////////////////////////////// INSERTING TO DB WITH InsertOne
	// animeResult, err := animeCollection.InsertOne(ctx, bson.D{
	// 	{"Title", "Attack On Titan"},
	// 	{"Rating", "10.0"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(animeResult.InsertedID)

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
	r.HandleFunc("/movies", getAnimes).Methods("GET")
	r.HandleFunc("/episodes", getEpisodes).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
