package main

import (
	"database/sql"

	db2 "github.com/kalinskilk/arquitetura-hexagonal/adapters/db"
	"github.com/kalinskilk/arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product example", 30)

	productService.Enable(product)

}
