package models

import (
	"log"

	"github.com/gothello/go-web-studies/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.SqlConnection()

	rows, err := db.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}
	products := []Product{}

	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err := rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Fatal(err)
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()

	return products
}

func GetProduct(id int) Product {
	db := db.SqlConnection()

	sql := "SELECT * FROM products WHERE id = $1"

	row, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}

	for row.Next() {
		err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return product
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.SqlConnection()

	sql := "UPDATE products SET name = $2, description = $3, price = $4, quantity = $4 WHERE id = $1;"

	_, err := db.Exec(
		sql,
		id,
		name,
		description,
		price,
		quantity,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func SaveProduct(id int, name, description string, price float64, quantity int) {
	db := db.SqlConnection()

	sql := "INSERT INTO products (id, name, description, price, quantity) VALUES ($1, $2, $3, $4, $5);"

	_, err := db.Exec(sql, id, name, description, price, quantity)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteProduct(id int) {
	db := db.SqlConnection()

	sql := "DELETE FROM products WHERE id = $1;"

	_, err := db.Exec(sql, id)
	if err != nil {
		log.Fatal(err)
	}
}
