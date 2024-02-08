package dto

type CreateProductDTO struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	IsActive bool    `json:"is_active"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}

type UpdateProductDTO struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	IsActive bool    `json:"is_active"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}

type ProductDTO struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	IsActive  bool    `json:"is_active"`
	Discount  float64 `json:"discount"`
	Store     string  `json:"store"`
	CreatedAt string  `json:"created_at"`
}

type ProductListDTO struct {
	Products []ProductDTO `json:"products"`
}
