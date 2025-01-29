package db_test	

import (
	"database/sql"
	"log"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/minorigox/projeto-arquitetura-hexagonal/adapters/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minorigox/projeto-arquitetura-hexagonal/application"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
}

func createTable(db *sql.DB) {
	table := `create table products (
	"id" string,
	"name" string,
	"price" float,
	"status" string
	);`
	stmt, err := db.Prepare(table)
	if (err != nil) {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if (err != nil) {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	createProduct(Db)
	productDB := db.NewProductDB(Db)
	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Create(t *testing.T) {
	setUp()
	defer Db.Close()
	productDB := db.NewProductDB(Db)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}