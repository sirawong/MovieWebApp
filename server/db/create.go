package db

import (
	"context"
	"log"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMovie(movie models.Movie) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	MovieCollection := DB.Database.Collection("movies")
	result, err := MovieCollection.InsertOne(ctx, movie)
	if err != nil {
		log.Println(err)
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID).Hex()
	return oid, nil
}
