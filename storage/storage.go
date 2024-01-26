package storage

import "developer/api/models"



type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale()   ISaleStorage
}

type IBranchStorage interface {
	Create(models.CreateBranch)(string, error)
	GetByID(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchesResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}

type ISaleStorage interface {
	Create(models.CreateSale)(string, error)
	GetByID(models.PrimaryKey) (models.Sale, error)
	GetList(models.GetListRequest) (models.SalesResponse, error)
	Update(models.UpdateSale) (string, error)
	Delete(models.PrimaryKey) error
}