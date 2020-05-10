package user

import (
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

// User struct for user model
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Cover     string `json:"cover"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// SetEmail to set email in model
func (u *User) SetEmail(email string) {
	u.Email = email
}

// GetEmail get email in model
func (u User) GetEmail() string {
	return u.Email
}

// SetFirstName to set firstname in model
func (u *User) SetFirstName(firstname string) {
	u.FirstName = firstname
}

// GetFirstName get firstname in model
func (u User) GetFirstName() string {
	return u.FirstName
}

// SetLastName to set firstname in model
func (u *User) SetLastName(lastname string) {
	u.LastName = lastname
}

// GetLastName get firstname in model
func (u User) GetLastName() string {
	return u.LastName
}

// SetPassword to set password in model
func (u *User) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalln(err)
	}
	u.Password = string(hash)
}

// RandomString Generate a random string of A-Z chars with len = l
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

// RandomInt Returns an int >= min, < max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// SetImage to upload user image and set the name
func (u *User) SetImage(file multipart.File, handle *multipart.FileHeader) {
	filename := RandomString(20)
	path := "/uploads/users/" + filename

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(path, data, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	u.Image = filename
}

// GetImage to get image for user
func (u User) GetImage() string {
	return "http://localhost:8080/uploads/users/" + u.Image
}

// SetCover to upload user image and set the name
func (u *User) SetCover(file multipart.File, handle *multipart.FileHeader) {
	filename := RandomString(20)
	path := "/uploads/users/" + filename

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(path, data, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	u.Cover = filename
}

// GetCover to get image for user
func (u User) GetCover() string {
	return "http://localhost:8080/uploads/users/" + u.Image
}
