package user

import (
	"yalla_go/middleware"

	"github.com/gin-gonic/gin"
)

// Routes generate routes for user
func Routes(route *gin.Engine, USERAPI API) {
	public := route.Group("/user")
	private := route.Group("/user")

	private.Use(middleware.Auth())
	public.POST("/login", USERAPI.Login)
	public.POST("/register", USERAPI.Register)
	private.GET("/myprofile", USERAPI.GetProfile)
}
