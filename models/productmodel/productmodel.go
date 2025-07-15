package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT
		products.id,
		products.name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at,
		categories.id AS category_id,
		categories.name AS category_name
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Category.Id,
			&product.Category.Name,
		)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products	
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec("INSERT INTO products (name, stock, description, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		product.Name, product.Stock, product.Description, product.Category.Id, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}

func GetById(id int) *entities.Product {
	var product entities.Product
	err := config.DB.QueryRow(`
		SELECT
		products.id,
		products.name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at,
		categories.id AS category_id,
		categories.name AS category_name
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = ?`, id).Scan(
			&product.Id,
			&product.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Category.Id,
			&product.Category.Name,
	)
	if err != nil {
		return nil
	}
	return &product
}


func Update(product entities.Product) bool {
	result, err := config.DB.Exec("UPDATE products SET name = ?, stock = ?, description = ?, category_id = ?, updated_at = ? WHERE id = ?",
		product.Name, product.Stock, product.Description, product.Category.Id, product.UpdatedAt, product.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

func Delete(id int) bool {
	result, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}