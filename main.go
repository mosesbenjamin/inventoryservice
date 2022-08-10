package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mosesbenjamin/inventoryservice/database"
	"github.com/mosesbenjamin/inventoryservice/product"
	"github.com/mosesbenjamin/inventoryservice/receipt"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	receipt.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
