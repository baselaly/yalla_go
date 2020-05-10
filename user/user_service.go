package user

import (
	"errors"

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
func (u *Service) Login(email string, password string) (User, error) {
	credintials := make(map[string]string)
	credintials["email"] = email
	user, err := u.UserRepository.getUserBy(credintials)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("Wrong Credentials")
	}

	return user, nil
}
