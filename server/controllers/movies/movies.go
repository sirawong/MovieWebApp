package controllerMovies

import (
	"log"
	"net/http"
	"server/db"
	"server/middleware"
	"server/models"
	"strings"
	"time"

	controllersAuth "server/controllers/auth"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMovies(c *gin.Context) {
	limit := 20
	// pages := c.Param("page")
	// page, _ := strconv.Atoi(pages)
	page := 1
	title := c.Query("searchQuery")
	genres := c.Query("genres")
	genreA := strings.Split(genres, ",")
	var filter = bson.M{}
	if genres != "" || title != "" {
		filter = bson.M{"$or": []bson.M{bson.M{"title": title}, bson.M{"genres": bson.M{"$in": genreA}}}}
	}
	startIndex := (page - 1) * limit
	totalMovies, err := db.CountDocuments()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	movies, err := db.GetMovies(limit, startIndex, filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": movies, "currentPage": page, "numberOfPages": int(totalMovies) / limit})
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	filter := bson.D{{"_id", oid}}
	movie, err := db.GetMovie(filter)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func getUserId(c *gin.Context) (string, error) {
	tokenAuth, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.Println("ExtractTokenMetadata")
		return "", err
	}
	userId, err := controllersAuth.FetchAuth(tokenAuth)
	if err != nil {
		log.Println("FetchAuth")
		return "", err
	}
	return userId, nil
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	log.Println(userId)

	l, _ := time.LoadLocation("Asia/Bangkok")
	movie.CreatedAt = time.Now().In(l)

	oid, _ := primitive.ObjectIDFromHex(userId)
	movie.Creator = oid

	inserted, err := db.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": inserted})

}

func UpdateMovie(c *gin.Context) {
	movieId := c.Param("id")
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	movie.ID, _ = primitive.ObjectIDFromHex(movieId)
	l, _ := time.LoadLocation("Asia/Bangkok")
	movie.Lastupdate = time.Now().In(l)

	result, err := db.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, gin.H{"ModifiedCount": result})

}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	ID, _ := primitive.ObjectIDFromHex(id)
	result, err := db.DeleteMovie(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
	}

	c.JSON(http.StatusOK, gin.H{"DeletedCount": result})

}
