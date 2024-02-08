package domain

import "time"

type Product struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	IsActive  bool      `json:"is_active"`
	Discount  float64   `json:"discount"`
	Store     string    `json:"store"`
	CreatedAt time.Time `json:"created_at"`
}
