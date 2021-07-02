package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Creator          primitive.ObjectID `json:"creator,omitempty" bson:"creator,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`
	Plot             string             `json:"plot,omitempty" bson:"plot,omitempty"`
	Directors        []string           `json:"directors,omitempty" bson:"directors,omitempty"`
	Cast             []string           `json:"cast,omitempty" bson:"cast,omitempty"`
	Writers          []string           `json:"writers,omitempty" bson:"writers,omitempty"`
	Genres           []string           `json:"genres,omitempty" bson:"genres,omitempty"`
	Rated            string             `json:"rated,omitempty" bson:"rated,omitempty"`
	Released         time.Time          `json:"released,omitempty" bson:"released,omitempty"`
	Year             interface{}        `json:"year,omitempty" bson:"year,omitempty"`
	Awards           Awards             `json:"Awards,omitempty" bson:"Awards,omitempty"`
	Imdb             Imdb               `json:"imdb,omitempty" bson:"imdb,omitempty"`
	Poster           string             `json:"poster,omitempty" bson:"poster,omitempty,omitempty"`
	Languages        []string           `json:"languages,omitempty" bson:"languages,omitempty"`
	NumMflixComments int                `json:"num_mflix_comments,omitempty" bson:"num_mflix_comments,omitempty"`
	CreatedAt        time.Time          `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	Lastupdate       interface{}        `json:"lastupdate,omitempty" bson:"lastupdate,omitempty"`
}

type Awards struct {
	Wins        int    `json:"wins" bson:"wins,omitempty"`
	Nominations int    `json:"nominations" bson:"nominations,omitempty"`
	Text        string `json:"text" bson:"text,omitempty"`
}

type Imdb struct {
	Rating float32 `json:"rating" bson:"rating,omitempty"`
	Votes  uint32  `json:"votes" bson:"votes,omitempty"`
	Id     uint32  `json:"id" bson:"id,omitempty"`
}
