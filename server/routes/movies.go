package routes

import (
	controllerMovies "server/controllers/movies"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func (r routes) movieRouter(rg *gin.RouterGroup) {
	rg.GET("/search", controllerMovies.GetMovies)
	rg.GET("/", controllerMovies.GetMovies)
	rg.GET("/:id", controllerMovies.GetMovie)

	rg.POST("/", middleware.TokenAuthMiddleware(), controllerMovies.CreateMovie)
	rg.PATCH("/:id", controllerMovies.UpdateMovie)
	rg.DELETE("/:id", middleware.TokenAuthMiddleware(), controllerMovies.DeleteMovie)
	// rg.PATCH("/:id/likePost", middleware.TokenAuthMiddleware(), likePost)
}
