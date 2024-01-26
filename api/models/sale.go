package models

import "time"

type Sale struct {
	ID 			    string`json:"id"`
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistant_id"`
	CashierID       string`json:"cashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
	CreatedAt       string`json:"created_at"`
	UpdatedAt       string`json:"updated_at"`
	DeletedAt       string`json:"deleted_at"`
}

type CreateSale struct {
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistant_id"`
	CashierID       string`json:"cashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
	CreatedAt       string`json:"created_at"`
	UpdatedAt       string`json:"updated_at"`
	DeletedAt       string`json:"deleted_at"`
}

type UpdateSale struct {
	ID 			    string`json:"id"`
	BranchID        string`json:"branch_id"`
	ShopAssistantID string`json:"shop_assistant_id"`
	CashierID       string`json:"cashier_id"`
	PaymentType     string`json:"payment_type"`
	Price             uint`json:"price"`
	Status          string`json:"status"`
	ClientName      string`json:"client_name"`
	UpdatedAt       time.Time`json:"updated_at"`
}

type SalesResponse struct {
	Sales []Sale`json:"sales"`
	Count    int`json:"count"`
}