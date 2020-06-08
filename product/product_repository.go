package product

import (
	"database/sql"
	"errors"
)

// Repository struct to product repository
type Repository struct {
	DB *sql.DB
}

// ProvideProductRepository function to provide product repository
func ProvideProductRepository(DB *sql.DB) Repository {
	return Repository{DB: DB}
}

// Create function to create product in repo
func (repo *Repository) Create(p Product) error {
	stmt, err := repo.DB.Prepare("INSERT INTO `products` (name,description,user_id) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	rows, err := stmt.Exec(p.Name, p.Description, p.UserID)

	if err != nil {
		return err
	}

	_, err = rows.LastInsertId()

	if err != nil {
		return err
	}
	return nil
}

// GetProductBy function to get product
func (repo *Repository) GetProductBy(columns map[string]string) (Product, error) {
	var product Product
	var query string = "SELECT * FROM `products` WHERE 1=1"
	for key, value := range columns {
		query += " AND " + key + " = '" + value + "'"
	}
	query += " Limit 1;"
	err := repo.DB.QueryRow(query).Scan(&product.ID, &product.Name, &product.Description, &product.UserID)
	if err != nil {
		return Product{}, errors.New("Not Found")
	}
	return product, nil
}
