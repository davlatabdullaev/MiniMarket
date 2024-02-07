package models

import "time"

type Sale struct {
	ID 			    string`json:"id"`
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistent_id"`
	CashierID       string`json:"chashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
	CreatedAt time.Time`json:"created_at"` 
	UpdatedAt time.Time`json:"updated_at"`
	DeletedAt time.Time`json:"deleted_at"`
}

type CreateSale struct {
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistent_id"`
	CashierID       string`json:"chashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
}

type UpdateSale struct {
	ID 			    string`json:"id"`
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistent_id"`
	CashierID       string`json:"chashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
}

type SalesResponse struct {
	Sales []Sale`json:"sales"`
	Count    int`json:"count"`
}

type StartSale struct {
	ProductId string`json:"product_id"`
	Quantity     int`json:"quantity"`
}