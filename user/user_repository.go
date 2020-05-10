package user

import (
	"database/sql"
)

// Repository connect model to service
type Repository struct {
	DB *sql.DB
}

// ProvideUserRepository construct repository
func ProvideUserRepository(DB *sql.DB) Repository {
	return Repository{DB: DB}
}

// getUserBy function to get user by columns
func (u *Repository) getUserBy(columns map[string]string) (User, error) {
	var user User
	var query string = "SELECT * FROM `users` WHERE 1=1"
	for key, value := range columns {
		query += " AND " + key + " = '" + value + "'"
	}
	query += " Limit 1;"
	err := u.DB.QueryRow(query).Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Image, &user.Cover, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
