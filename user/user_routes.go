package user

import (
	"github.com/gin-gonic/gin"
)

// Routes generate routes for user
func Routes(route *gin.Engine, USERAPI API) {
	user := route.Group("/user")
	user.POST("/login", USERAPI.Login)
	user.POST("/register", USERAPI.Register)
}
