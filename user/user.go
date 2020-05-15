package user

import (
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"yalla_go/nullsql"

	"golang.org/x/crypto/bcrypt"
)

// User struct for user model
type User struct {
	ID        uint               `json:"id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Image     nullsql.NullString `json:"image"`
	Cover     nullsql.NullString `json:"cover"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
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
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
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
func (u *User) SetImage(file multipart.File, handle *multipart.FileHeader) error {

	// handle if there is no image sent, user image is optional
	if file == nil && handle == nil {
		u.Cover = nullsql.ToNullString("")
		return nil
	}

	// upload user image if sent
	filename := RandomString(20) + "*.jpg"
	defer file.Close()

	tempFile, err := ioutil.TempFile("uploads/users", filename)
	if err != nil {
		return err
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	tempFile.Write(fileBytes)

	u.Image = nullsql.ToNullString(tempFile.Name())
	return nil
}

// GetImage to get image for user
func (u User) GetImage() string {
	image := u.Image
	if image.Valid && len(image.String) > 0 {
		return "http://localhost:8080/" + image.String
	}
	return "http://localhost:8080/uploads/users/profile.png"
}

// SetCover to upload user image and set the name
func (u *User) SetCover(file multipart.File, handle *multipart.FileHeader) error {
	// handle if there is no cover sent, user cover is optional
	if file == nil && handle == nil {
		u.Cover = nullsql.ToNullString("")
		return nil
	}

	// upload user cover if sent
	filename := RandomString(20) + "*.jpg"
	defer file.Close()

	tempFile, err := ioutil.TempFile("uploads/users", filename)
	if err != nil {
		return err
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	tempFile.Write(fileBytes)

	u.Cover = nullsql.ToNullString(tempFile.Name())
	return nil
}

// GetCover to get image for user
func (u User) GetCover() string {
	image := u.Cover
	if image.Valid && len(image.String) > 0 {
		return "http://localhost:8080/" + image.String
	}
	return "http://localhost:8080/uploads/users/cover.png"
}
