package main

import (
	"ProductApiApp/common/app"
	"ProductApiApp/common/postgresql"
	"ProductApiApp/controllers"
	"ProductApiApp/persistence"
	"ProductApiApp/services"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)
	productService := services.NewProductService(productRepository)

	productController := controllers.NewProductController(productService)

	productController.RegisterRoutes(e)

	e.Start(":5173")
}
