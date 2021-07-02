package controllerAuth

import (
	"context"
	"log"
	"net/http"
	"server/db"
	"server/middleware"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	userDB, err := db.QueryUser(u.Email)
	checkLogin := CheckPasswordHash(u.Password, userDB.Password)
	if err != nil || !checkLogin {
		c.JSON(http.StatusUnauthorized, "Username or Password incorrect")
		return
	}

	token, err := CreateToken(userDB.ID.Hex())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(userDB.ID.Hex(), token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func SignUp(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	userDB, _ := db.QueryUser(u.Email)
	if userDB.Email != "" {
		c.JSON(http.StatusNotAcceptable, "That username is taken, Try another.")
		return
	}

	hashpass, err := HashPassword(u.Password)
	if err != nil {
		log.Print(err)
	}

	u.Password = hashpass
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.DB.Database.Collection("users")
	result, err := userCollection.InsertOne(ctx, u)
	if err != nil {
		log.Print(err)
	}
	oid := result.InsertedID.(primitive.ObjectID)

	token, err := CreateToken(oid.Hex())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusCreated, token)
}

func Logout(c *gin.Context) {
	au, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
