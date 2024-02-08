package controllers

import (
	"ProductApiApp/dto"
	"ProductApiApp/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService services.IProductService
}

func NewProductController(productService services.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products/store/:store", productController.GetProductsByStore)
	e.POST("/api/v1/products", productController.AddNewProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProduct)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	result := productController.productService.GetProductByID(idInt)

	if result.Success {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	result := productController.productService.GetAllProducts()

	if result.Success {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}

func (productController *ProductController) GetProductsByStore(c echo.Context) error {
	store := c.Param("store")
	result := productController.productService.GetProductsByStore(store)

	if result.Success {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}

func (productController *ProductController) AddNewProduct(c echo.Context) error {
	product := new(dto.CreateProductDTO)
	if err := c.Bind(product); err != nil {
		return c.JSON(400, "Invalid product")
	}
	result := productController.productService.AddNewProduct(*product)

	if result.Success {
		return c.JSON(http.StatusCreated, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}

func (productController *ProductController) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	product := new(dto.UpdateProductDTO)
	if err := c.Bind(product); err != nil {
		return c.JSON(400, "Invalid product")
	}
	result := productController.productService.UpdateProduct(idInt, *product)

	if result.Success {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	result := productController.productService.DeleteProduct(idInt)

	if result.Success {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusBadRequest, result)
}
