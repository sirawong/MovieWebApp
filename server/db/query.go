package db

import (
	"context"
	"fmt"
	"log"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func QueryUser(username string) (models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var user models.User
	userCollection := DB.Database.Collection("users")
	err := userCollection.FindOne(ctx, bson.M{"email": username}).Decode(&user)

	return user, err
}

func CountDocuments() (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	MovieCollection := DB.Database.Collection("movies")
	itemCount, err := MovieCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	return itemCount, err
}

func GetMovies(limit, startIndex int, filter interface{}) ([]models.Movie, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	MovieCollection := DB.Database.Collection("movies")

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(startIndex))

	cursor, err := MovieCollection.Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []models.Movie
	for cursor.Next(ctx) {
		var movie models.Movie
		if err = cursor.Decode(&movie); err != nil {
			log.Println(err)
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func GetMovie(filter interface{}) (models.Movie, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	MovieCollection := DB.Database.Collection("movies")

	var Movie models.Movie
	if err := MovieCollection.FindOne(ctx, filter).Decode(&Movie); err != nil {
		log.Println(err)
		return models.Movie{}, err
	}
	return Movie, nil
}
