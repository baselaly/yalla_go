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
	public.GET("/profile/:id", USERAPI.GetUserProfile)
	private.GET("/myprofile", USERAPI.GetProfile)
}
