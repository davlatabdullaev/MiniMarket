package models

import "time"

type Branch struct {
	ID 	      string`json:"id"`
	Name      string`json:"name"`
	Address   string`json:"address"`
	CreatedAt time.Time`json:"created_at"`
	UpdatedAt time.Time`json:"updated_at"`
	DeletedAt time.Time`json:"deleted_at"`
}

type CreateBranch struct {
	Name      string`json:"name"`
	Address   string`json:"address"`
	CreatedAt string`json:"created_at"`
}

type UpdateBranch struct {
	ID 	      string   `json:"id"`
	Name      string   `json:"name"`
	Address   string   `json:"address"`
}

type BranchesResponse struct {
	Branches []Branch`json:"branches"`
	Count         int`json:"count"`
}