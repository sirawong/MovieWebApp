package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func SetupRouter() routes {
	r := routes{
		router: gin.Default(),
	}

	usersv1 := r.router.Group("/v1/users")
	r.userRouter(usersv1)

	moviesv1 := r.router.Group("/v1")
	r.movieRouter(moviesv1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr[0])
}
