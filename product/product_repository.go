package product

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func ProvideProductRepository(DB *sql.DB) Repository {
	return Repository{DB: DB}
}

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
