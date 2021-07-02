package routes

import (
	auth "server/controllers/auth"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func (r routes) userRouter(rg *gin.RouterGroup) {

	rg.POST("/signup", auth.SignUp)
	rg.POST("/signin", auth.SignIn)
	rg.POST("/logout", middleware.TokenAuthMiddleware(), auth.Logout)
	rg.POST("/token/refresh", middleware.TokenAuthMiddleware(), auth.Refresh)
}
