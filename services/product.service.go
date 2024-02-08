package services

import (
	"ProductApiApp/domain"
	"ProductApiApp/dto"
	"ProductApiApp/persistence"
	"ProductApiApp/utils/result"
	"time"
)

type IProductService interface {
	AddNewProduct(createProductDto dto.CreateProductDTO) result.Result[bool]
	UpdateProduct(id int64, updateProductDto dto.UpdateProductDTO) result.Result[dto.ProductDTO]
	DeleteProduct(id int64) result.Result[bool]
	GetAllProducts() result.Result[dto.ProductListDTO]
	GetProductByID(id int64) result.Result[dto.ProductDTO]
	GetProductsByStore(store string) result.Result[dto.ProductListDTO]
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func validateProductCreate(createProductDto dto.CreateProductDTO) result.Result[bool] {
	if createProductDto.Name == "" {
		return result.Error[bool](false, "Name is required")
	}

	if createProductDto.Price < 0 {
		return result.Error[bool](false, "Price is required")
	}

	if createProductDto.Store == "" {
		return result.Error[bool](false, "Store is required")
	}

	return result.Success[bool](true, "Validated")
}

// AddNewProduct implements IProductService.
func (productService *ProductService) AddNewProduct(createProductDto dto.CreateProductDTO) result.Result[bool] {

	validationResult := validateProductCreate(createProductDto)

	if validationResult.Success == false {
		return validationResult
	}

	var product domain.Product = domain.Product{
		Name:     createProductDto.Name,
		Price:    createProductDto.Price,
		IsActive: createProductDto.IsActive,
		Discount: createProductDto.Discount,
		Store:    createProductDto.Store,
	}

	result := productService.productRepository.Save(product)

	return result
}

// DeleteProduct implements IProductService.
func (productService *ProductService) DeleteProduct(id int64) result.Result[bool] {

	if id <= 0 {
		return result.Error[bool](false, "Invalid ID")
	}

	result := productService.productRepository.Delete(id)

	return result
}

// GetAllProducts implements IProductService.
func (productService *ProductService) GetAllProducts() result.Result[dto.ProductListDTO] {
	var productsListDTO dto.ProductListDTO

	res := productService.productRepository.FindAll()

	if res.Success {
		for _, product := range res.Data {
			productsListDTO.Products = append(productsListDTO.Products, dto.ProductDTO{
				Id:        product.Id,
				Name:      product.Name,
				Price:     product.Price,
				IsActive:  product.IsActive,
				Discount:  product.Discount,
				Store:     product.Store,
				CreatedAt: time.Time.String(product.CreatedAt),
			})
		}

		return result.Success[dto.ProductListDTO](productsListDTO, res.Message)
	}
	return result.Error[dto.ProductListDTO](dto.ProductListDTO{}, res.Message)
}

// GetProductByID implements IProductService.
func (productService *ProductService) GetProductByID(id int64) result.Result[dto.ProductDTO] {
	if id <= 0 {
		return result.Error[dto.ProductDTO](dto.ProductDTO{}, "Invalid ID")
	}

	res := productService.productRepository.FindByID(id)

	if res.Success {
		product := res.Data
		return result.Success[dto.ProductDTO](dto.ProductDTO{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			IsActive:  product.IsActive,
			Discount:  product.Discount,
			Store:     product.Store,
			CreatedAt: time.Time.String(product.CreatedAt),
		}, res.Message)
	}
	return result.Error[dto.ProductDTO](dto.ProductDTO{}, res.Message)
}

// GetProductsByStore implements IProductService.
func (productService *ProductService) GetProductsByStore(store string) result.Result[dto.ProductListDTO] {
	if store == "" {
		return result.Error[dto.ProductListDTO](dto.ProductListDTO{}, "Invalid Store")
	}

	var productsListDTO dto.ProductListDTO

	res := productService.productRepository.FindAllByStore(store)

	if res.Success {
		for _, product := range res.Data {
			productsListDTO.Products = append(productsListDTO.Products, dto.ProductDTO{
				Id:        product.Id,
				Name:      product.Name,
				Price:     product.Price,
				IsActive:  product.IsActive,
				Discount:  product.Discount,
				Store:     product.Store,
				CreatedAt: time.Time.String(product.CreatedAt),
			})
		}

		return result.Success[dto.ProductListDTO](productsListDTO, res.Message)
	}
	return result.Error[dto.ProductListDTO](dto.ProductListDTO{}, res.Message)
}

// UpdateProduct implements IProductService.
func (productService *ProductService) UpdateProduct(id int64, updateProductDto dto.UpdateProductDTO) result.Result[dto.ProductDTO] {
	if id <= 0 {
		return result.Error[dto.ProductDTO](dto.ProductDTO{}, "Invalid ID")
	}

	var product domain.Product = domain.Product{
		Id:       updateProductDto.Id,
		Name:     updateProductDto.Name,
		Price:    updateProductDto.Price,
		IsActive: updateProductDto.IsActive,
		Discount: updateProductDto.Discount,
		Store:    updateProductDto.Store,
	}

	res := productService.productRepository.Update(id, product)

	if res.Success {
		product := res.Data
		return result.Success[dto.ProductDTO](dto.ProductDTO{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			IsActive:  product.IsActive,
			Discount:  product.Discount,
			Store:     product.Store,
			CreatedAt: time.Time.String(product.CreatedAt),
		}, res.Message)
	}
	return result.Error[dto.ProductDTO](dto.ProductDTO{}, res.Message)
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}
