package models

type Product struct {
	ID	   	   string`json:"id"`
	Name       string`json:"name"`
	Price      string`json:"price"`
	BarCode    string`json:"bar_code"`
	CategoryID string`json:"category_id"`
	CreatedAt  string`json:"created_at"`
	UpdatedAt  string`json:"updated_at"`
	DeletedAt  string`json:"deleted_at"`
}

type CreateProduct struct {
	Name       string`json:"name"`
	Price      string`json:"price"`
	BarCode    string`json:"bar_code"`
	CategoryID string`json:"category_id"`
}

type UpdateProduct struct {
	ID 		   string`json:"id"`
	Name       string`json:"name"`
	Price      string`json:"price"`
	CategoryID string`json:"category_id"`
}

type ProductsResponse struct {
	Products []Product`json:"products"`
	Count          int`json:"count"`
}