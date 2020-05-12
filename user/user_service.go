package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Service connect repository to api file
type Service struct {
	UserRepository Repository
}

// ProvideUserService provide service to repository
func ProvideUserService(Repo Repository) Service {
	return Service{UserRepository: Repo}
}

// Login function
func (u *Service) Login(email string, password string) (TransformedUser, error) {
	credintials := make(map[string]string)
	credintials["email"] = email
	user, err := u.UserRepository.getUserBy(credintials)
	if err != nil {
		return TransformedUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return TransformedUser{}, errors.New("Wrong Credentials")
	}
	Tuser := user.TransformUser()
	return Tuser, nil
}

// Register to register user
func (u *Service) Register(c *gin.Context) (int, error) {
	var user User
	email := c.PostForm("email")
	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	password := c.PostForm("password")
	image, imageHeader, err := c.Request.FormFile("image")
	if err != nil {
		return 0, err
	}

	cover, coverHeader, err := c.Request.FormFile("cover")
	if err != nil {
		return 0, err
	}

	err = user.SetCover(cover, coverHeader)
	if err != nil {
		return 0, err
	}
	err = user.SetImage(image, imageHeader)
	if err != nil {
		return 0, err
	}
	err = user.SetPassword(password)
	if err != nil {
		return 0, err
	}

	user.SetEmail(email)
	user.SetFirstName(firstname)
	user.SetLastName(lastname)

	id, err := u.UserRepository.Create(user)

	if err != nil {
		return 0, err
	}
	return id, nil
}
