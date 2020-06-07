package product

import "database/sql"

// Repository struct to product repository
type Repository struct {
	DB *sql.DB
}

// ProvideProductRepository function to provide product repository
func ProvideProductRepository(DB *sql.DB) Repository {
	return Repository{DB: DB}
}

// Create function to create product in repo
func (repo Repository) Create(p Product) error {
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
