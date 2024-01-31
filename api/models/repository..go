package models

type Repository struct {
	ID 		  string`json:"id"`
	ProductID string`json:"product_id"`
	BranchID  string`json:"branch_id"`
	Count 	  string`json:"count"`
	CreatedAt string`json:"created_at"`
	UpdatedAt string`json:"updated_at"`
	DeletedAt string`json:"deleted_at"`
}

type CreateRepository struct {
	ProductID string`josn:"product_id"`
	BranchID  string`json:"branch_id"`
	Count 	  string`json:"count"`
}

type UpdateRepository struct {
	ID 		  string`json:"id"`
	ProductID string`json:"product_id"`
	BranchID  string`json:"branch_id"`
	Count 	  string`json:"count"`
}

type RepositoriesResponse struct {
	Repositories []Repository`json:"repositories"`
	Count                 int`json:"count"`
}
