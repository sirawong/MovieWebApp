package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMovie(ID primitive.ObjectID) (int, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	MovieCollection := DB.Database.Collection("movies")
	result, err := MovieCollection.DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	return int(result.DeletedCount), nil

}
