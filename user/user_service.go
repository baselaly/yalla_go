package user

import (
	"errors"
	"strconv"

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
func (u *Service) Register(c *gin.Context) (uint, error) {
	var user User
	email := c.PostForm("email")
	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	password := c.PostForm("password")

	err := user.SetPassword(password)
	if err != nil {
		return 0, err
	}
	image, imageHeader, _ := c.Request.FormFile("image")
	err = user.SetImage(image, imageHeader)
	if err != nil {
		return 0, err
	}

	cover, coverHeader, _ := c.Request.FormFile("cover")
	err = user.SetCover(cover, coverHeader)
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

	return uint(id), nil
}

// GetUser func to get single user
func (u *Service) GetUser(ID uint) (TransformedUser, error) {
	ByData := make(map[string]string)
	ByData["id"] = strconv.FormatUint(uint64(ID), 10)
	user, err := u.UserRepository.getUserBy(ByData)
	if err != nil {
		return TransformedUser{}, err
	}
	return user.TransformUser(), nil
}
