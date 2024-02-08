package infrastructure

import (
	"ProductApiApp/common/postgresql"
	"ProductApiApp/persistence"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "postgres",
		Database:              "product_db",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)

	code := m.Run()

	dbPool.Close()

	os.Exit(code)
}

func TestFindAllProducts(t *testing.T) {
	t.Run("FindAllProducts", func(t *testing.T) {
		products:= productRepository.FindAll()
		fmt.Println("Products: ", products)
	})
}
