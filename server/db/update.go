package db

import (
	"context"
	"log"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateMovie(movie models.Movie) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	MovieCollection := DB.Database.Collection("movies")
	result, err := MovieCollection.UpdateOne(ctx, bson.M{"_id": movie.ID}, bson.M{"$set": movie})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result.ModifiedCount, err
}
