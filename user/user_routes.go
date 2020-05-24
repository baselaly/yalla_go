package user

import (
	"github.com/gin-gonic/gin"
)

// Routes generate routes for user
func Routes(route *gin.Engine, USERAPI API) {
	public := route.Group("/user")
	private := route.Group("/user")
	public.POST("/login", USERAPI.Login)
	public.POST("/register", USERAPI.Register)
	private.GET("/myprofile", USERAPI.GetProfile)
}
