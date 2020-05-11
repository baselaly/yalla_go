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

// Create function to create user
func (u *Repository) Create(user User) (int, error) {
	stmt, err := u.DB.Prepare("INSERT INTO `users` (first_name,last_name,email,password,image,cover)VALUES (?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	rs, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Image.String, user.Cover.String)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
