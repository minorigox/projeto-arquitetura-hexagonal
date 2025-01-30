package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/minorigox/projeto-arquitetura-hexagonal/adapters/db"
	"github.com/minorigox/projeto-arquitetura-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDBAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDBAdapter)
	product, _ := productService.Create("Product Example", 30)
	productService.Enable(product)
}