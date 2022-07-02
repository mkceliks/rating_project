package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
