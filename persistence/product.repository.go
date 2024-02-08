package persistence

import (
	"ProductApiApp/domain"
	"ProductApiApp/utils/messages/constants"
	"ProductApiApp/utils/result"
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	FindAll() result.Result[[]domain.Product]
	FindAllByStore(store string) result.Result[[]domain.Product]
	FindByID(id int64) result.Result[domain.Product]
	Save(product domain.Product) result.Result[bool]
	Update(id int64, product domain.Product) result.Result[domain.Product]
	Delete(id int64) result.Result[bool]
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

/**
 * @description: FindAll method to get all products
 * @return {[]domain.Product}
 * @param {*ProductRepository} productRepository
 */
func (productRepository *ProductRepository) FindAll() result.Result[[]domain.Product] {
	ctx := context.Background()
	rows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")

	if err != nil {
		log.Error("Error while querying products: ", err)
		return result.Error[[]domain.Product]([]domain.Product{}, constants.ERROR_WHILE_FETCHING_PRODUCTS)
	}

	// if rows.CommandTag().RowsAffected() == 0 {
	// 	return result.Error[[]domain.Product]([]domain.Product{}, constants.PRODUCTS_NOT_FOUND)
	// }

	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		var createdAt time.Time
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.IsActive, &product.Discount, &product.Store, &createdAt)

		if err != nil {
			log.Error("Error while scanning products: ", err)
			return result.Error[[]domain.Product]([]domain.Product{}, constants.PRODUCT_SCAN_ERROR)
		}

		product.CreatedAt, _ = time.Parse(time.RFC3339, createdAt.Format(time.RFC3339))
		products = append(products, product)
	}

	return result.Success[[]domain.Product](products, constants.PRODUCTS_FOUND)
}

func (productRepository *ProductRepository) FindAllByStore(store string) result.Result[[]domain.Product] {

	ctx := context.Background()
	rows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products WHERE store = $1", store)

	if err != nil {
		log.Error(constants.ERROR_WHILE_QUERYING_PRODUCTS, err)
		return result.Error[[]domain.Product](nil, constants.ERROR_WHILE_QUERYING_PRODUCTS)
	}

	// if rows.CommandTag().RowsAffected() == 0 {
	// 	return result.Error[[]domain.Product](nil, constants.PRODUCTS_NOT_FOUND)
	// }

	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.IsActive, &product.Discount, &product.Store, &product.CreatedAt)

		if err != nil {
			log.Error(constants.ERROR_WHILE_SCANNING_PRODUCTS, err)
			return result.Error[[]domain.Product]([]domain.Product{}, constants.ERROR_WHILE_SCANNING_PRODUCTS)
		}

		products = append(products, product)
	}

	return result.Success[[]domain.Product](products, constants.PRODUCTS_FOUND)
}

// Save implements IProductRepository.
func (productRepository *ProductRepository) Save(product domain.Product) result.Result[bool] {
	ctx := context.Background()

	addedProduct, err := productRepository.dbPool.Exec(ctx,
		"INSERT INTO products (name, price, is_active, discount, store) VALUES ($1, $2, $3, $4, $5)",
		product.Name, product.Price, product.IsActive, product.Discount, product.Store)

	if err != nil {
		log.Error(constants.PRODUCT_INSERT_ERROR, err)
		return result.Error[bool](false, constants.PRODUCT_INSERT_ERROR)
	}

	log.Info("Product added: ", addedProduct)

	if addedProduct.RowsAffected() > 0 {
		return result.Success[bool](true, constants.PRODUCT_ADDED)
	}

	return result.Error[bool](false, constants.PRODUCT_INSERT_ERROR)
}

// Delete implements IProductRepository.
func (productRepository *ProductRepository) Delete(id int64) result.Result[bool] {
	ctx := context.Background()
	_, err := productRepository.dbPool.Exec(ctx, "DELETE FROM products WHERE id = $1", id)

	if err != nil {
		log.Error(constants.PRODUCT_DELETE_ERROR, err)
		return result.Error[bool](false, constants.PRODUCT_DELETE_ERROR)
	}

	return result.Success[bool](true, constants.PRODUCT_DELETED)
}

// FindByID implements IProductRepository.
func (productRepository *ProductRepository) FindByID(id int64) result.Result[domain.Product] {
	ctx := context.Background()
	row := productRepository.dbPool.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", id)

	var product domain.Product
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.IsActive, &product.Discount, &product.Store, &product.CreatedAt)

	if err != nil {
		log.Error(constants.PRODUCT_SCAN_ERROR, err)
		return result.Error[domain.Product](domain.Product{}, constants.PRODUCT_SCAN_ERROR)
	}

	return result.Success[domain.Product](product, constants.PRODUCT_FOUND)
}

// Update implements IProductRepository.
func (productRepository *ProductRepository) Update(id int64, product domain.Product) result.Result[domain.Product] {
	ctx := context.Background()

	updatedProduct, err := productRepository.dbPool.Exec(ctx,
		"UPDATE products SET name = $1, price = $2, is_active = $3, discount = $4, store = $5 WHERE id = $7",
		product.Name, product.Price, product.IsActive, product.Discount, product.Store, id)

	if err != nil {
		log.Error(constants.PRODUCT_UPDATE_ERROR, err)
		return result.Error[domain.Product](domain.Product{}, constants.PRODUCT_UPDATE_ERROR)
	}

	log.Info("Product updated: ", updatedProduct)

	if updatedProduct.RowsAffected() > 0 {
		return result.Success[domain.Product](product, constants.PRODUCT_UPDATED)
	}

	return result.Error[domain.Product](domain.Product{}, constants.PRODUCT_UPDATE_ERROR)
}
