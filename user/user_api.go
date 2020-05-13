package user

import (
	"net/http"
	"yalla_go/jwt"

	"github.com/gin-gonic/gin"
)

// API struct
type API struct {
	UserService Service
}

// ProvideUserAPI construct user api
func ProvideUserAPI(UserService Service) API {
	return API{UserService: UserService}
}

// Login API Function
func (UserAPI *API) Login(c *gin.Context) {
	validationErrors := ValidateLoginRequest(c)
	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrors})
		return
	}

	user, err := UserAPI.UserService.Login(c.PostForm("email"), c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	JwtToken, err := jwt.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": JwtToken})
}

// Register API Function
func (UserAPI *API) Register(c *gin.Context) {
	userID, err := UserAPI.UserService.Register(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	JwtToken, err := jwt.CreateToken(uint(userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": JwtToken})
}
